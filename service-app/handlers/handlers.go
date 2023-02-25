package handlers

import (
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"
	"service-app/auth"
	"service-app/data/users"
	"service-app/middlewares"
	"service-app/web"
)

func Api(log *log.Logger, a *auth.Auth, s *users.Service) http.Handler {

	//r := chi.NewRouter()
	//r.MethodFunc(http.MethodGet, "/check", check)
	app := web.App{
		Mux: chi.NewRouter(),
	}
	h := handler{
		Service: s,
		Auth:    a,
	}

	m := middlewares.NewMid(log, a)
	app.HandleFunc(http.MethodGet, "/check", m.Logger(m.Error(m.Panic(m.Authenticate(m.Authorize(check, auth.RoleAdmin, auth.RoleUser))))))
	app.HandleFunc(http.MethodPost, "/signup", m.Logger(m.Error(m.Panic(h.SignUp))))
	app.HandleFunc(http.MethodPost, "/add", m.Logger(m.Error(m.Panic(m.AuthenticateCookie(m.Authorize(h.AddInventory, auth.RoleAdmin, auth.RoleUser))))))
	app.HandleFunc(http.MethodPost, "/view", m.Logger(m.Error(m.Panic(m.AuthenticateCookie(m.Authorize(h.ViewInventory, auth.RoleAdmin, auth.RoleUser))))))

	app.HandleFunc(http.MethodPost, "/login", m.Logger(m.Error(m.Panic(h.Login))))
	//app.HandleFunc(http.MethodPost, "/add", m.Logger(m.Error(m.Panic(m.Authenticate(m.Authorize(h.AddInventory, auth.RoleAdmin, auth.RoleUser))))))
	//app.HandleFunc(http.MethodPost, "/view", m.Logger(m.Error(m.Panic(m.Authenticate(m.Authorize(h.ViewInventory, auth.RoleAdmin, auth.RoleUser))))))
	return app
}
