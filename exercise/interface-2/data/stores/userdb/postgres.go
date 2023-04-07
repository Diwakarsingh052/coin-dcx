package userdb

import (
	"coin-dcx/exercise/interface-2/data"
	"context"
	"database/sql"
	"fmt"
)

type Conn struct {
	db *sql.DB
}

func NewConn(db *sql.DB) *Conn {
	return &Conn{db: db}
}

func (p *Conn) Create(ctx context.Context, usr data.User) error {
	fmt.Println("adding to postgres", usr)
	return nil
}
func (p *Conn) Update(ctx context.Context, usr data.User) error {
	fmt.Println("updating in postgres", usr)
	return nil
}
func (p *Conn) Delete(usr data.User) error {
	fmt.Println("deleting in postgres", usr)
	return nil
}
