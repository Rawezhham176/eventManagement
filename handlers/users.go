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

	token, err := utils.GenerateToken(user.Email, strconv.FormatInt(user.ID, 10))

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusAccepted, gin.H{"user authorised": user.Email, "token": token})
}
