package data

import (
	"context"
)

// Storer interface declares the behavior this package needs to perists and
// retrieve data.
type Storer interface {
	Create(ctx context.Context, usr User) error
	Update(ctx context.Context, usr User) error
	Delete(ctx context.Context, usr User) error
}

// Store manages the set of APIs for data access.
type Store struct {
	Storer
}

// NewStore constructs a core for data api access.
func NewStore(storer Storer) *Store {
	//initialize the store struct
}
