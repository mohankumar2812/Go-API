package controller

import (
	"example/model"
	"example/service"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	var newUser model.User

	if err := c.ShouldBindJSON(&newUser); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "worng model"})
	}

	result, err := service.CreateUser(&newUser)

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, err)
	}

	c.IndentedJSON(http.StatusOK, result)

}

func Login(c *gin.Context) {

	var newUser model.User;

	if err := c.ShouldBindJSON(&newUser); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "worng model"})
	}

	result, err := service.LoginUser(&newUser);

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, err.Error())
	}

	c.IndentedJSON(http.StatusOK, result)

}

func GetUser(c *gin.Context) {

	mail := c.GetHeader("Authorization")

	email, err := service.ValidateJWT(mail)

	fmt.Println("email address",email)

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, err.Error())
	}

	user, err := model.GetUser(email.(string))

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message":"mail not found"})
	}

	c.IndentedJSON(http.StatusOK, user)

}

func UpdateUser(c *gin.Context) {
	mail := c.Param("mail")
	name := c.Param("name")

	if mail == "" {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "please give mail"})
	}

	if name == "" {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "please give updated name"})
	}

	result, err := model.UpdateUser(mail, name)

	if err != nil {
		c.IndentedJSON(http.StatusExpectationFailed, gin.H{"message": "error occure"})
	}

	c.IndentedJSON(http.StatusOK, result)
}

func DeleteUser(c *gin.Context) {
	mail := c.Param("mail")

	if mail == "" {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "please give mail address"})
	}

	result, err := model.DeleteUser(mail) 

	if err != nil {
		c.IndentedJSON(http.StatusExpectationFailed, gin.H{"message": "user not found"})
	}

	c.IndentedJSON(http.StatusOK, result)
}
