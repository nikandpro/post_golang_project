package database

import (
	"backend/basic/pkg/vars"
	"database/sql"
	"fmt"
)

// Создаем БД
func CreateDB(name string) {
	// CREATE DATABASE IF NOT EXISTS
	db, err := sql.Open("mysql", vars.DBConn+vars.DBName)
	if err != nil {
		panic(err)
	}
	defer db.Close() // закрываем подключение к БД

	// Делаем запрос
	if _, err = db.Exec(fmt.Sprintf(`CREATE DATABASE IF NOT EXISTS %s`, name)); err != nil {
		panic(err)
	}
}

func InitTables() {
	db, err := sql.Open("mysql", vars.DBConn+vars.DBName)
	if err != nil {
		panic(err)
	}
	defer db.Close() // закрываем подключение к БД

	// Делаем запрос
	if _, err = db.Exec(fmt.Sprint(`
	CREATE TABLE IF NOT EXISTS users (
		id INT NOT NULL AUTO_INCREMENT,
		name VARCHAR(20) NOT NULL DEFAULT '',
		age INT NOT NULL DEFAULT '',
		password VARCHAR(255) NOT NULL DEFAULT ''
	  );`)); err != nil {
		panic(err)
	}

	if _, err = db.Exec(fmt.Sprint(`
	CREATE TABLE IF NOT EXISTS articles (
		id INT NOT NULL AUTO_INCREMENT,
		title VARCHAR(100) NOT NULL DEFAULT '',
		anons VARCHAR(300) NOT NULL DEFAULT '',
		full_text text NOT NULL DEFAULT ''
	  );`)); err != nil {
		panic(err)
	}
}