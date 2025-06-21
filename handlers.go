package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"net/http"
)

var users = []User{}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(Response{Message: "Добро пожаловать в API"})
}

func createUserHandler(w http.ResponseWriter, r *http.Request) {
	var user User
	json.NewDecoder(r.Body).Decode(&user)
	user.ID = fmt.Sprintf("%d", len(users)+1)
	users = append(users, user)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}

func getUsersHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	connStr := "postgres://user:password@localhost/dbname?sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	rows, err := db.Query("SELECT id, name, description, content, date_create, image FROM news ORDER BY id DESC")
	if err != nil {
		fmt.Println("error in returning result")
	}
	defer rows.Close()
	news := []newItem{}
	for rows.Next() {
		p := newItem{}
		err := rows.Scan(&p.Id, &p.Name, &p.Description, &p.Content, &p.Date_create, &p.Image)
		if err != nil {
			fmt.Println(err)
			continue
		}
		news = append(news, p)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(news)
}

func getUserHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	params := mux.Vars(r)
	id := params["id"]
	fmt.Println(id)
	connStr := "postgres://user:password@localhost/dbname?sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	row := db.QueryRow("SELECT id, name, description, content, date_create, image FROM news WHERE id = $1", id)
	news := newItem{}
	err2 := row.Scan(&news.Id, &news.Name, &news.Description, &news.Content, &news.Date_create, &news.Image)
	if err2 == sql.ErrNoRows {
		http.NotFound(w, r)
		return
	} else if err2 != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(news)

	//http.Error(w, "Пользователь не найден", http.StatusNotFound)
}

func contactsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	connStr := "postgres://user:password@localhost/dbname?sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	rows, err := db.Query("SELECT id, address FROM contacts ORDER BY id ASC")
	if err != nil {
		fmt.Println("error in returning result")
	}
	defer rows.Close()
	news := []contactItem{}
	for rows.Next() {
		p := contactItem{}
		err := rows.Scan(&p.Id, &p.Address)
		if err != nil {
			fmt.Println(err)
			continue
		}
		news = append(news, p)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(news)
}

func updateUserHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	for i, user := range users {
		if user.ID == id {
			json.NewDecoder(r.Body).Decode(&users[i])
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(users[i])
			return
		}
	}

	http.Error(w, "Пользователь не найден", http.StatusNotFound)
}

func deleteUserHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	for i, user := range users {
		if user.ID == id {
			users = append(users[:i], users[i+1:]...)
			w.WriteHeader(http.StatusNoContent)
			return
		}
	}

	http.Error(w, "Пользователь не найден", http.StatusNotFound)
}
