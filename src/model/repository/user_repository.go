package repository

import (
	"os"

	"github.com/cassianobraz/crudGo/src/configuration/rest_err"
	"github.com/cassianobraz/crudGo/src/model"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

const (
	MONGODB_USER_DB         = "MONGODB_USER_DB"
	MONGODB_USER_COLLECTION = "MONGODB_USER_COLLECTION"
)

func NewUserRepository(
	database *mongo.Database,
) UserRepository {
	return &userRepository{
		database,
	}
}

type userRepository struct {
	databaseConnection *mongo.Database
}

func getUserCollectionName() string {
	collectionName := os.Getenv(MONGODB_USER_COLLECTION)
	if collectionName == "" {
		return os.Getenv(MONGODB_USER_DB)
	}
	return collectionName
}

type UserRepository interface {
	CreateUser(
		userDomain model.UserDomainInterface,
	) (model.UserDomainInterface, *rest_err.RestErr)

	UpdateUser(
		userId string,
		userDomain model.UserDomainInterface,
	) *rest_err.RestErr

	DeleteUser(
		userId string,
	) *rest_err.RestErr

	FindUserByEmail(
		email string,
	) (model.UserDomainInterface, *rest_err.RestErr)

	FindUserByID(
		id string,
	) (model.UserDomainInterface, *rest_err.RestErr)
}
