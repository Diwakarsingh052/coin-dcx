package middlewares

import (
	"context"
	"errors"
	"net/http"
	"service-app/auth"
	"service-app/web"
	"strings"
)

var ErrForbidden = web.NewRequestError(
	errors.New("you are not authorized for that action"),
	http.StatusForbidden,
)

func (m *Mid) Authenticate(next web.HandlerFunc) web.HandlerFunc {
	return func(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
		// fetching data from Authorization header
		authHeader := r.Header.Get("Authorization")

		//token format :- bearer <token> // we trying to separate both strings with space
		parts := strings.Split(authHeader, " ") // parts would be slice

		if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
			err := errors.New("expected authorization header format: Bearer <token>")
			return web.NewRequestError(err, http.StatusUnauthorized)
		}

		claims, err := m.a.ValidateToken(parts[1])
		if err != nil {
			return web.NewRequestError(err, http.StatusUnauthorized)
		}

		// putting the token in the context so we can see the values in the claims struct in the request life cycle //
		// specifically we will look for the subject field
		//as it stores unique user id which will helpful to identify for whom this token was generated
		ctx = context.WithValue(ctx, auth.Key, claims)

		return next(ctx, w, r)

	}
}

func (m *Mid) Authorize(next web.HandlerFunc, requiredRoles ...string) web.HandlerFunc {
	return func(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
		claims, ok := ctx.Value(auth.Key).(auth.Claims)

		if !ok {
			return errors.New("claims missing from context: Authorize called without/before Authenticate")
		}

		ok = claims.HasRoles(requiredRoles...)
		if !ok {
			return ErrForbidden
		}
		return next(ctx, w, r)
	}

}
