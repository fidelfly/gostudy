package jwtTool

import (
	"github.com/dgrijalva/jwt-go"
	"time"
	"fmt"
)

const secretKey = "com.fidelfly.jwt.secretKey@ltismdg&4370157"

func NewToken(playload map[string]interface{}) string {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := make(jwt.MapClaims)
	claims["iat"] = time.Now().Unix()
	claims["exp"] = time.Now().Add(time.Hour * time.Duration(1))

	if playload != nil {
		for k, v := range playload {
			claims[k] = v
		}
	}

	token.Claims = claims

	tokenStr, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return ""
	}
	return tokenStr
}

func ValidToken(tokenStr string) (bool, map[string]interface{}) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token)(interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected siging method: %v", token.Header["alg"])
		}
		return []byte(secretKey), nil
	});
	if err != nil {
		return false, nil
	}
	if token.Valid {
		return true, token.Claims.(jwt.MapClaims)
	}
	return false, nil
}