package main

import (
	"context"
	"crypto/tls"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/joho/godotenv"
	"github.com/segmentio/kafka-go"
	"github.com/segmentio/kafka-go/sasl/scram"
	"github.com/yaninyzwitty/kafka-producer-go/repository"
	"github.com/yaninyzwitty/kafka-producer-go/service"
	"github.com/yaninyzwitty/kafka-producer-go/transport"
)

func main() {
	ctx := context.Background()
	// load env variables
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file", err)
	}
	KAFKA_BOOTSTRAP_SERVER := os.Getenv("KAFKA_BOOTSTRAP_SERVER")
	KAFKA_API_KEY := os.Getenv("KAFKA_API_KEY")
	KAFKA_API_SECRET := os.Getenv("KAFKA_API_SECRET")
	KAFKA_TOPIC := os.Getenv("KAFKA_TOPIC")

	// setup kafka here using kafka-go (producer)
	mechanism, _ := scram.Mechanism(scram.SHA256, KAFKA_API_KEY, KAFKA_API_SECRET)
	writer := kafka.Writer{
		Addr:  kafka.TCP(KAFKA_BOOTSTRAP_SERVER),
		Topic: KAFKA_TOPIC,
		Transport: &kafka.Transport{
			SASL: mechanism,
			TLS:  &tls.Config{},
		},
	}
	defer writer.Close()

	// Initialize dependencies (products)

	productRepo := repository.NewProductRepository(ctx, &writer)
	productService := service.NewProductService(productRepo)
	productHandler := transport.NewProductHandler(productService)

	// initialize the router, chi, gin or echo
	router := chi.NewRouter()
	router.Use(middleware.Logger)
	// Define routes (products)

	router.Post("/products", productHandler.CreateProduct)

	// Start the server
	log.Println("Server started on :3000 😂")
	log.Fatal(http.ListenAndServe(":3000", router))

}
