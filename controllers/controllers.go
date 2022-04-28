package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jclaudiotomasjr/alura/apip-gin-rest-app/database"
	"github.com/jclaudiotomasjr/alura/apip-gin-rest-app/models"
)

func AllUsers(c *gin.Context) {
	var users []models.User
	database.DB.Find(&users)
	c.JSON(http.StatusOK, users)
}

func CreateUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"erro": err.Error(),
		})
	}
	database.DB.Create(&user)
	c.JSON(http.StatusOK, gin.H{
		"usuário": "Criado com Sucesso!"})
}

func DeleteUser(c *gin.Context) {
	var user models.User
	id := c.Params.ByName("id")
	database.DB.First(&user, id)
	if user.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"status": "Usuário não encontrado!"})
		return
	}
	database.DB.Delete(&user, id)
	c.JSON(http.StatusOK, user)
}

func ReturnUser(c *gin.Context) {
	var user models.User
	id := c.Params.ByName("id")
	database.DB.First(&user, id)
	if user.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"Status": "usuário não encontrado!"})
	}
	c.JSON(http.StatusOK, user)
}

func UpdateUser(c *gin.Context) {
	var user models.User
	id := c.Params.ByName("id")
	database.DB.First(&user, id)
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"Erro": err.Error()})
		return
	}
	database.DB.Model(&user).UpdateColumns(user)
	c.JSON(200, gin.H{
		"Status": "usuário atualizado com sucesso!"})
}
