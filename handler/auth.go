package handler

import (
	"fmt"
	"net/http"
	"time"

	"github.com/Damianen/norm/model"
	"github.com/golang-jwt/jwt/v5"
)

var secretKey = []byte("secret-key")

func createToken(management model.Management) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"management": management,
			"exp":        time.Now().Add(time.Hour * 24).Unix(),
		})

	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func verifyToken(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})

	if err != nil {
		return &jwt.Token{}, err
	}

	if !token.Valid {
		return &jwt.Token{}, fmt.Errorf("invalid token")
	}

	return token, nil
}

func getCookieData(r *http.Request) (model.Management, error) {
	cookie, err := r.Cookie("JWT-token")

	if err != nil {
		return model.Management{}, err
	}

	tokenString := cookie.Value
	token, err := verifyToken(tokenString)

	if err != nil {
		return model.Management{}, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		data := claims["management"].(map[string]interface{})
		management := model.Management{}
		management.Email = data["Email"].(string)
		management.Function = data["Function"].(string)
		management.Id = data["Id"].(string)
		management.Name = data["Name"].(string)
		management.Pnl = data["Pnl"].(string)
		return management, nil
	}

	return model.Management{}, nil
}

func verifyCookie(r *http.Request) error {
	cookie, err := r.Cookie("JWT-token")

	if err != nil {
		return err
	}

	tokenString := cookie.Value
	token, err := verifyToken(tokenString)

	if err != nil && !token.Valid {
		return err
	}

	return nil
}
