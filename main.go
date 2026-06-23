package main

import (
	"context"
	"log"
	"os"

	"github.com/cassianobraz/crudGo/src/configuration/database/mongodb"
	"github.com/cassianobraz/crudGo/src/configuration/logger"
	"github.com/cassianobraz/crudGo/src/controller/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

// @title Meu Primeiro CRUD em GO
// @version 1.0
// @description API for crud operations on users
// @host localhost:8080
// @BasePath /
// @schemes http
// @license MIT
func main() {
	logger.Info("About to start user application")

	godotenv.Load()

	database, err := mongodb.NewMongoDBConnection(context.Background())
	if err != nil {
		log.Fatalf(
			"Error trying to connect to database, error=%s \n",
			err.Error())
		return
	}

	userController := initDependencies(database)

	router := gin.Default()
	gin.SetMode(gin.ReleaseMode)
	routes.InitRoutes(&router.RouterGroup, userController)

	if err := router.Run(":" + os.Getenv("PORT")); err != nil {
		log.Fatal(err)
	}
}
