package main


import (
	"log"
	"os"
	"testing"

	"."
)
const tableCreationQuery = `CREATE TABLE IF NOT EXISTS products
(
    id SERIAL,
    name TEXT NOT NULL,
    price NUMERIC(10,2) NOT NULL DEFAULT 0.00,
    CONSTRAINT products_pkey PRIMARY KEY (id)
)`


var app App

func TestMain(m *testing.M){
	app = App{}
	app.Initialize(os.Getenv("APP_DB_USERNAME"),os.Getenv("APP_DB_PASSWORD"),os.Getenv("APP_DB_NAME"))

	ensureTableExists()

	code := m.Run()
	clearTable()
	os.Exit(code)
	
	
}

func ensureTableExists()  {
	if _,err := app.DB.Exec(tableCreationQuery);err != nil{
		log.Fatal(err)
	}
}

func clearTable()  {
	
	app.DB.Exec("DELETE FROM products")
	app.DB.Exec("ALTER SEQUENCE products_id_seq RESTART WITH 1")
}
