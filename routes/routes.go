package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/jclaudiotomasjr/alura/apip-gin-rest-app/controllers"
)

func HandleRequests() {
	r := gin.Default()
	r.GET("/users/", controllers.AllUsers)
	r.POST("/users/", controllers.CreateUser)
	r.DELETE("/users/:id", controllers.DeleteUser)
	r.GET("/users/:id", controllers.ReturnUser)
	r.PATCH("/users/:id", controllers.UpdateUser)
	r.Run()

}
