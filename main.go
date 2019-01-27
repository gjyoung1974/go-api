package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"rest-api/model"
)

var people []model.Person

func GetPersonEndpoint(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for _, item := range people {
		if item.Firstname == params["Firstname"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&model.Person{})
}

func GetPeopleEndpoint(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(people)
}

func CreatePersonEndpoint(w http.ResponseWriter, r *http.Request) {
	db := model.DBMigrate()
	defer db.Close()

	var person model.Person
	_ = json.NewDecoder(r.Body).Decode(&person)
	people = append(people, person)
	//json.NewEncoder(w).Encode(people)
	db.Create(&person)

}
func DeletePersonEndpoint(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for index, item := range people {
		if item.Firstname == params["Firstname"] {
			people = append(people[:index], people[index+1:]...)
			break
		}
		json.NewEncoder(w).Encode(people)
	}
}

func main() {

	router := mux.NewRouter()

	router.HandleFunc("/people", GetPeopleEndpoint).Methods("GET")
	router.HandleFunc("/people/{id}", GetPersonEndpoint).Methods("GET")
	router.HandleFunc("/people/{id}", CreatePersonEndpoint).Methods("POST")
	router.HandleFunc("/people/{id}", DeletePersonEndpoint).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8000", router))
}

// curl --header "Content-Type: application/json" --request POST --data '{Firstname: "John", Lastname: "Doe", City: "City X", State: "State X"}' http://localhost:8000/people/1
