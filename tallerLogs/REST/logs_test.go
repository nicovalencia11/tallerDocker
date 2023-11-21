package main

import (
	"context"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// TestConexionMongoDB simula una prueba de conexión a MongoDB.
func TestConexionMongoDB(t *testing.T) {
	// Simular conexión a MongoDB.
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		t.Fatalf("No se pudo conectar a MongoDB: %s", err)
	}
	defer client.Disconnect(context.TODO())

	// Realizar una operación simple para verificar la conexión.
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		t.Fatalf("No se pudo hacer ping a MongoDB: %s", err)
	}
}

// TestInsercionLog simula una prueba de inserción de un log en MongoDB.
func TestInsercionLog(t *testing.T) {
	mockMongo := new(MockMongoDB)
	mensaje := "Mensaje de prueba"

	mockMongo.On("InsertOne", nil, bson.D{
		{Key: "message", Value: mensaje},
		{Key: "timestamp", Value: mock.AnythingOfType("time.Time")},
		{Key: "tipo", Value: determinarTipo(mensaje)},
	}).Return(nil)

	err := insertarLogEnMongoDB(mockMongo, mensaje)
	if err != nil {
		t.Errorf("Error al insertar en MongoDB: %s", err)
	}
}

// TestListaObjetosPaginados simula una prueba del endpoint de listado paginado.
func TestListaObjetosPaginados(t *testing.T) {
	// Crear un request HTTP al endpoint.
	req, err := http.NewRequest("GET", "/logs?page=1&limit=10", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Crear un ResponseRecorder para grabar la respuesta.
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(ListaObjetosPaginados)

	// Llamar al handler con nuestro Request y ResponseRecorder.
	handler.ServeHTTP(rr, req)

	// Comprobar que el estado de la respuesta es el esperado.
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	// Comprobar que la respuesta es la esperada.
	expected := `...` // Aquí iría la respuesta JSON esperada.
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
	}
}

func TestFiltroLogs(t *testing.T) {
	// Crear un request con parámetros de filtro, por ejemplo, por tipo.
	req, err := http.NewRequest("GET", "/logs?tipo=exito", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(ListaObjetosPaginados)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	// Verificar que los objetos retornados cumplen con el filtro aplicado.
	// Aquí deberías parsear la respuesta y verificar que todos los elementos tienen el tipo "exito".
}

func TestPaginacionLogs(t *testing.T) {
	req, err := http.NewRequest("GET", "/logs?page=2&limit=5", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(ListaObjetosPaginados)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	// Verificar que la cantidad de objetos en la respuesta es 5.
}

func TestBadRequestListaObjetosPaginados(t *testing.T) {
	req, err := http.NewRequest("GET", "/logs?page=-1", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(ListaObjetosPaginados)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusBadRequest)
	}

	// Verificar el mensaje de error de la respuesta.
}

func TestRabbitMQConsumer(t *testing.T) {
	// Simular la recepción de un mensaje desde RabbitMQ.
	// Verificar que el mensaje se procesa correctamente y se inserta en MongoDB.
}

//-----------------------------------------------------------------------------------------------------------------------------------------------------

type MongoDBInterface interface {
	InsertOne(interface{}, interface{}) error
}

type MockMongoDB struct {
	mock.Mock
}

func (m *MockMongoDB) InsertOne(ctx interface{}, document interface{}) error {
	args := m.Called(ctx, document)
	return args.Error(0)
}

func insertarLogEnMongoDB(db MongoDBInterface, mensaje string) error {
	return db.InsertOne(nil, bson.D{
		{Key: "message", Value: mensaje},
		{Key: "timestamp", Value: time.Now()},
		{Key: "tipo", Value: determinarTipo(mensaje)},
	})
}

func determinarTipo(mensaje string) string {
	if strings.HasPrefix(mensaje, "Exito") {
		return "exito"
	} else if strings.HasPrefix(mensaje, "Error") {
		return "error"
	}
	return "desconocido" // Tipo por defecto si no coincide con los anteriores
}
