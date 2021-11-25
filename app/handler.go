package app

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type Customer struct {
	Name    string `json:"name" xml:"name`
	City    string `json:"city" xml:"city`
	ZipCode string `json:"zip_code" xml:"zip_code`
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "hello from homepage")
}

func getAllCustomer(w http.ResponseWriter, r *http.Request) {
	customers := []Customer{
		{Name: "majid", City: "kendari", ZipCode: "12345"},
		{Name: "majid", City: "kendari", ZipCode: "12345"},
		{Name: "majid", City: "kendari", ZipCode: "12345"},
	}

	if r.Header.Get("Content-Type") == "application/xml" {
		xml.NewEncoder(w).Encode(customers)
	} else {
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(customers)
	}

}

func getCustomers(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fmt.Fprint(w, vars["customer_id"])
}

func getPostCustomer(w http.ResponseWriter, r *http.Request) {

	fmt.Fprint(w, "Posst Intttt")
}
