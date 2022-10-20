package tokens

import (
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/vimalkumar-2124/sample-authentication/config"
)

// Generate JSON Web Token
func GenerateJWT(email string) (string, error) {
	jwtKey := config.Config("SECRET")
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["email"] = email
	claims["exp"] = time.Now().Add(time.Minute * 1).Unix()
	tokenInit := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := tokenInit.SignedString([]byte(jwtKey))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
