package main

import (
	"finalProject/database"
	"finalProject/routes"
)

func main() {
	database.StartDB()
	r := routes.StartApp()
	r.Run(":8080")
}
