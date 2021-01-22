package main

import (
	"github.com/go-redis/redis/v8"
	"os"
	//"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"log"
	"net/http"
	//"strings"
	//"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
	//"time"
	"context"
	"github.com/rs/cors"
	"gitlab.com/cesarcedanov/social_expenses_tracker/model"
)

var (
	ctx = context.Background()
	client *redis.Client
)

func main(){
	log.Print("Starting expernses server")

	db, err = gorm.Open( "postgres",
		"host=localhost port=5432 user=cesarcedanov dbname=postgres sslmode=disable password=postgres")
	if err != nil {
		log.Fatalf("error pinging database: %v", err)
	}

	// Close Database at the end of the Scope or in case something fails
	defer db.Close()

	db.AutoMigrate(&model.User{})
	db.AutoMigrate(&model.Expense{})

	// Initialize Redis Client
	InitializeRedis()

	// Created a new Router and Set the Routes
	r := mux.NewRouter()
	addRoutes(r)

	handler := cors.Default().Handler(r)
	log.Fatal(http.ListenAndServe(":8080", handler))
}



func InitializeRedis() {
	//Initializing redis
	dsn := os.Getenv("REDIS_DSN")
	if len(dsn) == 0 {
		dsn = "localhost:6379"
	}
	client = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	_, err := client.Ping(ctx).Result()
	if err != nil {
		panic(err)
	}
	log.Print("Initialized Redis Client")
}