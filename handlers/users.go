package handlers

import (
	"eventManagement/model"
	"eventManagement/utils"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	var newUser model.User
	err := c.ShouldBindJSON(&newUser)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err = newUser.Save()
	if err != nil {
		panic(fmt.Sprintf("create user error: %s", err))
		return
	}
	c.JSON(http.StatusCreated, gin.H{"user created": newUser})
}

func LoginUser(c *gin.Context) {
	var user model.User
	err := c.ShouldBindJSON(&user)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = user.ValidateUserCredential()

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	token, err := utils.GenerateToken(user.Email, strconv.FormatInt(user.UserId, 10))

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusAccepted, gin.H{"user authorised": user.Email, "token": token})
}

func ForgotPassword(c *gin.Context) {
	var requestBody map[string]interface{}
	if err := c.BindJSON(&requestBody); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	email := requestBody["email"].(string)

	user, err := model.ValidateUserByEmail(email)

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	token, err := utils.GenerateToken(user.Email, strconv.FormatInt(user.UserId, 10))

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if err := utils.SendResetEmail(user.Email, token); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not send email"})
		return
	}

	c.JSON(http.StatusAccepted, gin.H{"Reset link sent to your email": email, "token": token})
}

func ResetPassword(c *gin.Context) {
	var requestBody map[string]interface{}
	if err := c.BindJSON(&requestBody); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	token := requestBody["email"].(string)
	newPassword := requestBody["email"].(string)
	userId, err := utils.VerifyToken(token)

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"Invalid token": token})
		return
	}

	err = model.ResetUserPassword(userId, newPassword)

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"Message": "user password has been changed"})
}
