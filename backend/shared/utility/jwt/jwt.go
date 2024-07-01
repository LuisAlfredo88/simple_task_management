package jwt

import (
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
)

var (
	secretKey = os.Getenv("JWT_KEY")
	algorithm = "HS256"
	audience  = os.Getenv("JWT_AUDIENCE")
	issuer    = os.Getenv("JWT_ISSUER")
)

func Parse(tokenString string) (jwt.MapClaims, error) {

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Check the signing method
		if token.Method.Alg() != algorithm {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(secretKey), nil
	})

	if err != nil {
		fmt.Printf("Error parsing token: %v\n", err)
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		fmt.Println("Invalid token")
		return nil, fmt.Errorf("invalid token")
	}

	if customData, ok := claims["customData"].(map[string]interface{}); ok {
		return customData, nil
	}

	return claims, nil
}

func GenerateJWT(customClaims map[string]interface{}) (string, error) {
	// Define the signing method (HS256)
	algorithm := jwt.SigningMethodHS256

	// Create a new JWT token with custom claims
	token := jwt.NewWithClaims(algorithm, jwt.MapClaims{
		"aud": audience,
		"iss": issuer,
		"exp": time.Now().Add(time.Hour * 1).Unix(),
		"uid": customClaims["userId"],
	})

	// Sign the token with the secret key
	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
