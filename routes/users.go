package routes

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"notes.com/app/models"
	"notes.com/app/utils"
)

func GetUsers(context *gin.Context) {

	users, err := models.GetUsers()

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Couldnt fetch user"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "User found", "user": users})
}

func GetUserById(context *gin.Context) {
	userId, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Invalid id"})
		return
	}

	user, err := models.GetUserById(userId)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Couldnt fetch user"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "User found", "user": user})
}

func CreateUser(context *gin.Context) {

	var user models.User
	err := context.ShouldBindJSON(&user)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Couldnt create user"})
		return
	}

	err = user.CreateUser()

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Couldnt create user"})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "Created user"})
}

func LoginUser(context *gin.Context) {
	var user models.User
	err := context.ShouldBindJSON(&user)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Couldnt bind user"})
		return
	}

	err = user.LoginUser()

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Couldnt login user"})
		return
	}

	token, err := utils.GenerateToken(user.Email, user.Id)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Couldnt autehnticate user"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "User logged in successfully!", "token": token})
}
