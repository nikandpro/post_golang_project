package handlers

import (
	"database/sql"
	"fmt"
	"html/template"
	"net/http"

	"backend/basic/pkg/models"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

// var user = models.User{1, "Nikita", 22, "123"}
var posts = []models.Article{}
var showPost = models.Article{}

func Index(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("templates/index.html", "templates/header.html", "templates/footer.html")

	if err != nil {
		fmt.Fprintf(w, err.Error())
	}

	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/posts")
	if err != nil {
		panic(err)
	}

	defer db.Close()

	res, err := db.Query("SELECT * FROM `articles`")
	if err != nil {
		panic(err)
	}

	posts = []models.Article{}
	for res.Next() {
		var post models.Article
		err = res.Scan(&post.Id, &post.Title, &post.Anons, &post.FullText)
		if err != nil {
			panic(err)
		}
		posts = append(posts, post)
	}

	t.ExecuteTemplate(w, "index", posts)
}

func Create(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("templates/create.html", "templates/header.html", "templates/footer.html")

	if err != nil {
		fmt.Fprintf(w, err.Error())
	}

	t.ExecuteTemplate(w, "create", nil)
}

func Save_article(w http.ResponseWriter, r *http.Request) {
	title := r.FormValue("title")
	anons := r.FormValue("anons")
	full_text := r.FormValue("full_text")

	if title == "" || anons == "" || full_text == "" {
		fmt.Fprintf(w, "Не все данные заполные")
	} else {
		db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/posts")
		if err != nil {
			panic(err)
		}

		defer db.Close()

		// Запись данных

		insert, err := db.Query(fmt.Sprintf("INSERT INTO `articles` (`title`, `anons`, `full_text`) VALUES('%s','%s', '%s')", title, anons, full_text))
		if err != nil {
			panic(err)
		}
		defer insert.Close()

		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}

func Sign(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("templates/sign.html", "templates/header.html", "templates/footer.html")

	if err != nil {
		fmt.Fprintf(w, err.Error())
	}

	t.ExecuteTemplate(w, "sign", nil)
}

func Show_post(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	t, err := template.ParseFiles("templates/show.html", "templates/header.html", "templates/footer.html")

	if err != nil {
		fmt.Fprintf(w, err.Error())
	}

	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/posts")
	if err != nil {
		panic(err)
	}

	defer db.Close()

	res, err := db.Query(fmt.Sprintf("SELECT * FROM `articles` WHERE `id` = '%s'", vars["id"]))
	if err != nil {
		panic(err)
	}

	showPost = models.Article{}
	for res.Next() {
		var post models.Article
		err = res.Scan(&post.Id, &post.Title, &post.Anons, &post.FullText)
		if err != nil {
			panic(err)
		}
		showPost = post
	}

	t.ExecuteTemplate(w, "show", showPost)
}

func Sign_user(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("name")
	age := r.FormValue("age")
	password := r.FormValue("password")

	if name == "" || age == "" || password == "" {
		fmt.Fprintf(w, "Не все данные заполные")
	} else {
		db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/posts")
		if err != nil {
			panic(err)
		}

		defer db.Close()

		// Запись данных

		insert, err := db.Query(fmt.Sprintf("INSERT INTO `users` (`name`, `age`, `password`) VALUES('%s','%s', '%s')", name, age, password))
		if err != nil {
			panic(err)
		}
		defer insert.Close()

		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}

// func IndexUser(w http.ResponseWriter, r *http.Request) {
// 	t, err := template.ParseFiles("templates/show.html", "templates/header.html", "templates/footer.html")

// 	if err != nil {
// 		fmt.Fprintf(w, err.Error())
// 	}

// 	t.ExecuteTemplate(w, "header", user)
// }




