package main

import (
	"github.com/gorilla/mux"
	"github.com/lib/pq"
	"os"
)

func main() {
	app := App{}
	app.Initialize(
		os.Getenv("APP_DB_USERNAME"),
		os.Getenv("APP_DB_PASSWORD"),
		os.Getenv("APP_DB_NAME"))

	app.Run(":8080")

	fmt.Println()
}
