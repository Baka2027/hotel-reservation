package handlers

import (
    "net/http"
    "github.com/gin-gonic/gin"

    "hotel/pkg/models"
    "hotel/pkg/config"
	"hotel/pkg/jwt"
)

func CreateUser(c *gin.Context){
    var user models.User
    if err := c.ShouldBindJSON(&user); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    config.DB.Create(&user)
    c.JSON(http.StatusOK, user)
}

func CreateListing(c *gin.Context){
    var listing models.Listing
    if err := c.ShouldBindJSON(&listing); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    config.DB.Create(&listing)
    c.JSON(http.StatusOK, listing)
}


func CreateBooking(c *gin.Context){
    var booking models.Booking
    if err := c.ShouldBindJSON(&booking); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    config.DB.Create(&booking)
    c.JSON(http.StatusOK, booking)
}

func ProtectedResource(c *gin.Context) {
    userID := c.MustGet("userID").(string)
    c.JSON(http.StatusOK, gin.H{"message": "Hello, authenticated user!", "userID": userID})
}

// Login handles user authentication and generates JWT token
func Login(c *gin.Context) {
    // Validate user credentials
    // If valid, generate JWT token
    userID := "123" // Example user ID
    token, err := jwt.GenerateJWT(userID)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"token": token})
}



