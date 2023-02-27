package user

import (
	"errors"
	"log"
)

// Service struct inject dependency in data package
type Service struct {
	db string
	l  *log.Logger
}

var ErrConn = errors.New("please provide a valid connection")

// NewService func is initializing the Config struct
func NewService(host, port, conn string, l *log.Logger) (*Service, error) {
	if host == "" || port == "" || conn == "" || l == nil {
		return nil, ErrConn
		// without pointer // Config{},ErrCon

	}

	db := host + port + conn

	//c := Service{db: db, l: l}

	return &Service{
		db: db,
		l:  l,
	}, nil

}

func (s *Service) AddToDb(name string, age int) {
	u := user{
		name: name,
		age:  age,
	}
	if s.db == "" {
		s.l.Println("cannot add we don't have a db connection")
	}

	s.l.Println("adding to db", u)

}

func (s *Service) Update(name string, age int) {
	u := user{
		name: name,
		age:  age,
	}

	if s.db == "" {
		s.l.Println("cannot add we don't have a db connection")
	}

	s.l.Println("updating in db", u)

}
