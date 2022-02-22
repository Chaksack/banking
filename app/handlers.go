package app

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"net/http"
)

type Customer struct {
	Name    string `json:"full_name"`
	City    string `json:"city"`
	Zipcode string `json:"zip_code"`
}

func greet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello World!")
}

func getAllCustomers(w http.ResponseWriter, r *http.Request) {
	customers := []Customer{
		{Name: "Keziah", City: "New Delhi", Zipcode: "2330"},
		{Name: "Chak", City: "New Delhi", Zipcode: "2330"},
	}

	if r.Header.Get("Content-Type") == "application/xml" {
		w.Header().Add(
			"Content-Type", "application/xml",
		)
		xml.NewEncoder(w).Encode(customers)
	} else {
		w.Header().Add(
			"Content-Type", "application/json",
		)
		json.NewEncoder(w).Encode(customers)
	}
}
