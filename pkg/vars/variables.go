package vars

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

var (
	VERSION = "1.1"
	_       = godotenv.Load()
	// .env vars
	PORT = os.Getenv("PORT")
	// Получение переменных окружения
	DBSQL = os.Getenv("DB_SQL")
	DBUser = os.Getenv("DB_USER")
	DBPass = os.Getenv("DB_PASS")
	DBName = os.Getenv("DB_NAME")
	DBHost = os.Getenv("DB_HOST")
	DBPort = os.Getenv("DB_PORT")

	DBConn = fmt.Sprintf("%s:%s@tcp(%s:%s)/", DBUser, DBPass, DBHost, DBPort)
)