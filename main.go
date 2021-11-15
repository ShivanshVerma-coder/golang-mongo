package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/ShivanshVerma-coder/golang-mongo/database"
	"github.com/ShivanshVerma-coder/golang-mongo/router"
	"github.com/joho/godotenv"
)

func goDotEnvVariable(key string) string {

	// load .env file
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	return os.Getenv(key)
}

func main() {
	fmt.Println("Golang-Mongo_Backend")
	connectionString := goDotEnvVariable("MONGO_URI")
	database.ConnectDB(connectionString)
	r := router.Router()
	fmt.Println("Server is getting started...")
	log.Fatal(http.ListenAndServe(":8080", r))
	fmt.Println("Listening at port 8080")
}
