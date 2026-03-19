package routes

import (
	"github.com/FernandaIshida/go-gin-api.git/controllers"
	"github.com/gin-gonic/gin"
)

func HandleRequests() {
	r := gin.Default()
	r.GET("/students", controllers.ShowAllStudents)
	r.GET("/students/:id", controllers.ShowStudentByID)
	r.GET("/students/cpf/:cpf", controllers.SearchStudentByCPF)
	r.POST("/students", controllers.CreateNewStudent)
	r.DELETE("students/:id", controllers.DeleteStudentByID)
	r.PATCH("/students/:id", controllers.UpdateStudentByID)
	r.Run()
}
