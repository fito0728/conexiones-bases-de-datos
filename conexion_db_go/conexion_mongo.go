package main

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func conectar_mongo() {
	// Establecer la conexión con MongoDB
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017") // Cambia la URI si es necesario
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatalf("Error al conectar a MongoDB: %v", err)
	}
	defer client.Disconnect(context.TODO())

	// Verificar la conexión
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatalf("Error al hacer ping a MongoDB: %v", err)
	}
	fmt.Println("Conexión exitosa a MongoDB")

	// Seleccionar la base de datos y la colección
	collection := client.Database("testdb").Collection("users")

	// Crear un documento para insertar
	user := bson.D{
		{"name", "John Doe"},
		{"age", 30},
	}

	// Insertar el documento en la colección
	insertResult, err := collection.InsertOne(context.TODO(), user)
	if err != nil {
		log.Fatalf("Error al insertar documento: %v", err)
	}
	fmt.Printf("Documento insertado con ID: %v\n", insertResult.InsertedID)

	// Consultar el documento
	var result bson.D
	err = collection.FindOne(context.TODO(), bson.D{{"name", "John Doe"}}).Decode(&result)
	if err != nil {
		log.Fatalf("Error al consultar documento: %v", err)
	}

	// Mostrar el resultado
	fmt.Printf("Documento encontrado: %v\n", result)
}