package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
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

func getUsersHandler(w http.ResponseWriter, r *http.Request) {
	//w.Header().Set("Content-Type", "application/json")
	//json.NewEncoder(w).Encode(users)
	w.Header().Set("Access-Control-Allow-Origin", "*")
	//db, err := sql.Open("mysql", "usershop2:04051960@/dbshop2")
	//
	//if err != nil {
	//	panic(err)
	//}
	//defer db.Close()
	//
	//if err := db.Ping(); err != nil {
	//	log.Fatalln(err)
	//}
	//
	//rows, err := db.Query("select el.ID as ID, NAME, PREVIEW_TEXT, FILE_NAME, SUBDIR  from dbshop2.b_iblock_element as el left join dbshop2.b_file as file ON el.PREVIEW_PICTURE=file.ID WHERE IBLOCK_ID=1")
	//if err != nil {
	//	fmt.Println("error in returning result")
	//}
	//defer rows.Close()
	//news := []new{}
	//
	//for rows.Next() {
	//	p := new{}
	//	err := rows.Scan(&p.ID, &p.NAME, &p.PREVIEW_TEXT, &p.FILE_NAME, &p.SUBDIR)
	//	if err != nil {
	//		fmt.Println(err)
	//		continue
	//	}
	//	news = append(news, p)
	//}

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

	rows, err := db.Query("SELECT id, name, description, content, date_create, image FROM news")
	if err != nil {
		fmt.Println("error in returning result")
	}
	defer rows.Close()
	news := []newItem{}
	var tom = person{"Tom", 23}
	//bob := new{'id', 'name', 'description', 'content', 'date_create'}
	fmt.Println("success1")
	for rows.Next() {
		p := newItem{}
		err := rows.Scan(&p.Id, &p.Name, &p.Description, &p.Content, &p.Date_create, &p.Image)
		if err != nil {
			fmt.Println(err)
			continue
		}
		news = append(news, p)
	}

	//for _, p := range news {
	//	fmt.Println(p.id, p.name, p.description, p.content)
	//}
	//js := `{"status": "available", "environment": %q, "version": %q}`
	//arr := [4]int{3, 2, 5, 4}

	fmt.Println(news)
	//longerArr := []int{5, 7, 1, 2, 0}
	var myMap map[string]int

	// Инициализируем map
	myMap = make(map[string]int)
	myMap["апельсин"] = 10
	myMap["яблоко"] = 25
	fmt.Println(tom)
	w.Header().Set("Content-Type", "application/json")
	//w.Write([]byte(js))
	//fmt.Println(json.Marshal(news))
	//fmt.Println(json.NewEncoder(w).Encode(news))
	//json.NewEncoder(w).Encode(tom)
	json.NewEncoder(w).Encode(news)
}

func getUserHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	params := mux.Vars(r)
	id := params["id"]
	db, err := sql.Open("mysql", "usershop2:04051960@/dbshop2")

	if err != nil {
		panic(err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatalln(err)
	}

	rows, err := db.Query("select ID, NAME from dbshop2.b_iblock_element WHERE IBLOCK_ID=1 AND ID=?", id)
	if err != nil {
		fmt.Println("error in returning result")
	}
	defer rows.Close()
	news := []newItem{}

	for rows.Next() {
		p := newItem{}
		err := rows.Scan(&p.Id, &p.Name)
		if err != nil {
			fmt.Println(err)
			continue
		}

		news = append(news, p)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(news)

	//http.Error(w, "Пользователь не найден", http.StatusNotFound)
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
