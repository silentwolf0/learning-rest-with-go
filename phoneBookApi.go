package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Contact struct {
	ID         string   `json:"id,omitempty"`
	FirstName  string   `json:"firstname,omitempty"`
	SecondName string   `json:"secondname,omitempty"`
	Address    *Address `json:"address,omitempty"`
}

type Address struct {
	City   string `json:"city,omitempty"`
	Street string `json:"street,omitempty"`
}

var conts []Contact

func main() {
	myRouter := mux.NewRouter()
	conts = append(conts, Contact{ID: "1", FirstName: "Cetric", SecondName: "Okola",
		Address: &Address{City: "Nairobi", Street: "Tom Mboya"}})
	conts = append(conts, Contact{ID: "2", FirstName: "Elvis", SecondName: "Mutende",
		Address: &Address{City: "Kisumu", Street: "Raila"}})

	fmt.Println(len(conts))

	myRouter.HandleFunc("/conts", GetContacts).Methods("GET")
	myRouter.HandleFunc("/conts/{ID}", GetContact).Methods("GET")
	myRouter.HandleFunc("/conts/{ID}", CreateContact).Methods("POST")
	myRouter.HandleFunc("/conts/{ID}", DeleteContact).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8080", myRouter))
}

func GetContacts(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(conts)

}
func GetContact(w http.ResponseWriter, r *http.Request) {

	param := mux.Vars(r)
	for _, item := range conts {
		if item.ID == param["ID"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}

	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte("404 - Contact Not Found"))

}

// CreateContact ...
func CreateContact(w http.ResponseWriter, r *http.Request) {
	param := mux.Vars(r)
	var con Contact
	
	if r.Body == nil {
		http.Error(w, "Please send a request body", 400)
		return
	}
	err := json.NewDecoder(r.Body).Decode(&con) //send json data to the server
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	con.ID = param["id"]
	conts = append(conts, con)
	json.NewEncoder(w).Encode(conts) //write json to the server
}

func DeleteContact(w http.ResponseWriter, r *http.Request) {
	param := mux.Vars(r)
	for index, item := range conts {
		if item.ID == param["id"] {
			conts = append(conts[:index], conts[index+1:]...) //deletes a given conduct
			break
		}
	}
	json.NewEncoder(w).Encode(conts)
}
