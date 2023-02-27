package main

import (
	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
	"mysql-bookstore/pkg/routes"
	"net/http"
)

func main() {
	myRoutes := mux.NewRouter()
	routes.RegisterBookstoreRoutes(myRoutes)
	http.Handle("/", myRoutes)
	log.Fatal(http.ListenAndServe("localhost:9010", myRoutes))
}
