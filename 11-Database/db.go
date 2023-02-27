package main

import (
	"context"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"time"
)

var db *sql.DB //nil

const (
	host     = "localhost"
	port     = 5432
	user     = "diwakar"
	password = "root"
	dbname   = "users"
)

type usr struct {
	name      string
	last_name string
}

func main() {
	var err error
	psqlInfo := fmt.Sprintf("host=%s port=%d data=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	db, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	//db.PingContext()
	if err := db.Ping(); err != nil {
		panic(err)
	}
	fmt.Println("connected")
	//Insert()
	//Insert2(ctx)
	//Update(ctx)
	//delete(ctx)
	//querySingleRecords(ctx)
	QueryMultipleRecords(ctx)
}

func Insert() {

	sqlStatement := `INSERT INTO users (age, email, first_name,last_name)
					VALUES ($1, $2, $3, $4)`
	var name = "abc"
	//exec query //
	res, err := db.Exec(sqlStatement, 32, "abc@email.com", name, "abc")
	//db.ExecContext()
	if err != nil {
		log.Println(err)
		return
	}

	fmt.Println(res.LastInsertId())

}

func Insert2(ctx context.Context) {

	sqlStatement := `INSERT INTO users (age, email, first_name,last_name)
					VALUES ($1, $2, $3, $4)
					RETURNING id,email
					`
	var name = "abc"
	var (
		id    int
		email string
	)

	//exec the query and return one row back
	err := db.QueryRowContext(ctx, sqlStatement, 32, "efg@email.com", name, "arora").Scan(&id, &email)
	//db.QueryRowContext()
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println(id, email)

}

func Update(ctx context.Context) {

	sqlStatement := `Update users 
                    Set last_name= $1
					 where id =$2;
`
	u := usr{
		name:      "ajay",
		last_name: "singh",
	}

	res, err := db.ExecContext(ctx, sqlStatement, u.last_name, 2)

	if err != nil {
		log.Println(err)
	}

	fmt.Println(res.RowsAffected())

}

func delete(ctx context.Context) {

	sqlStatement := `Delete FROM users
                    where id =$1;
`

	res, err := db.ExecContext(ctx, sqlStatement, 2)
	if err != nil {
		log.Println(err)
		return
	}

	fmt.Println(res.RowsAffected())

}

func querySingleRecords(ctx context.Context) {

	sqlStatement := `Select id, email FROM users where id = $1;`

	var (
		id    int
		email string
	)

	err := db.QueryRowContext(ctx, sqlStatement, 1).Scan(&id, &email)

	if err != nil {
		log.Println(err)

	}

	switch err {
	case sql.ErrNoRows:
		log.Println("no rows returned")
	case nil:
		fmt.Println(id, email)
	default:
		log.Println(err)

	}

}

func QueryMultipleRecords(ctx context.Context) {

	//exec the query // we expect the multiple rows back
	rows, err := db.QueryContext(ctx, "Select id, email FROM users LIMIT $1", 4)

	if err != nil {
		log.Println(err)
		return
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

	if err := rows.Err(); err != nil {
		log.Println(err)
		return
	}

}
