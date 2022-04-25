package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func main() {
	// Ping database
	bd, err := getDB()
	if err != nil {
		log.Printf("Error con la base de datos" + err.Error())
		return
	} else {
		err = bd.Ping()
		if err != nil {
			log.Printf("Error al conectar con la base de datos, Por favor revisar sus credenciales: " + err.Error())
			return
		}
	}
	// Define routes
	router := mux.NewRouter()
	setupRoutesForProductos(router)
	// .. here you can define more routes
	// ...
	// for example setupRoutesForGenres(router)

	// Setup and start server
	port := ":8000"

	server := &http.Server{
		Handler: router,
		Addr:    port,
		// timeouts so the server never waits forever...
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Printf("Server started at %s", port)
	log.Fatal(server.ListenAndServe())
}
