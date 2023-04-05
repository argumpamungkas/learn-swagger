package main

import (
	"DTS/Chapter-2/Sesi/sesi-5-swagger/repo"
	"DTS/Chapter-2/Sesi/sesi-5-swagger/routers"
)

// @title Car API
// @vesion 1.0
// @description This is a sample service for managing cars

// @host localhost:9000
// @BasePath /
func main() {

	repo.StartDB()

	routers.StartServer().Run(":9000")

}
