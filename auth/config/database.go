package config

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var schema = `
CREATE TABLE public.users (
	id serial NOT NULL,
	username varchar(128) NULL,
	"password" varchar(256) NULL,
	CONSTRAINT users_pkey PRIMARY KEY (id)
	);`


var DB *sqlx.DB

// ConnectDB to get all needed db connections for application
func ConnectDB() *sqlx.DB {
	DB = getDBConnection()

	return DB
}

func getDBConnection() *sqlx.DB {
	var dbConnectionStr string

	dbConnectionStr = fmt.Sprintf(
		"host=%s port=%d dbname=%s user=%s password=%s sslmode=disable",
		"localhost",
		5432,
		"grpc_auth",
		"postgres",
		"postgres",
	)
	

	db, err := sqlx.Open("postgres", dbConnectionStr)

	//db.MustExec(schema)
	
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	//TODO: experiment with correct values
	db.SetMaxIdleConns(1)
	db.SetMaxOpenConns(5)

	fmt.Println("Connected to DB")
	return db
}
