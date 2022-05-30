package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/axtoneIO/go_rest_api/internal/database"
	transport_http "github.com/axtoneIO/go_rest_api/internal/transport/http"
	"github.com/joho/godotenv"
)

// App - the struct which contains things like pointers
// to database connections
type App struct {}

// Run - sets up our application
func (app *App) Run() error {
	fmt.Println("Setting up our App")

	godotenv.Load(".env")

	var err error

	_, err = database.NewDatabase()
	if err != nil {
		return err
	}

	handler := transport_http.NewHandler()
	handler.SetupRoutes()

	if err := http.ListenAndServe("localhost:8000",handler.Router); err != nil {
		log.Print(err)
		return err
	}
	
	return nil
}

func main() {
	fmt.Println("Go REST API")
	app := App{}
	if err := app.Run(); err != nil{
		fmt.Println("Error starting up our REST API")
		fmt.Println(err)
	}
}