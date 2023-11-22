package main

import (
	"context"
	"log"
	"strings"
	"time"

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

	collection := client.Database("myDatabase").Collection("messages")

	// Procesar mensajes
	go func() {
		for d := range msgs {
			log.Printf("Received a message: %s", d.Body)

			mensaje := string(d.Body)
			tipo := determinarTipo(mensaje)
			appId := extraerAppId(mensaje)

			_, err := collection.InsertOne(context.TODO(), bson.D{
				{Key: "message", Value: mensaje},
				{Key: "timestamp", Value: time.Now()},
				{Key: "tipo", Value: tipo},
				{Key: "application", Value: appId},
			})
			if err != nil {
				log.Printf("Error al insertar mensaje en MongoDB: %s", err)
			}
		}
	}()

	log.Printf(" [*] Esperando mensajes. Para salir presiona CTRL+C")
	<-make(chan bool)
}

func determinarTipo(mensaje string) string {
	if strings.HasPrefix(mensaje, "Exito") {
		return "exito"
	} else if strings.HasPrefix(mensaje, "Error") {
		return "error"
	}
	return "desconocido" // Tipo por defecto si no coincide con los anteriores
}

func extraerAppId(mensaje string) string {
	if strings.Contains(mensaje, "Aut") {
		return "aut"
	} else if strings.Contains(mensaje, "Rest") {
		return "rest"
	}
	return "desconocido" // Tipo por defecto si no coincide con los anteriores
}
