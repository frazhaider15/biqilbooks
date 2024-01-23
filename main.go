package main

import (
	"fmt"
	"os"

	"github.com/biqilbooks/config"
	"github.com/biqilbooks/controllers"
	"github.com/biqilbooks/db"
	"github.com/gin-gonic/gin"
)

func main() {
	config.LoadEnvVariables()

	db.ConnectToDb()
	db.SyncDatabase()

	r := gin.Default()
	controllers.GetGinRoutes(r)

	fmt.Println("Starting on http://0.0.0.0:" + os.Getenv("PORT"))
	r.Run()
}
