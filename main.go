package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/Harsh-apk/notesPostgres/db"
	"github.com/Harsh-apk/notesPostgres/handlers"
	"github.com/Harsh-apk/notesPostgres/types"
	"github.com/gofiber/fiber/v2"
)

func main() {
	file, err := os.ReadFile("env.json")
	if err != nil {
		log.Fatal(err)
	}
	var data types.ENV_DATA
	err = json.Unmarshal(file, &data)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(data.USER)

	db := db.NewPostgressUserDb(Connect())
	userHandler := handlers.NewUserHandler(db, &data)

	app := fiber.New()
	app.Get("/", userHandler.HandleGetAllUsers)
	app.Get("/auth/:id", userHandler.HandleAuthenticateEmail)
	app.Post("/create", userHandler.HandleCreateUser)
	app.Listen(":3000")

}
