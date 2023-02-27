package middlewares

import (
	"context"
	"errors"
	"net/http"
	"service-app/auth"
	"service-app/web"
)

func (m *Mid) AuthenticateCookie(next web.HandlerFunc) web.HandlerFunc {
	return func(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
		cookie, err := r.Cookie("token")
		if err != nil {
			err := errors.New("please login first")
			//http.Redirect(w, r, "/home", http.StatusOK)

			return web.NewRequestError(err, http.StatusUnauthorized)
		}

		claims, err := m.a.ValidateToken(cookie.Value)
		if err != nil {
			return web.NewRequestError(err, http.StatusUnauthorized)
		}

		// putting the token in the context so we can see the values in the claims struct in the request life cycle //
		// specifically we will look for the subject field as it stores unique data id which will helpful to identify for whom this token was genrated
		ctx = context.WithValue(ctx, auth.Key, claims)
		return next(ctx, w, r)
	}

}
