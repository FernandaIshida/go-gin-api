package main

import (
	"fmt"

	"github.com/FernandaIshida/go-gin-api.git/database"
	"github.com/FernandaIshida/go-gin-api.git/messaging"
	"github.com/FernandaIshida/go-gin-api.git/routes"
)

func main() {
	database.ConnectDataBase()

	conn, ch := messaging.ConnectRabbitMQ()

	fmt.Println("Connected to RabbitMQ successfully!")

	defer conn.Close()
	defer ch.Close()

	routes.HandleRequests()
}
