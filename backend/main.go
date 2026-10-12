package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/GPSVIRK/Personal_Finance_Manager/backend/handlers"
	"github.com/GPSVIRK/Personal_Finance_Manager/backend/internal/database"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"

	_ "github.com/lib/pq"
)



func main(){
	err := godotenv.Load()
	if err != nil{
		log.Fatalf("ERROR loading .env: %v",err)
	}

	portString := ":"+os.Getenv("PORT")
	db_url := os.Getenv("DB_URL")

	conn, err := sql.Open("postgres", db_url)
	defer conn.Close()

	if err!=nil{
		log.Fatal(err)
	}
	if err=conn.Ping(); err!=nil{
		log.Fatal(err)
	}

	db := database.New(conn)
	apiCfg := handlers.ApiConfig{
		DB: db,
	}

	router := chi.NewRouter()

	router.Use(middleware.Logger)
	
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"http://*", "https://*"},
		AllowedMethods: []string{"GET","PUT","POST","DELETE","OPTIONS"},
		AllowedHeaders: []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders: []string{"Link"},
		AllowCredentials: false,
		MaxAge: 300,
	}))

	router.Get("/test",handlers.Test)

	router.Get("/expenses",apiCfg.GetExpenses)
	router.Post("/expenses",apiCfg.PostExpenses)
	router.Patch("/expenses",apiCfg.PatchExpenses)
	router.Delete("/expenses",apiCfg.DeleteExpense)

	fmt.Printf("db open and server on %v\n",portString)
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
