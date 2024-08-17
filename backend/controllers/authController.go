package controllers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/ShahSau/CinemaNow/backend/database"
	"github.com/ShahSau/CinemaNow/backend/helpers"
	"github.com/ShahSau/CinemaNow/backend/models"
	"github.com/ShahSau/CinemaNow/backend/types"
)

var userCollection *mongo.Collection = database.GetCollection(database.DB, "users")

// Register is a function that handles the registration of a user
func Register(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	count, err := userCollection.CountDocuments(c.Request.Context(), bson.D{{Key: "email", Value: user.Email}})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if count > 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "email already exists"})
		return
	}

	count, err = userCollection.CountDocuments(c.Request.Context(), bson.D{{Key: "phone", Value: user.Phone}})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if count > 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "phone number already exists"})
		return
	}

	password := HashPassword(user.Password)
	user.Password = password

	user.CreatedAt, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	user.UpdatedAt, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	user.ID = primitive.NewObjectID()
	user.User_id = user.ID.Hex()
	user.Role = "User"

	token, refreshToken, _ := helpers.GenerateAllTokens(user.User_id, user.Email, user.First_name, user.Last_name)
	user.Token = token
	user.RefreshToken = refreshToken

	_, err = userCollection.InsertOne(c.Request.Context(), user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"error": false, "message": "User created successfully", "data": user, "status": http.StatusCreated, "success": true})

}

// Login is a function that handles the login of a user
func Login(c *gin.Context) {
	var user types.Loginuser
	var foundUser models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := userCollection.FindOne(c.Request.Context(), bson.D{{Key: "email", Value: user.Email}}).Decode(&foundUser)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	passwordIsValid, msg := ComparePassword(foundUser.Password, user.Password)

	if !passwordIsValid {
		c.JSON(http.StatusUnauthorized, gin.H{"error": msg})
		return
	}

	token, refreshToken, _ := helpers.GenerateAllTokens(foundUser.User_id, foundUser.Email, foundUser.First_name, foundUser.Last_name)

	helpers.UpdateAllTokens(foundUser.User_id, token, refreshToken)

	c.JSON(http.StatusOK, gin.H{"error": false, "message": "User logged in successfully", "data": foundUser, "status": http.StatusOK, "success": true})

}

// Logout is a function that handles the logout of a user
func Logout(c *gin.Context) {
	var user struct {
		User_id string `json:"user_id" binding:"required"`
	}
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	helpers.UpdateAllTokens(user.User_id, "", "")

	c.Set("email", "")
	c.Set("first_name", "")
	c.Set("last_name", "")
	c.Set("user_id", "")

	c.JSON(http.StatusOK, gin.H{"error": false, "message": "User logged out successfully", "status": http.StatusOK, "success": true})
}

// ForgotPassword is a function that handles the forgot password of a user
func ForgotPassword(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Forgot password",
	})
}

// ResetPassword is a function that handles the reset password of a user
func ResetPassword(c *gin.Context) {
	var user types.PasswordReset
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var foundUser models.User
	err := userCollection.FindOne(c.Request.Context(), bson.D{{Key: "email", Value: user.Email}}).Decode(&foundUser)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	passwordIsValid, msg := ComparePassword(foundUser.Password, user.OldPassword)

	if !passwordIsValid {
		c.JSON(http.StatusUnauthorized, gin.H{"error": msg})
		return
	}

	new_password := HashPassword(user.NewPassword)
	foundUser.Password = new_password

	foundUser.UpdatedAt, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))

	_, err = userCollection.UpdateOne(c.Request.Context(), bson.D{{Key: "email", Value: user.Email}}, bson.D{{Key: "$set", Value: foundUser}})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"error": false, "message": "Password reset successfully", "status": http.StatusOK, "success": true, "data": foundUser})
}
