package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Objeto representa la estructura de tus objetos en la colección MongoDB.
type Objeto struct {
	ID        string    `json:"id" bson:"_id,omitempty"`
	Nombre    string    `json:"message" bson:"message"`
	Timestamp time.Time `json:"timestamp" bson:"timestamp"` // Fecha de creación
	Tipo      string    `json:"tipo" bson:"tipo"`           // Tipo de log
}

// ListaObjetosPaginados devuelve una página de objetos desde la colección MongoDB.
func ListaObjetosPaginados(w http.ResponseWriter, r *http.Request) {
	// Configura la conexión a MongoDB
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatal(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(ctx)

	collection := client.Database("myDatabase").Collection("messages")

	// Obtener parámetros de filtrado y paginación
	page, limit, err := obtenerParametrosPaginacion(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	filtro, err := obtenerFiltro(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	opts := options.Find().SetLimit(int64(limit)).SetSkip(int64((page - 1) * limit)).SetSort(bson.D{{"timestamp", 1}}) // Ordenar por fecha de creación

	cur, err := collection.Find(ctx, filtro, opts)
	if err != nil {
		http.Error(w, "Error al obtener objetos desde MongoDB", http.StatusInternalServerError)
		return
	}
	defer cur.Close(ctx)

	// Itera sobre los documentos y los agrega a la lista de objetos
	var objetos []Objeto
	for cur.Next(ctx) {
		var objeto Objeto
		err := cur.Decode(&objeto)
		if err != nil {
			http.Error(w, "Error al decodificar objeto desde MongoDB", http.StatusInternalServerError)
			return
		}
		objetos = append(objetos, objeto)
	}

	// Verifica si hubo algún error durante la iteración
	if err := cur.Err(); err != nil {
		http.Error(w, "Error durante la iteración de documentos en MongoDB", http.StatusInternalServerError)
		return
	}

	// Convierte la lista de objetos a formato JSON
	jsonBytes, err := json.Marshal(objetos)
	if err != nil {
		http.Error(w, "Error al convertir a JSON", http.StatusInternalServerError)
		return
	}

	// Establece el encabezado Content-Type como application/json
	w.Header().Set("Content-Type", "application/json")

	// Escribe la respuesta
	w.Write(jsonBytes)
}

// obtenerParametrosPaginacion procesa los parámetros de paginación de la solicitud
func obtenerParametrosPaginacion(r *http.Request) (int, int, error) {
	pageStr := r.URL.Query().Get("page")
	page, err := strconv.Atoi(pageStr)
	if err != nil || page < 1 {
		page = 1
	}

	limitStr := r.URL.Query().Get("limit")
	limit, err := strconv.Atoi(limitStr)
	if err != nil || limit < 1 {
		limit = 10
	}

	return page, limit, nil
}

// obtenerFiltro construye un filtro para la consulta MongoDB basado en los parámetros de la solicitud
func obtenerFiltro(r *http.Request) (bson.M, error) {
	filtro := bson.M{}

	// Filtrar por tipo de log si está presente
	tipo := r.URL.Query().Get("tipo")
	if tipo != "" {
		filtro["tipo"] = tipo
	}

	// Filtrar por rango de fechas si está presente
	desde := r.URL.Query().Get("desde")
	hasta := r.URL.Query().Get("hasta")
	if desde != "" && hasta != "" {
		desdeTime, err := time.Parse(time.RFC3339, desde)
		if err != nil {
			return nil, fmt.Errorf("formato de fecha 'desde' inválido: %s", err)
		}
		hastaTime, err := time.Parse(time.RFC3339, hasta)
		if err != nil {
			return nil, fmt.Errorf("formato de fecha 'hasta' inválido: %s", err)
		}
		filtro["timestamp"] = bson.M{"$gte": desdeTime, "$lte": hastaTime}
	}

	return filtro, nil
}

func main1() {
	// Crea un enrutador usando Gorilla Mux
	r := mux.NewRouter()

	// Define la ruta para el endpoint de listar objetos paginados
	r.HandleFunc("/logs", ListaObjetosPaginados).Methods("GET")

	// Configura el servidor HTTP con el enrutador
	port := 8090
	fmt.Printf("Servidor escuchando en el puerto %d...\n", port)
	http.Handle("/", r)
	http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
}
