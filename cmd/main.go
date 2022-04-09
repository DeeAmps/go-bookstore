package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/DeeAmps/go-bookstore-api/pkg/db"
	"github.com/DeeAmps/go-bookstore-api/pkg/routes"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterBookRoutes(r)
	_, err := db.Connect()
	if err != nil {
		panic(err)
	}
	// log.Fatal(db.AutoMigrate(&models.Book{}))
	fmt.Println("Database Migration Completed")
	fmt.Println("Application listening on port 8000")
	log.Fatal(http.ListenAndServe(":8000", r))

}
