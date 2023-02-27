package userdb

import (
	"coin-dcx/excercise/data"
	"context"
	"database/sql"
	"fmt"
)

type Postgres struct {
	db *sql.DB
}

func NewPostgres(db *sql.DB) Postgres {
	return Postgres{db: db}
}

func (p *Postgres) Create(ctx context.Context, usr data.User) error {
	fmt.Println("adding to postgres", usr)
	return nil
}
func (p *Postgres) Update(ctx context.Context, usr data.User) error {
	fmt.Println("updating in postgres", usr)
	return nil
}
func (p *Postgres) Delete(ctx context.Context, usr data.User) error {
	fmt.Println("deleting in postgres", usr)
	return nil
}
