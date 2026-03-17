package main

import (
	"github.com/FernandaIshida/go-gin-api.git/models"
	"github.com/FernandaIshida/go-gin-api.git/routes"
)

func main() {
	models.Students = []models.Student{
		{Name: "John Doe", CPF: "12345678900", RG: "MG1234567"},
		{Name: "Jane Smith", CPF: "98765432100", RG: "SP7654321"},
	}
	routes.HandleRequests()
}
