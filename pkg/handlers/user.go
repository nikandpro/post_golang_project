package handlers

import (
	"backend/basic/pkg/models"
	"backend/basic/pkg/vars"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"golang.org/x/crypto/bcrypt"
)

//CRUD

func Login(w http.ResponseWriter, r *http.Request) {
	var user = models.User{}
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		panic(err)
	}
	db, err := sql.Open(vars.DBSQL, vars.DBConn+vars.DBName)
	if err != nil {
		panic(err)
	}

	defer db.Close()


	res, err := db.Query(fmt.Sprintf("SELECT * FROM `users` WHERE `name` = '%s' AND PASSWORD = '%s'", user.Name, user.Password))
	if err != nil {
		panic(err)
	}

	var userFind = models.User{}
	for res.Next() {
		var user models.User
		err = res.Scan(&user.Id, &user.Name, &user.Age, &user.Password)
		if err != nil {
			panic(err)
		}
		userFind = user
	}

	if userFind.Name != "" {
		json_data, err := json.Marshal(userFind)
		if err != nil {
        	panic(err)
    	}
		w.Write(json_data)
	} else {
		http.Error(w, "No found", http.StatusNotFound)
	}
	
	
	
}

func GetUsers(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open(vars.DBSQL, vars.DBConn+vars.DBName)
	if err != nil {
		panic(err)
	}

	defer db.Close()

	res, err := db.Query("SELECT * FROM `users`")
	if err != nil {
		panic(err)
	}

	var users = []models.User{}
	for res.Next() {
		var user models.User
		err = res.Scan(&user.Id, &user.Name, &user.Age, &user.Password)
		if err != nil {
			panic(err)
		}
		users = append(users, user)
	}

	json_data, err := json.Marshal(users)
	if err != nil {
        panic(err)
    }
	// fmt.Println(users)
	w.Write(json_data)
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	varsURL := mux.Vars(r)
	
	db, err := sql.Open(vars.DBSQL, vars.DBConn+vars.DBName)
	if err != nil {
		panic(err)
	}

	defer db.Close()

	res, err := db.Query(fmt.Sprintf("SELECT * FROM `users` WHERE `id` = '%s'", varsURL["id"]))
	if err != nil {
		panic(err)
	}

	var userId = models.User{}
	for res.Next() {
		var user models.User
		err = res.Scan(&user.Id, &user.Name, &user.Age, &user.Password)
		if err != nil {
			panic(err)
		}
		userId = user
	}

	json_data, err := json.Marshal(userId)
	if err != nil {
        panic(err)
    }
	w.Write(json_data)
}

func CreateUser(w http.ResponseWriter, r *http.Request) {

	var user = models.User{}
	err := json.NewDecoder(r.Body).Decode(&user)
	// fmt.Println(user.Id, user.Name, user.Age, user.Password)

	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		panic(err)
	}

	db, err := sql.Open(vars.DBSQL, vars.DBConn+vars.DBName)
	if err != nil {
		panic(err)
	}

	defer db.Close()

	password := []byte(user.Password)

	hashPas, _ := bcrypt.GenerateFromPassword(password, 10)

	// Запись данных

	insert, err := db.Query("INSERT INTO users (name, age, password) VALUES ( ?, ?, ?)", user.Name, user.Age, hashPas)
	// insert, err := db.Query(fmt.Sprintf("INSERT INTO `users` (`name`, `age`, `password`) VALUES('%s', '%s', '%s')", user.Name, user.Age, user.Password))
	if err != nil {
		panic(err)
	}
	defer insert.Close()
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	varsURL := mux.Vars(r)
	id := varsURL["id"]
	
	db, err := sql.Open(vars.DBSQL, vars.DBConn+vars.DBName)
	if err != nil {
		panic(err)
	}

	defer db.Close()

	var user = models.User{}
	err = json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		panic(err)
	}

	_, err = db.Query("UPDATE users SET name = ?, age = ?, password = ? WHERE id = ?", user.Name, user.Age, user.Password, id)
	if err != nil {
		fmt.Println("Updat err")
		panic(err)
	}

		
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	varsURL := mux.Vars(r)
	id := varsURL["id"]
	
	db, err := sql.Open(vars.DBSQL, vars.DBConn+vars.DBName)
	if err != nil {
		panic(err)
	}

	defer db.Close()

	_, err = db.Query("DELETE FROM users WHERE id = ?",  id)
	if err != nil {
		fmt.Println("Delete err")
		panic(err)
	}
}

