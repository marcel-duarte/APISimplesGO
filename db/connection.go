package db

import (
	"apisimples/config"
	"database/sql"
	"fmt"
)

func OpenConnection() (*sql.DB, error) {
	conf := config.GetDB()

	sc := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		conf.Host, conf.Port, conf.User, conf.Pass, conf.Database)

	conn, err := sql.Open("postgres", sc)

	if err != nil {
		panic(err) // nao deve ser usado em produção - apenas pra teste
	}
	err = conn.Ping()

	return conn, err

}
