package db

import (
	"database/sql"
	"fmt"

	// _ impede do compilador reclamar que a biblioteca não está sendo usada
	_ "github.com/lib/pq"
)

const (
	// host     = "go_db"
	host     = "localhost"
	port     = 5432
	dbname   = "postgres"
	user     = "postgres"
	password = "1234"
)

func ConnectDB() (*sql.DB, error) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}

	fmt.Println("Conectado ao banco de dados" + dbname)

	return db, nil
}
