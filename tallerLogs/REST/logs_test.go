package main

import (
	"context"
	"encoding/json"
	"log"
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

// TestListaObjetosPaginados simula una prueba del endpoint de listado paginado.
func TestListaObjetosPaginados(t *testing.T) {
	// Crear un request HTTP al endpoint.
	req, err := http.NewRequest("GET", "/logs?page=1&limit=1", nil)
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

	// Parsear la respuesta en una estructura de logs
	var logs []Objeto // Objeto es la estructura que representa un log
	err = json.NewDecoder(rr.Body).Decode(&logs)
	if err != nil {
		t.Fatal(err)
	}

	// Comprobar que la respuesta es la esperada.
	if len(logs) != 1 {
		t.Errorf("handler returned unexpected number of logs: got %v want %v", len(logs), 1)
	}

}

// TestFiltroLogs para el filtro de logs solo chequea que el tipo sea exito en todos los logs
func TestFiltroLogs(t *testing.T) {
	// Conexión a MongoDB
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(context.TODO())

	collection := client.Database("myDatabase").Collection("messages")

	// Recuperar los logs
	cursor, err := collection.Find(context.TODO(), bson.D{})
	if err != nil {
		t.Fatal(err)
	}

	var logs []bson.M
	if err = cursor.All(context.TODO(), &logs); err != nil {
		t.Fatal(err)
	}

	// Verificar que los objetos retornados cumplen con el filtro aplicado.
	for _, log := range logs {
		mensaje := log["message"].(string)
		if tipo := determinarTipo(mensaje); tipo != "exito" {
			t.Errorf("El log no cumple con el filtro: got %v want %v", tipo, "exito")
		}
	}
}

// TestInsercionLog simula una prueba de inserción de un log en MongoDB.
func TestInsercionLog(t *testing.T) {
	mockCollection := new(MockMongoCollection)
	mensaje := "Mensaje de prueba"
	tipoMensaje := determinarTipo(mensaje)

	mockCollection.On("InsertOne", mock.Anything, bson.D{
		{Key: "message", Value: mensaje},
		{Key: "timestamp", Value: mock.AnythingOfType("time.Time")},
		{Key: "tipo", Value: tipoMensaje},
	}).Return(nil, nil)

	err := insertarLogEnMongoDB(mockCollection, mensaje)
	if err != nil {
		t.Errorf("Error al insertar en MongoDB: %s", err)
	}
}

//-----------------------------------------------------------------------------------------------------------------------------------------------------

type MongoCollection interface {
	InsertOne(ctx context.Context, document interface{}, opts ...*options.InsertOneOptions) (*mongo.InsertOneResult, error)
}

func insertarLogEnMongoDB(collection MongoCollection, mensaje string) error {
	tipo := determinarTipo(mensaje)

	_, err := collection.InsertOne(context.TODO(), bson.D{
		{Key: "message", Value: mensaje},
		{Key: "timestamp", Value: time.Now()},
		{Key: "tipo", Value: tipo},
	})

	return err
}

type MockMongoCollection struct {
	mock.Mock
}

func (m *MockMongoCollection) InsertOne(ctx context.Context, document interface{}, opts ...*options.InsertOneOptions) (*mongo.InsertOneResult, error) {
	args := m.Called(ctx, document)
	return args.Get(0).(*mongo.InsertOneResult), args.Error(1)
}

func determinarTipo(mensaje string) string {
	if strings.HasPrefix(mensaje, "Exito") {
		return "exito"
	} else if strings.HasPrefix(mensaje, "Error") {
		return "error"
	}
	return "desconocido" // Tipo por defecto si no coincide con los anteriores
}
