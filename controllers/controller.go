package controllers

import (
	"github.com/FernandaIshida/go-gin-api.git/models"
	"github.com/gin-gonic/gin"
)

func ShowAllStudents(c *gin.Context) {
	c.JSON(200, models.Students)
}
