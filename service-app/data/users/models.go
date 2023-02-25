package users

import (
	"github.com/lib/pq"
	"time"
)

type User struct {
	ID           string         `json:"id"`
	Name         string         `json:"name"`
	Email        string         `json:"email"`
	Roles        pq.StringArray `json:"roles"`
	PasswordHash []byte         `json:"-"`
	DateCreated  time.Time      `json:"date_created"`
	DateUpdated  time.Time      `json:"date_updated"`
}

type NewUser struct {
	Name     string   `json:"name"`
	Email    string   `json:"email"`
	Roles    []string `json:"roles"`
	Password string   `json:"password"`
}

type ShirtInventory struct {
	ID          string    `json:"id"`
	UserId      string    `json:"user_id"`
	ItemName    string    `json:"item_name"`
	Quantity    int       `json:"quantity"`
	DateCreated time.Time `json:"date_created"`
	DateUpdated time.Time `json:"date_updated"`
}

// NewShirtInventory contains information needed to create a ShirtInventory.
type NewShirtInventory struct {
	ItemName string `json:"item_name"`
	Quantity int    `json:"quantity"`
}
