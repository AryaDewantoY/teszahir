package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)


type User struct {
	ID string `json:"id"`
	Name string `json:"name"`
	Gender string `json:"gender"`
	Phone string `json:"phone"`
	Email string `json:"email"`
	Created_at string `json:"created_at"`
	Updated_at string `json:"updated_at"`

}

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Bismillah Tes Backend Zahir")
}

func allUsers(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(Users)
}

func singleUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	for _, user := range Users {
		if user.ID == id {
			json.NewEncoder(w).Encode(user)
		}
	}
}

func createUser(w http.ResponseWriter, r *http.Request) {
	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	Users = append(Users, user)
	json.NewEncoder(w).Encode(Users)
}

func updateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	params := mux.Vars(r)
	id := params["id"]

	var user User

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	for i, p := range Users {
		if p.ID == id {
			Users[i].ID = user.ID
			Users[i].Name = user.Name
			Users[i].Gender = user.Gender
			Users[i].Phone = user.Phone
			Users[i].Email = user.Email
			Users[i].Created_at = user.Created_at
			Users[i].Updated_at = user.Updated_at
			json.NewEncoder(w).Encode(Users[1])
			return
		}
	}
}

func deleteUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	params := mux.Vars(r)
	id := params["id"]

	for i, p := range Users {
		if p.ID == id {
			Users = append(Users[:i], Users[i+1:]...)
			json.NewEncoder(w).Encode(p)
			return
		}
	}
}

func handleRequest() {
	r :=  mux.NewRouter().StrictSlash(true)

	r.HandleFunc("/", home)
	r.HandleFunc("/user", allUsers).Methods("GET")
	r.HandleFunc("/user/{id}", singleUser).Methods("GET")
	r.HandleFunc("/user:", createUser).Methods("POST")
	r.HandleFunc("/user/{id}", updateUser).Methods("PUT")
	r.HandleFunc("/user/{id}", deleteUser).Methods("DELETE")

	fmt.Println("Rest API running!")
	log.Fatal(http.ListenAndServe(":8000", r))
}


func main() {
	Users = []User{
		User{ID: "1", Name: "Fulanbinfulan", Gender: "Male", Phone: "62813456789", Email: "fulan@gmail.com", Created_at: "2022-08-17T08:12:48+07:00", Updated_at: "2022-08-17T08:12:48+07:00"},
		User{ID: "2", Name: "Fulanibinfulani", Gender: "Female", Phone: "628134567234", Email: "fulani@gmail.com", Created_at: "2022-08-17T08:12:48+07:00", Updated_at: "2022-08-17T08:12:48+07:00"},
	}
	handleRequest()
}

var Users []User