package controllers

import (
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"

	"github.com/ShahSau/CinemaNow/backend/helpers"
	"github.com/ShahSau/CinemaNow/backend/models"
)

// GetAllUsers is a function that returns all the users
func GetAllUsers(c *gin.Context) {
	recordPerPage, err := strconv.Atoi(c.Query("recordPerPage"))
	if err != nil || recordPerPage < 1 {
		recordPerPage = 10
	}

	page, errp := strconv.Atoi(c.Query("page"))
	if errp != nil || page < 1 {
		page = 1
	}

	startIndex := (page - 1) * recordPerPage
	startIndex, err = strconv.Atoi(c.Query("startIndex"))

	matchStage := bson.D{{Key: "$match", Value: bson.D{{}}}}
	projectStage := bson.D{
		{
			Key: "$project", Value: bson.D{
				{Key: "id", Value: 1},
				{Key: "first_name", Value: 1},
				{Key: "last_name", Value: 1},
				{Key: "email", Value: 1},
				{Key: "role", Value: 1},
				{Key: "phone", Value: 1},
			},
		},
	}

	record := bson.D{{Key: "$skip", Value: recordPerPage * (page - 1)}}
	limit := bson.D{{Key: "$limit", Value: recordPerPage}}

	result, errAgg := userCollection.Aggregate(c.Request.Context(), mongo.Pipeline{matchStage, projectStage, record, limit})

	if errAgg != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": errAgg.Error()})
		return
	}

	var allUsers []bson.M
	if err = result.All(c.Request.Context(), &allUsers); err != nil {
		log.Fatal(err)
	}

	c.JSON(http.StatusOK, gin.H{"data": allUsers, "page": page, "recordPerPage": recordPerPage, "startIndex": startIndex})

}

// GetUser is a function that returns a single user
func GetUser(c *gin.Context) {
	userId := c.Param("id")

	var user models.User

	defer c.Request.Body.Close()

	err := userCollection.FindOne(c.Request.Context(), bson.D{{Key: "user_id", Value: userId}}).Decode(&user)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": user, "status": http.StatusOK, "success": true, "error": false, "message": "User retrieved successfully"})

}

// UpdateUser is a function that updates a user
func UpdateUser(c *gin.Context) {
	userId := c.Param("id")
	var user models.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// find the user
	var foundUser models.User
	err := userCollection.FindOne(c.Request.Context(), bson.D{{Key: "user_id", Value: userId}}).Decode(&foundUser)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	user.Password = foundUser.Password
	user.CreatedAt = foundUser.CreatedAt
	user.ID = foundUser.ID
	user.Role = foundUser.Role
	user.User_id = foundUser.User_id
	user.UpdatedAt, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))

	_, err = userCollection.UpdateOne(c.Request.Context(), bson.D{{Key: "user_id", Value: userId}}, bson.D{{Key: "$set", Value: user}})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"error": false, "message": "User updated successfully", "status": http.StatusOK, "success": true, "data": user})
}

// DeleteUser is a function that deletes a user
func DeleteUser(c *gin.Context) {
	userId := c.Param("id")

	defer c.Request.Body.Close()

	userEmail, _ := c.Get("first_name")
	var isAdmin = helpers.IsAdmin(userEmail.(string))

	if !isAdmin {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "You are not authorized to view this resource"})
		return
	}

	_, err := userCollection.DeleteOne(c.Request.Context(), bson.D{{Key: "id", Value: userId}})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"error": false, "message": "User deleted successfully", "status": http.StatusOK, "success": true, "data": nil})
}

func ComparePassword(hashedPassword string, password string) (bool, string) {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		return false, "Invalid password"
	}
	return true, "Password is valid"
}

func HashPassword(password string) string {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		log.Panic("Error hashing password")
	}
	return string(bytes)
}
