package main

import (
	"book-management/pkg/routes"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	router := mux.NewRouter()
	routes.RegisterBookStoreRoutes(router)
	http.Handle("/", router)
	log.Println("Server started at http://localhost:9010")
	log.Fatal(http.ListenAndServe("localhost:9010", router))
}
