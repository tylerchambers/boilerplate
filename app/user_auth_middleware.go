package app

import (
	"context"
	"fmt"
	"net/http"

	"github.com/golang-jwt/jwt"
)

type uidString string

// userAuthMiddleware checks if a user is authenticated.
func (s *Server) userAuthMiddleware(next http.Handler) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		c, err := r.Cookie("token")
		if err != nil {
			if err == http.ErrNoCookie {
				// If the cookie is not set, return an unauthorized status
				w.WriteHeader(http.StatusUnauthorized)
				return
			}
			// For any other type of error, return a bad request status
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		// Get the JWT string from the cookie
		tokenString := c.Value

		// Initialize a new instance of `Claims`
		// Parse takes the token string and a function for looking up the key. The latter is especially
		// useful if you use multiple keys for your application.  The standard is to use 'kid' in the
		// head of the token to identify which key to use, but the parsed token (head and claims) is provided
		// to the callback, providing flexibility.
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			// Don't forget to validate the alg is what you expect:
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return s.secretKey, nil
		})
		if err != nil {
			http.Error(w, "could not parse token", http.StatusInternalServerError)
		}
		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			r = r.WithContext(context.WithValue(r.Context(), uidString("uid"), claims["UserID"]))
			next.ServeHTTP(w, r)
		} else {
			http.Error(w, "invalid login", http.StatusInternalServerError)
		}

	})
}
