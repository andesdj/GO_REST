package app

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"net/http"
)

type Customer struct {
	Name    string `json:"first_name" xml:"name"`
	City    string `json:"city" xml:"city"`
	Zipcode string `json:"zip" xml:"zipcode"`
}

func greet(w http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(w, "Hello World!!")
}

func getAllCustomers(w http.ResponseWriter, r *http.Request) {
	customers := []Customer{
		{Name: "ANDES", City: "Bogota", Zipcode: "1234"},
		{Name: "SOL", City: "Cali", Zipcode: "5678"},
		{Name: "KATA", City: "Cartagena", Zipcode: "2580"},
		{Name: "Noralba", City: "jamundi", Zipcode: "1232131231"},
	}

	// IF is XML or JSON Request
	if r.Header.Get("Content-Type") == "application/xml" {
		w.Header().Add("Content-Type", "application/xml")
		xml.NewEncoder(w).Encode(customers)
	} else {
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(customers)

	}

}
