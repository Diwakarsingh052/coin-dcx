package middlewares

import (
	"log"
	"service-app/auth"
)

type Mid struct {
	log *log.Logger
	a   *auth.Auth
}

func NewMid(log *log.Logger, a *auth.Auth) Mid {
	return Mid{
		log: log,
		a:   a,
	}
}
