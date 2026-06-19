package main

import (
	"github.com/cassianobraz/crudGo/src/controller"
	"github.com/cassianobraz/crudGo/src/model/repository"
	"github.com/cassianobraz/crudGo/src/model/service"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

func initDependencies(
	database *mongo.Database,
) controller.UserControllerInterface {
	repo := repository.NewUserRepository(database)
	service := service.NewUserDomainService(repo)
	return controller.NewUserControllerInterface(service)
}
