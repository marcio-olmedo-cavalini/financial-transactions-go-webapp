package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/marcio-olmedo-cavalini/financial-transactions-go-webapp/models"
)

func ShowUserListPage(c *gin.Context) {
	var usuarios = models.ListUsers()
	c.HTML(http.StatusOK, "user.html", gin.H{
		"usuarios": usuarios,
	})
}

func ShowNewUserPage(c *gin.Context) {
	c.HTML(http.StatusOK, "newuser.html", nil)
}

func SaveNewUser(c *gin.Context) {
	emailValue := c.PostForm("email")
	nameValue := c.PostForm("nome")
	passwdValue := models.GenerateRandomPassword()
	passwdHashValue, _ := models.HashPassword(passwdValue)
	user := models.User{Nome: nameValue, Email: emailValue, Password: passwdHashValue}
	models.CreateUser(user)
	sendEmailToUser(nameValue, passwdValue)
	c.Redirect(http.StatusFound, "/user")
}

func sendEmailToUser(nameValue string, passwdValue string) {
	fmt.Printf("[ATENCAO] - %s, Sua senha de acesso ao sistema é: %s\n", nameValue, passwdValue)
}

func ShowEditUserPage(c *gin.Context) {
	var user models.User
	id := c.Query("id")
	fmt.Println("id: " + id)
	user = models.FindUserById(id)
	c.HTML(http.StatusOK, "edituser.html", gin.H{
		"nome":  user.Nome,
		"email": user.Email,
		"id":    user.ID,
	})
}

func UpdateUser(c *gin.Context) {
	emailValue := c.PostForm("email")
	nameValue := c.PostForm("nome")
	idValue := c.PostForm("id")
	user := models.User{Nome: nameValue, Email: emailValue}
	models.UpdateUser(user, idValue)
	c.Redirect(http.StatusFound, "/user")
}
