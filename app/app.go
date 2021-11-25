package app

import (
	"net/http"
	"hexagonal/data"

	"github.com/gorilla/mux"
)

func Start() {
	router := mux.NewRouter()

	router.HandleFunc("/", homePage).Methods("GET")
	router.HandleFunc("/customers", getAllCustomer).Methods("GET")
	router.HandleFunc("/customers", getPostCustomer).Methods("POST")
	router.HandleFunc("/customers/{customer_id:[0-9]+}", getCustomers).Methods("GET")
	http.ListenAndServe(data.Port(), router)
}
