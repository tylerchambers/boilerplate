package app

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt"

	"github.com/tylerchambers/boilerplate/models"
)

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Claims struct {
	UserID string `json:"UserID"`
	jwt.StandardClaims
}

func (s *Server) loginHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req LoginRequest

		json.NewDecoder(r.Body).Decode(&req)

		var user models.User
		s.db.First(&user, "email = ?", req.Email)
		if user.CheckPassword(req.Password) {
			expirationTime := time.Now().Add(5 * time.Minute)
			// Create the JWT claims, which includes the username and expiry time
			claims := &Claims{
				UserID: user.ID.String(),
				StandardClaims: jwt.StandardClaims{
					// In JWT, the expiry time is expressed as unix milliseconds
					ExpiresAt: expirationTime.Unix(),
				},
			}

			token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
			// Create the JWT string
			tokenString, err := token.SignedString(s.secretKey)
			if err != nil {
				// If there is an error in creating the JWT return an internal server error
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			// Finally, we set the client cookie for "token" as the JWT we just generated
			// we also set an expiry time which is the same as the token itself
			http.SetCookie(w, &http.Cookie{
				Name:    "token",
				Value:   tokenString,
				Expires: expirationTime,
			})
			w.Write([]byte("Login successful!"))
		} else {
			http.Error(w, "login failed", http.StatusBadRequest)
		}
	}
}
