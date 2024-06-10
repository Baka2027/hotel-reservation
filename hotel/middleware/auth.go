package middleware

import (
    "fmt"
    "net/http"
    "github.com/gin-gonic/gin"
    "github.com/dgrijalva/jwt-go"
)

var secretKey = []byte("keyone")

// AuthMiddleware validates JWT token
func AuthMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        tokenString := c.GetHeader("Authorization")
        if tokenString == "" {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization token is missing"})
            c.Abort()
            return
        }

        // Strip the 'Bearer ' prefix from the token string if present
        if len(tokenString) > 7 && tokenString[:7] == "Bearer " {
            tokenString = tokenString[7:]
        }

        // Parse and validate token
        token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
            // Check the signing method
            if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
                return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
            }
            return secretKey, nil
        })
        if err != nil {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
            c.Abort()
            return
        }

        // Check if token is valid
        if !token.Valid {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Token is invalid"})
            c.Abort()
            return
        }

        // Token is valid, extract user ID
        if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
            c.Set("userID", claims["userID"])
        } else {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Token is invalid"})
            c.Abort()
            return
        }

        // Token is valid, proceed with the request
        c.Next()
    }
}
