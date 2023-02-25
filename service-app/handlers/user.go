package handlers

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"service-app/auth"
	"service-app/data/users"
	"service-app/web"
)

type handler struct {
	*users.Service
	*auth.Auth
}

func (h *handler) SignUp(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	v, ok := ctx.Value(web.KeyValue).(*web.Values)
	if !ok {
		return fmt.Errorf("web.Values missing from the context")
	}

	var nu users.NewUser

	//decode put the json data in struct by converting it
	err := web.Decode(r, &nu)

	if err != nil {
		return fmt.Errorf("%w", err)
	}

	//creating the user in the db
	usr, err := h.Create(ctx, nu, v.Now)
	if err != nil {
		return fmt.Errorf("user signup problem: %w", err)
	}

	return web.Respond(ctx, w, usr, http.StatusOK)

}

func (h *handler) Login(ctx context.Context, w http.ResponseWriter, r *http.Request) error {

	v, ok := ctx.Value(web.KeyValue).(*web.Values)
	if !ok {
		return fmt.Errorf("web.Values missing from the context")
	}

	var login struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	err := web.Decode(r, &login)
	if err != nil {
		return fmt.Errorf("%w", err)
	}

	claims, err := h.Authenticate(ctx, login.Email, login.Password, v.Now)

	if err != nil {
		return web.NewRequestError(errors.New("invalid email or password"), http.StatusUnauthorized)
	}

	var tkn struct {
		Token string `json:"token"`
	}

	tkn.Token, err = h.GenerateToken(claims)

	if err != nil {
		return fmt.Errorf("generating token %w", err)
	}

	h.SetCookie(w, tkn.Token)
	return web.Respond(ctx, w, tkn, http.StatusOK)

}

func (h *handler) SetCookie(w http.ResponseWriter, token string) {
	cookie := http.Cookie{
		Name:     "token",
		Value:    token,
		HttpOnly: true,
	}

	http.SetCookie(w, &cookie)
}
