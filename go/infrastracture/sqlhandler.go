package infrastracture

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

type SqlHandler struct {
	DB *sql.DB
}

func ConnectSql() *SqlHandler {
	dbConf := fmt.Sprintf("%s:%s@tcp(db:3306)/%s?charset=utf8&parseTime=true", os.Getenv("MYSQL_USER"), os.Getenv("MYSQL_PASSWORD"), os.Getenv("MYSQL_DATABASE"))

	db, err := sql.Open("mysql", dbConf)
	if err != nil {
		log.Fatal(err)
	}

	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}

	return &SqlHandler{db}
}

func (handler *SqlHandler) Execute(query string, args ...interface{}) (sql.Result, error) {
	type SqlResult struct {
		Result sql.Result
	}

	result, err := handler.DB.Exec(query, args...)
	res := SqlResult{
		Result: result,
	}
	if err != nil {
		return res.Result, err
	}

	return res.Result, nil
}
