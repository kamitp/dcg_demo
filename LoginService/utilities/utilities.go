package utilities

import (
	"encoding/base64"
	"encoding/json"
	"strings"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

func CreateResponse(status string, msg interface{}) string {
	json_bytes, _ := json.Marshal(map[string]interface{}{"status": status, "message": msg})
	return string(json_bytes)
}

func CreateJWT(userName string, secrete string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := make(jwt.MapClaims)
	claims["iss"] = "Hello Fresh"                         // issuer
	claims["sub"] = "Test Assignmnet"                     // subject
	claims["user"] = userName                             // user name
	claims["admin"] = false                               // admin role
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix() // expiration time
	token.Claims = claims
	return token.SignedString([]byte(secrete))
}

func ValidateJWTToken(jwtToken string, secrete string) (jwt.MapClaims, error) {

	// Validation steps
	// 1. If token exist && token not expired
	claims := jwt.MapClaims{}
	_, err := jwt.ParseWithClaims(jwtToken, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(secrete), nil
	})
	if err != nil {
		return nil, err
	}
	return claims, nil
}

func DecodeBasicAuthorizationKey(authKey string) ([]string, error) {
	key := strings.Split(authKey, " ")[1]
	nameAndPasswd, err := base64.StdEncoding.DecodeString(key)
	if err != nil {
		return nil, err
	}
	return strings.Split(string(nameAndPasswd), ":"), nil
}
