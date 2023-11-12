package main

import (
	"context"
	"log"

	"github.com/streadway/amqp"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	// Conexión a RabbitMQ
	conn, err := amqp.Dial("amqp://guest:guest@rabbitmq:5672/")
	if err != nil {
		log.Fatalf("Failed to connect to RabbitMQ: %s", err)
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("Failed to open a channel: %s", err)
	}
	defer ch.Close()

	// Configuración de Exchange
	err = ch.ExchangeDeclare(
		"rootExchange", // name
		"direct",       // type
		true,           // durable
		false,          // auto-deleted
		false,          // internal
		false,          // no-wait
		nil,            // arguments
	)
	if err != nil {
		log.Fatalf("Failed to declare an exchange: %s", err)
	}

	// Declaración de Queue
	q, err := ch.QueueDeclare(
		"tuCola", // name
		true,     // durable
		false,    // delete when unused
		false,    // exclusive
		false,    // no-wait
		nil,      // arguments
	)
	if err != nil {
		log.Fatalf("Failed to declare a queue: %s", err)
	}

	// Bind Queue
	err = ch.QueueBind(
		q.Name,         // queue name
		"tuRoutingKey", // routing key
		"rootExchange", // exchange
		false,
		nil,
	)
	if err != nil {
		log.Fatalf("Failed to bind a queue: %s", err)
	}

	// Consumir mensajes
	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	if err != nil {
		log.Fatalf("Failed to register a consumer: %s", err)
	}

	// Conexión a MongoDB
    	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://mongodb:27017"))
    	if err != nil {
    		log.Fatal(err)
    	}
    	defer client.Disconnect(context.TODO())

    	// Obtén una lista de todas las bases de datos
    	databases, err := client.ListDatabaseNames(context.TODO(), bson.M{})
    	if err != nil {
    		log.Fatal(err)
    	}

    	// Verifica si la base de datos deseada está en la lista
    	databaseExists := false
    	for _, db := range databases {
    		if db == "myDatabase" {
    			databaseExists = true
    			break
    		}
    	}

    	// Si la base de datos no existe, crea una insertando un documento en una colección
    	if !databaseExists {
    		collection := client.Database("myDatabase").Collection("messages")
    		_, err := collection.InsertOne(context.TODO(), bson.D{
    			{Key: "message", Value: "Initial message to create database"},
    		})
    		if err != nil {
    			log.Printf("Error al insertar mensaje en MongoDB: %s", err)
    		}
    	}

    	// Obtén una colección (o crea una si no existe)
    	collection := client.Database("myDatabase").Collection("messages")

	// Procesar mensajes
	go func() {
		for d := range msgs {
			log.Printf("Received a message: %s", d.Body)

			// Insertar mensaje en MongoDB
			_, err := collection.InsertOne(context.TODO(), bson.D{
				{Key: "message", Value: string(d.Body)},
			})
			if err != nil {
				log.Printf("Error al insertar mensaje en MongoDB: %s", err)
			}
		}
	}()

	log.Printf(" [*] Esperando mensajes. Para salir presiona CTRL+C")
	<-make(chan bool)
}
