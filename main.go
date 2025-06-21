//package main
//
//import (
//	"database/sql"
//	"fmt"
//	_ "github.com/lib/pq"
//)
//
//func main() {
//	connStr := "postgres://user:password@localhost/dbname?sslmode=disable"
//	db, err := sql.Open("postgres", connStr)
//	if err != nil {
//		panic(err)
//	}
//	defer db.Close()
//
//	err = db.Ping()
//	if err != nil {
//		panic(err)
//	}
//
//	rows, err := db.Query("select id, name, description, content, date_create  from news")
//	if err != nil {
//		fmt.Println("error in returning result")
//	}
//	defer rows.Close()
//	news := []new{}
//
//	for rows.Next() {
//		p := new{}
//		err := rows.Scan(&p.id, &p.name, &p.description, &p.content, &p.date_create)
//		if err != nil {
//			fmt.Println(err)
//			continue
//		}
//		news = append(news, p)
//	}
//	fmt.Println(news)
//	fmt.Println("Successfully connected to the database!")
//	fmt.Println("Successfully connected to the database333!")
//}

package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
)

type Response struct {
	Message string `json:"message"`
}

//type User struct {
//	ID    string `json:"id"`
//	Name  string `json:"name"`
//	Email string `json:"email"`
//}
//
//var users = []User{}
//
//func homeHandler(w http.ResponseWriter, r *http.Request) {
//	w.Header().Set("Content-Type", "application/json")
//	json.NewEncoder(w).Encode(Response{Message: "Добро пожаловать в API"})
//}
//
//func createUserHandler(w http.ResponseWriter, r *http.Request) {
//	var user User
//	json.NewDecoder(r.Body).Decode(&user)
//	user.ID = fmt.Sprintf("%d", len(users)+1)
//	users = append(users, user)
//
//	w.Header().Set("Content-Type", "application/json")
//	w.WriteHeader(http.StatusCreated)
//	json.NewEncoder(w).Encode(user)
//}
//
//func getUsersHandler(w http.ResponseWriter, r *http.Request) {
//	w.Header().Set("Content-Type", "application/json")
//	json.NewEncoder(w).Encode(users)
//}
//
//func updateUserHandler(w http.ResponseWriter, r *http.Request) {
//	params := mux.Vars(r)
//	id := params["id"]
//
//	for i, user := range users {
//		if user.ID == id {
//			json.NewDecoder(r.Body).Decode(&users[i])
//			w.Header().Set("Content-Type", "application/json")
//			json.NewEncoder(w).Encode(users[i])
//			return
//		}
//	}
//
//	http.Error(w, "Пользователь не найден", http.StatusNotFound)
//}
//
//func deleteUserHandler(w http.ResponseWriter, r *http.Request) {
//	params := mux.Vars(r)
//	id := params["id"]
//
//	for i, user := range users {
//		if user.ID == id {
//			users = append(users[:i], users[i+1:]...)
//			w.WriteHeader(http.StatusNoContent)
//			return
//		}
//	}
//
//	http.Error(w, "Пользователь не найден", http.StatusNotFound)
//}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", homeHandler).Methods("GET")
	r.HandleFunc("/users", createUserHandler).Methods("POST")
	r.HandleFunc("/users", getUsersHandler).Methods("GET")
	r.HandleFunc("/users/{id}", updateUserHandler).Methods("PUT")
	r.HandleFunc("/users/{id}", getUserHandler).Methods("GET")
	r.HandleFunc("/users/{id}", deleteUserHandler).Methods("DELETE")
	r.HandleFunc("/images/{id}", imageHandler).Methods("GET")

	//handler := http.HandlerFunc(handleRequest)
	//http.Handle("/image", handler)

	fmt.Println("Server is running on port 8081")
	http.ListenAndServe(":8081", r)
}

func imageHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]
	buf, err := ioutil.ReadFile("upload/" + id)

	if err != nil {

		log.Fatal(err)
	}

	w.Header().Set("Content-Type", "image/jpg")
	w.Write(buf)
}
