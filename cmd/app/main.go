package main

import (
	"net/http"
	"backend/basic/pkg/routes"
	_ "github.com/go-sql-driver/mysql"
)


func handleFunc() {

	http.Handle("/", routes.InitRoutes())
	http.ListenAndServe(":8081", nil)
}

func main() {
	handleFunc()
}
