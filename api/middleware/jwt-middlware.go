package middlewares

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/Devil39/enigma/api/views"
	"github.com/dgrijalva/jwt-go"
)

//JwtMiddleware acts as a jwt validator middleware
func JwtMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		message := make(map[string]interface{})
		err := isValidToken(r)
		if err != nil {
			message["message"] = "Invalid token"
			views.SendResponse(w, http.StatusBadRequest, "", message)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func isValidToken(r *http.Request) error {
	_, _, err := VerifyToken(r)
	if err != nil {
		return err
	}
	return nil
}

func extractToken(r *http.Request) string {
	bearToken := r.Header.Get("Authorization")
	strArr := strings.Split(bearToken, " ")
	if len(strArr) == 2 {
		return strArr[1]
	}
	return ""
}

//VerifyToken takes the request, verifies token and returns it
func VerifyToken(r *http.Request) (*jwt.Token, jwt.MapClaims, error) {

	tokenString := extractToken(r)

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected sigining method: %v", token.Header["alg"])
		}
		return []byte("ENIGMA"), nil
	})
	if err != nil {
		return nil, nil, err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok && !token.Valid {
		return nil, nil, err
	}

	return token, claims, nil
}
