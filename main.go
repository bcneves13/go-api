package main

import(
	"encoding/json"
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"fmt"
	// "database/sql"
)
import _ "github.com/go-sql-driver/mysql"

//funcao principal
func main() {
	router := mux.NewRouter()
	// db, err := sql.Open("mysql", "root:root@/golang")
	// if err == nil {
	// 	nameRows, nameQueryErr := db.Query("SELECT * FROM nomes;")
	// 	if nameQueryErr == nil {
	// 		for nameRows.Next() {

	// 			fmt.Println("+v", nameRows.Scan("ID","Firstname", "Lastname"))
	// 		}
	// 		// if errRow == nil {
	// 		// }
	// 	}
	// }
	router.HandleFunc("/contato", GetPeople).Methods("GET")
	router.HandleFunc("/contato/{id}", GetPerson).Methods("GET")
	router.HandleFunc("/contato/{id}", CreatePerson).Methods("POST")
	router.HandleFunc("/contato/{id}", DeletePerson).Methods("DELETE")
	
	port := ":8000"
	fmt.Println("Server running in port:", port)
	log.Fatal(http.ListenAndServe(port, router))

}

type Address struct {
    City  string
    State string
}

type Person struct {
    ID        string
    Firstname string
    Lastname  string
    Address	  *Address
}

var people = []Person{
	Person{ID: "1", Firstname: "John", Lastname: "Doe", Address: &Address{City: "City X", State: "State X"}},
	Person{ID: "2", Firstname: "Koko", Lastname: "Doe", Address: &Address{City: "City Z", State: "State Y"}},
	Person{ID: "3", Firstname: "Francis", Lastname: "Sunday"},
}

// people = append(people, Person{ID: "1", Firstname: "John", Lastname: "Doe", Address: &Address{City: "City X", State: "State X"}})
// people = append(people, Person{ID: "2", Firstname: "Koko", Lastname: "Doe", Address: &Address{City: "City Z", State: "State Y"}})
// append(people, Person{ID: "3", Firstname: "Francis", Lastname: "Sunday"})

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "http://localhost:6010")
}

func GetPeople(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)

	json.NewEncoder(w).Encode(people)
	fmt.Println("pegou people");
}
func GetPerson(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for _, item := range people {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Person{})
}
func CreatePerson(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r);
	var person Person

	_ = json.NewDecoder(r.Body).Decode(&person) //blank identifier to execute a body decode to the variable  person

	person.ID = params["id"]
	people = append(people, person)
	json.NewEncoder(w).Encode(people)
}
func DeletePerson(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for index, item := range people {
		if item.ID == params["id"] {
			people = append(people[:index], people[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(people)
}
