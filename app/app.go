package app

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func Start() {
	// mux := http.NewServeMux()
	router := mux.NewRouter()

	//define routes
	router.HandleFunc("/greet", greet)
	router.HandleFunc("/customer", getAllCustomers)
	router.HandleFunc("/customer/{customer_id:[0-9]+}", getCustomer)

	//starting server
	log.Fatal(
		http.ListenAndServe("localhost:8000", router),
	)
}

func getCustomer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fmt.Fprint(w, vars["customer_id"])
}
