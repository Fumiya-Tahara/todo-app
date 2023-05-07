package infrastracture

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

type Storage struct {
	DB *sql.DB
}

// connect MySQL
func NewStorage() *Storage {
	dbConf := fmt.Sprintf("%s:%s@tcp(db:3306)/%s?charset=utf8&parseTime=true", os.Getenv("MYSQL_USER"), os.Getenv("MYSQL_PASSWORD"), os.Getenv("MYSQL_DATABASE"))

	fmt.Println(dbConf)
	db, err := sql.Open("mysql", dbConf)
	if err != nil {
		log.Fatal(err)
	}

	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}

	fmt.Println("connected!")

	return &Storage{db}
}

func (handler *Storage) Execute(query string, args ...interface{}) (sql.Result, error) {
	result, err := handler.DB.Exec(query, args...)
	// res := SqlResult{
	// 	Result: result,
	// }
	if err != nil {
		return result, err
	}

	return result, nil
}
