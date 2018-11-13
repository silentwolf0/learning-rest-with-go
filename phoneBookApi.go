package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Conducts struct {
	ID         string   `json:"id,omitempty"`
	FirstName  string   `json:"firstname,omitempty"`
	SecondName string   `json:"secondname,omitempty"`
	Address    *Address `json:"address,omitempty"`
}

type Address struct {
	City   string `json:"city,omitempty"`
	Street string `json:"street,omitempty"`
}

var cond []Conducts

func main() {
	myRouter := mux.NewRouter()
	cond = append(cond, Conducts{ID: "1", FirstName: "Cetric", SecondName: "Okola",
		Address: &Address{City: "Nairobi", Street: "Tom Mboya"}})
	cond = append(cond, Conducts{ID: "2", FirstName: "Elvis", SecondName: "Mutende",
		Address: &Address{City: "Kisumu", Street: "Raila"}})

	myRouter.HandleFunc("/cond", GetConducts).Methods("GET")
	myRouter.HandleFunc("/cond/{ID}", GetConduct).Methods("GET")
	myRouter.HandleFunc("/cond/{ID}", CreateConduct).Methods("POST")
	myRouter.HandleFunc("/cond/{ID}", DeleteConduct).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8080", myRouter))
}

func GetConducts(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(cond)

}
func GetConduct(w http.ResponseWriter, r *http.Request) {
	param := mux.Vars(r)
	for _, item := range cond {
		if item.ID == param["ID"] {
			json.NewEncoder(w).Encode(cond)
			return
		}
	}
	json.NewEncoder(w).Encode(&Conducts{})

}

func CreateConduct(w http.ResponseWriter, r *http.Request) {
	param := mux.Vars(r)
	var con Conducts
	_ = json.NewDecoder(r.Body).Decode(&cond)
	con.ID = param["id"]
	cond = append(cond, con)
	json.NewEncoder(w).Encode(cond)
}

func DeleteConduct(w http.ResponseWriter, r *http.Request) {
	param := mux.Vars(r)
	for index, item := range cond {
		if item.ID == param["id"] {
			cond = append(cond[:index], cond[index+1:]...) //deletes a given conduct
			break
		}
		json.NewEncoder(w).Encode(cond)
	}

}
