package main

import (
	"github.com/gin-gonic/gin"
	"github.com/landofcoder/go-lang-gin-postgresql-example/controllers"
	"github.com/landofcoder/go-lang-gin-postgresql-example/models" // new
)

func main() {
	r := gin.Default()

	models.ConnectDatabase() // new

	r.GET("/contacts", controllers.FindContacts) // new
	r.POST("/contacts", controllers.CreateContact)
	r.GET("/contacts/:id", controllers.FindContact)
	r.PATCH("/contacts/:id", controllers.UpdateContact)
	r.DELETE("/contacts/:id", controllers.DeleteContact)

	r.Run()

	err := r.Run()
	if err != nil {
		return
	}
}
