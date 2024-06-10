package jwt

import (
    "github.com/dgrijalva/jwt-go"
    "time"

)

var secretKey = []byte("keyone")

// GenerateJWT generates a JWT token for the given user ID
func GenerateJWT(userID string) (string, error) {
    // Create a new JWT token
    token := jwt.New(jwt.SigningMethodHS256)
    
    // Set claims (payload)
    claims := token.Claims.(jwt.MapClaims)
    claims["userID"] = userID
    claims["exp"] = time.Now().Add(time.Hour * 24).Unix() // Token expires in 24 hours

    // Sign and get the complete encoded token as a string
    tokenString, err := token.SignedString(secretKey)
    if err != nil {
        return "", err
    }
    
    return tokenString, nil
}


