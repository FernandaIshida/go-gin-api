package routes

import (
	"github.com/FernandaIshida/go-gin-api.git/controllers"
	"github.com/gin-gonic/gin"
)

func HandleRequests() {
	r := gin.Default()
	r.GET("/students", controllers.ShowAllStudents)
	r.Run()
}
