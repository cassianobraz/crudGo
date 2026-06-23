package repository

import (
	"context"
	"os"

	"github.com/cassianobraz/crudGo/src/configuration/logger"
	"github.com/cassianobraz/crudGo/src/configuration/rest_err"
	"github.com/cassianobraz/crudGo/src/model"
	"github.com/cassianobraz/crudGo/src/model/repository/entity/converter"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.uber.org/zap"
)

// CreateUser Creates a new user
// @Summary Create a new user
// @Description Create a new user with the provided user information
// @Tags Users
// @Accept json
// @Produce json
// @Param userRequest body request.UserRequest true "User information for registration"
// @Success 200 {object} response.UserResponse
// @Failure 400 {object} rest_err.RestErr
// @Failure 500 {object} rest_err.RestErr
// @Router /createUser [post]
func (ur *userRepository) CreateUser(
	userDomain model.UserDomainInterface,
) (model.UserDomainInterface, *rest_err.RestErr) {
	logger.Info("Init createUser repository",
		zap.String("journey", "createUser"))

	collection_name := os.Getenv(MONGODB_USER_DB)

	collection := ur.databaseConnection.Collection(collection_name)

	value := converter.ConvertDomainToEntity(userDomain)

	result, err := collection.InsertOne(context.Background(), value)
	if err != nil {
		logger.Error("Error trying to create user",
			err,
			zap.String("journey", "createUser"))
		return nil, rest_err.NewInternalServerError(err.Error())
	}

	value.ID = result.InsertedID.(bson.ObjectID)

	logger.Info(
		"CreateUser repository executed successfully",
		zap.String("userId", value.ID.Hex()),
		zap.String("journey", "createUser"))

	return converter.ConvertEntityToDomain(*value), nil
}
