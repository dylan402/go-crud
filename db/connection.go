package db

import (
	"database/sql"
	"fmt"
	"go-crud/configs"

	_ "github.com/lib/pq"
)

func OpenConnection() (*sql.DB, error) {
	conf := configs.GetDbConfig()

	stringConnection := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", conf.Host, conf.Port, conf.User, conf.Pass, conf.Db)

	conn, err := sql.Open("postgres", stringConnection)
	if err != nil {
		panic(err)
	}

	err = conn.Ping()

	return conn, err
}
