package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/joho/godotenv"

	_ "github.com/lib/pq"
)

func main(){
	err := godotenv.Load()
	if err != nil{
		log.Fatalf("ERROR loading .env: %v",err)
	}

	db_url := os.Getenv("DB_URL")
	portString := ":"+os.Getenv("PORT")

	db, err := sql.Open("postgres", db_url)

	defer db.Close()
	if err!=nil{
		log.Fatalf("ERR opening database: %v",err)
	}
	if err=db.Ping(); err!=nil{
		log.Fatal(err)
	}

	router := chi.NewRouter()

	router.Use(middleware.Logger)

	fmt.Printf("db open and server on %v",portString)
	err = http.ListenAndServe(portString, router)
	if err != nil{
		log.Fatal(err)
	}
}

//handlerFunc is func(http.ResponseWriter, *http.Request) 

/* this will containe the router for to which the frontend connects
* then it will also initialise the connection to the database
* it will call the handlers for everything else
 */
