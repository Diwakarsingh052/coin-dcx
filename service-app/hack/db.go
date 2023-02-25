package main

import (
	"context"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"time"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "diwakar"
	password = "root"
	dbname   = "postgres"
)

var db *sql.DB

func main() {

	var err error
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err = sql.Open("postgres", psqlInfo)

	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatalln(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	//Insert(ctx)
	//Insert2(ctx)
	QueryMultipleRecords(ctx)
}

func Insert(ctx context.Context) {
	sqlStatement := `INSERT INTO users (age, email, first_name,last_name)
					VALUES ($1, $2, $3, $4)`

	//exec func is used to just exec the query, it doesn't return any value from the DB
	result, err := db.ExecContext(ctx, sqlStatement, 33, "xyz@email.com", "xyz", "abc")
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(result.LastInsertId())
	fmt.Println(result.RowsAffected())
}

func Insert2(ctx context.Context) {

	sqlStatement := `INSERT INTO users (age, email, first_name,last_name)
					VALUES ($1, $2, $3, $4)
					RETURNING id,email
					`
	var (
		id    int
		email string
	)
	//QueryRowContext returns one row at a time as a result
	row := db.QueryRowContext(ctx, sqlStatement, 34, "abc11@email.com", "john", "don")
	err := row.Scan(&id, &email)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(id, email)

}

func QueryMultipleRecords(ctx context.Context) {
	sqlStatement := `select id,email FROM users`

	//QueryContext is used when we are expecting multiple rows back
	rows, err := db.QueryContext(ctx, sqlStatement)
	if err != nil {
		log.Fatalln(err)
	}
	defer rows.Close()

	for rows.Next() {
		var (
			id    int
			email string
		)
		err = rows.Scan(&id, &email)
		if err != nil {
			log.Println(err)
		}
		fmt.Println(id, email)

	}

}
