package postgres

import (
	_ "github.com/jackc/pgx/stdlib"
	"github.com/jmoiron/sqlx"
	"log"
)

type Database struct {
	db *sqlx.DB
}

func New(dsn string) (*Database, error) {
	//db, err := sqlx.Open("postgres", dsn)
	db, err := sqlx.Open("pgx", dsn)
	if err != nil {
		log.Println("DataBase NOT WORK")
		return nil, err
	}
	log.Println("DataBase Work")
	return &Database{
		db: db,
	}, nil
}
