package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/lib/pq"
	_ "github.com/lib/pq"
	//	"gorm.io/gorm"
	"log"
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

func getNewsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	key := r.FormValue("key")
	//var router = mux.NewRouter()

	//u, errRouter := router.Get("getNewsHandler").URL("key", key)
	//if errRouter != nil {
	//	http.Error(w, errRouter.Error(), 500)
	//	return
	//}

	fmt.Println(key)
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

func getNewHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	params := mux.Vars(r)
	id := params["id"]
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
	contacts := []contactItem{}
	contactIds := []string{}
	for rows.Next() {
		p := contactItem{}
		err := rows.Scan(&p.Id, &p.Address)
		if err != nil {
			fmt.Println(err)
			continue
		}

		if !contains(contactIds, p.Id) {
			contactIds = append(contactIds, p.Id)
		}

		contacts = append(contacts, p)
	}

	rows2, err2 := db.Query("SELECT id, addressid, value FROM phones WHERE addressid = ANY($1)", pq.Array(contactIds))
	if err2 != nil {
		fmt.Println("error in returning result")
	}
	defer rows2.Close()

	var myMap map[string][]string

	// Инициализируем map
	myMap = make(map[string][]string)

	for rows2.Next() {
		var id string
		var addressid string
		var value string
		if err := rows2.Scan(&id, &addressid, &value); err != nil {
			log.Fatal(err)
		}
		myMap[addressid] = append(myMap[addressid], value)
	}

	for i := len(contacts) - 1; i >= 0; i-- {
		contacts[i].Phone = myMap[contacts[i].Id]
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(contacts)
}

func blocksHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Access-Control-Allow-Origin", "*")
	params := mux.Vars(r)
	code := params["code"]
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

	rows, err := db.Query("SELECT id, content, code FROM blocks WHERE code = $1", code)
	if err != nil {
		fmt.Println("error in returning result")
	}
	defer rows.Close()
	blocks := []blockItem{}
	for rows.Next() {
		p := blockItem{}
		err := rows.Scan(&p.Id, &p.Content, &p.Code)
		if err != nil {
			fmt.Println(err)
			continue
		}

		blocks = append(blocks, p)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(blocks)
}

func contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}

	return false
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
