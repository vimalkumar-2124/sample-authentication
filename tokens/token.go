package tokens

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/vimalkumar-2124/sample-authentication/config"
	"github.com/vimalkumar-2124/sample-authentication/models"
)

// Generate JSON Web Token
func GenerateJWT(role string) (string, error) {
	jwtKey := config.EnvConfig("SECRET")
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["role"] = role
	claims["exp"] = time.Now().Add(time.Minute * 15).Unix()
	tokenInit := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := tokenInit.SignedString([]byte(jwtKey))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

// Extract token from the request Header
func ExtractToken(r *http.Request) string {
	authToken := r.Header.Get("Authorization")
	if authToken == "" {
		return ""
	}
	authToken = strings.ReplaceAll(authToken, "Bearer ", "")
	return authToken
}

// Verify the JWT
func VerifyToken(r *http.Request) (*jwt.Token, error) {
	tokenString := ExtractToken(r)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		//Make sure that the token method conform to "SigningMethodHMAC"
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(config.EnvConfig("SECRET")), nil

	})
	if err != nil {
		return nil, err
	}
	return token, nil
}

// func TokenValid(r *http.Request) error {
// 	token, err := VerifyToken(r)
// 	if err != nil {
// 		return err
// 	}
// 	if _, ok := token.Claims.(jwt.MapClaims); !ok && !token.Valid {
// 		return err
// 	}
// 	return nil
// }

// Extract the metadatas from the JWT
func ExtractTokenMetaData(r *http.Request) (models.TokenMetaData, error) {
	token, err := VerifyToken(r)
	if err != nil {
		return models.TokenMetaData{}, err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		role, ok := claims["role"].(string)
		if !ok {
			log.Println("Type check failed for ROLE")
			return models.TokenMetaData{}, err
		}
		exp, ok := claims["exp"].(float64)
		if !ok {
			log.Println("Type check failed for EXP")
			return models.TokenMetaData{}, err
		}
		return models.TokenMetaData{
			Role:   role,
			Expiry: int64(exp),
		}, nil
	}
	return models.TokenMetaData{}, err

}
