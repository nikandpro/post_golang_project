package handlers

import (
	"backend/basic/pkg/models"
	"backend/basic/pkg/vars"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

//CRUD


func GetPosts(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open(vars.DBSQL, vars.DBConn+vars.DBName)
	if err != nil {
		panic(err)
	}

	defer db.Close()

	res, err := db.Query("SELECT * FROM `articles`")
	if err != nil {
		panic(err)
	}


	var posts = []models.Article{}
	for res.Next() {
		var post models.Article
		err = res.Scan(&post.Id, &post.Title, &post.Anons, &post.FullText)
		if err != nil {
			panic(err)
		}
		posts = append(posts, post)
	}

	json_data, err := json.Marshal(posts)
	if err != nil {
        panic(err)
    }
	// fmt.Println(users)
	w.Write(json_data)
}

func GetPost(w http.ResponseWriter, r *http.Request) {
	varsURL := mux.Vars(r)
	
	db, err := sql.Open(vars.DBSQL, vars.DBConn+vars.DBName)
	if err != nil {
		panic(err)
	}

	defer db.Close()

	res, err := db.Query(fmt.Sprintf("SELECT * FROM `articles` WHERE `id` = '%s'", varsURL["id"]))
	if err != nil {
		panic(err)
	}

	var postId = models.Article{}
	for res.Next() {
		var post models.Article
		err = res.Scan(&post.Id, &post.Title, &post.Anons, &post.FullText)
		if err != nil {
			panic(err)
		}
		postId = post
	}

	json_data, err := json.Marshal(postId)
	if err != nil {
        panic(err)
    }
	w.Write(json_data)
}

func CreatePost(w http.ResponseWriter, r *http.Request) {

	var post = models.Article{}
	err := json.NewDecoder(r.Body).Decode(&post)


	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		panic(err)
	}

	db, err := sql.Open(vars.DBSQL, vars.DBConn+vars.DBName)
	if err != nil {
		panic(err)
	}

	defer db.Close()

	// Запись данных

	insert, err := db.Query("INSERT INTO articles (title, anons, full_text) VALUES ( ?, ?, ?)", post.Anons, post.Anons, post.FullText)

	if err != nil {
		panic(err)
	}
	defer insert.Close()
}

func UpdatePost(w http.ResponseWriter, r *http.Request) {
	varsURL := mux.Vars(r)
	id := varsURL["id"]
	
	db, err := sql.Open(vars.DBSQL, vars.DBConn+vars.DBName)
	if err != nil {
		panic(err)
	}

	defer db.Close()

	var post = models.Article{}
	err = json.NewDecoder(r.Body).Decode(&post)

	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		panic(err)
	}

	_, err = db.Query("UPDATE articles SET title = ?, anons = ?, full_text = ? WHERE id = ?", post.Title, post.Anons, post.FullText, id)
	if err != nil {
		fmt.Println("Update err")
		panic(err)
	}

		
}

func DeletePost(w http.ResponseWriter, r *http.Request) {
	varsURL := mux.Vars(r)
	id := varsURL["id"]
	
	db, err := sql.Open(vars.DBSQL, vars.DBConn+vars.DBName)
	if err != nil {
		panic(err)
	}

	defer db.Close()

	_, err = db.Query("DELETE FROM articles WHERE id = ?",  id)
	if err != nil {
		fmt.Println("Delete err")
		panic(err)
	}
}