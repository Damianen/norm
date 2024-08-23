package handler

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var secretKey = []byte("secret-key")

func createToken(pnl string) (string, error) {
    token := jwt.NewWithClaims(jwt.SigningMethodHS256,
        jwt.MapClaims{
            "pnl": pnl,
            "exp": time.Now().Add(time.Hour * 24).Unix(),
        })

    tokenString, err := token.SignedString(secretKey)
    if err != nil {
        return "", err
    }

    return tokenString, nil
}


func verifyToken(tokenString string) (string, error) {
    token, err := jwt.Parse(tokenString, func (token *jwt.Token) (interface{}, error) {
        return secretKey, nil
    })

   if err != nil {
        return "", err
    }

    if !token.Valid {
        return "", fmt.Errorf("invalid token")
    }

    if claims, ok := token.Claims.(jwt.MapClaims); ok {
        return claims["pnl"].(string), nil
    }

    return "", fmt.Errorf("no pnl found")
}


