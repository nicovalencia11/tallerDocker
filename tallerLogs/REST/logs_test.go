package main

import (
	"bytes"
	"context"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// TestConexionMongoDB simula una prueba de conexión a MongoDB.
func TestConexionMongoDB(t *testing.T) {
	// Simular conexión a MongoDB.
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://localhost:27018"))
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

// TestFiltroLogs simula una prueba de filtrado de logs.
func TestFiltroLogs(t *testing.T) {
	// Conexión a MongoDB
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://localhost:27018"))
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
	exitosos := 0
	errores := 0
	for _, log := range logs {
		mensaje := log["message"].(string)
		tipo := determinarTipo(mensaje)
		if tipo == "exito" {
			exitosos++
		} else if tipo == "error" {
			errores++
		}
	}

	if exitosos == 0 {
		t.Errorf("No se encontraron logs de exito")
	}

	if errores == 0 {
		t.Errorf("No se encontraron logs de error")
	}
}

// TestAgregarLog simula una prueba de inserción de logs.
func TestAgregarLog(t *testing.T) {
	// Crear un servidor HTTP de prueba
	ts := httptest.NewServer(http.HandlerFunc(AgregarLog))
	defer ts.Close()

	// Crear un objeto log para insertar
	log := Objeto{
		Nombre:      "Test log",
		Timestamp:   time.Now(),
		Tipo:        "exito",
		Application: "Test app",
	}

	// Convertir el objeto log a JSON
	jsonBytes, err := json.Marshal(log)
	if err != nil {
		t.Fatalf("Error al convertir el log a JSON: %s", err)
	}

	// Crear un request POST
	req, err := http.NewRequest("POST", ts.URL, bytes.NewBuffer(jsonBytes))
	if err != nil {
		t.Fatalf("Error al crear el request: %s", err)
	}

	// Llamar a la función AgregarLog
	rr := httptest.NewRecorder()
	AgregarLog(rr, req)

	// Verificar que la respuesta HTTP tenga el código de estado 201
	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("AgregarLog devolvió un código de estado incorrecto: obtuvo %v, esperaba %v", status, http.StatusCreated)
	}

	// Verificar que el cuerpo de la respuesta sea "Log agregado exitosamente"
	expected := "Log agregado exitosamente"
	if rr.Body.String() != expected {
		t.Errorf("AgregarLog devolvió un cuerpo incorrecto: obtuvo %v, esperaba %v", rr.Body.String(), expected)
	}
	// Agregar un pequeño retraso para dar tiempo a que se complete la inserción
	time.Sleep(2 * time.Second)

	// Conectar a la base de datos MongoDB y verificar que el log se haya insertado correctamente
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		t.Fatalf("Error al conectar a MongoDB: %s", err)
	}
	err = client.Connect(context.Background())
	if err != nil {
		t.Fatalf("Error al conectar a MongoDB: %s", err)
	}
	collection := client.Database("myDatabase").Collection("messages")
	filter := bson.M{"message": log.Nombre, "application": log.Application}
	result := collection.FindOne(context.Background(), filter)
	if result.Err() != nil {
		t.Errorf("Error al buscar el log en MongoDB: %s", result.Err())
	}
	var foundLog Objeto
	err = result.Decode(&foundLog)
	if err != nil {
		t.Errorf("Error al decodificar el log encontrado: %s", err)
	}
	if foundLog.Nombre != log.Nombre || foundLog.Application != log.Application {
		t.Errorf("El log encontrado no coincide con el log insertado")
	}
}

func determinarTipo(mensaje string) string {
	if strings.HasPrefix(mensaje, "Exito") {
		return "exito"
	} else if strings.HasPrefix(mensaje, "Error") {
		return "error"
	}
	return "desconocido" // Tipo por defecto si no coincide con los anteriores
}
