package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func setupRoutesForProductos(router *mux.Router) {
	// First enable CORS. If you don't need cors, comment the next line
	enableCORS(router)

	router.HandleFunc("/productos", func(w http.ResponseWriter, r *http.Request) {
		productos, err := getProductos()
		if err == nil {
			respondWithSuccess(productos, w)
		} else {
			respondWithError(err, w)
		}
	}).Methods(http.MethodGet)
	router.HandleFunc("/productos/{idProducto}", func(w http.ResponseWriter, r *http.Request) {
		idAsString := mux.Vars(r)["idProducto"]
		id, err := stringToInt64(idAsString)
		if err != nil {
			respondWithError(err, w)
			// We return, so we stop the function flow
			return
		}
		productos, err := getProductosById(id)
		if err != nil {
			respondWithError(err, w)
		} else {
			respondWithSuccess(productos, w)
		}
	}).Methods(http.MethodGet)

	router.HandleFunc("/productos", func(w http.ResponseWriter, r *http.Request) {
		// Declare a var so we can decode json into it
		var Productos Producto
		err := json.NewDecoder(r.Body).Decode(&Productos)
		if err != nil {
			respondWithError(err, w)
		} else {
			err := createProductos(Productos)
			if err != nil {
				respondWithError(err, w)
			} else {
				respondWithSuccess(true, w)
			}
		}
	}).Methods(http.MethodPost)

	router.HandleFunc("/productos", func(w http.ResponseWriter, r *http.Request) {

		var productos Producto
		err := json.NewDecoder(r.Body).Decode(&productos)
		if err != nil {
			respondWithError(err, w)
		} else {
			err := updateProductos(productos)
			if err != nil {
				respondWithError(err, w)
			} else {
				respondWithSuccess(true, w)
			}
		}
	}).Methods(http.MethodPut)
	router.HandleFunc("/productos/{idProducto}", func(w http.ResponseWriter, r *http.Request) {
		idAsString := mux.Vars(r)["idProducto"]
		id, err := stringToInt64(idAsString)
		if err != nil {
			respondWithError(err, w)

			return
		}
		err = deleteProductos(id)
		if err != nil {
			respondWithError(err, w)
		} else {
			respondWithSuccess(true, w)
		}
	}).Methods(http.MethodDelete)
}

func enableCORS(router *mux.Router) {
	router.PathPrefix("/").HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", AllowedCORSDomain)
	}).Methods(http.MethodOptions)
	router.Use(middlewareCors)
}
func middlewareCors(next http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, req *http.Request) {

			w.Header().Set("Access-Control-Allow-Origin", AllowedCORSDomain)
			w.Header().Set("Access-Control-Allow-Credentials", "true")
			w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
			w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

			next.ServeHTTP(w, req)
		})
}

func respondWithError(err error, w http.ResponseWriter) {
	w.WriteHeader(http.StatusInternalServerError)
	json.NewEncoder(w).Encode(err.Error())
}

func respondWithSuccess(data interface{}, w http.ResponseWriter) {

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(data)
}
