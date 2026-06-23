package repository

import (
	"context"
	"os"

	"github.com/cassianobraz/crudGo/src/configuration/logger"
	"github.com/cassianobraz/crudGo/src/configuration/rest_err"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.uber.org/zap"
)

// DeleteUser deletes a user with the specified ID.
// @Summary Delete User
// @Description Deletes a user based on the ID provided as a parameter.
// @Tags Users
// @Accept json
// @Produce json
// @Param userId path string true "ID of the user to be deleted"
// @Success 200
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Failure 400 {object} rest_err.RestErr
// @Failure 500 {object} rest_err.RestErr
// @Router /deleteUser/{userId} [delete]
func (ur *userRepository) DeleteUser(
	userId string,
) *rest_err.RestErr {
	logger.Info("Init DeleteUser repository",
		zap.String("journey", "DeleteUser"))

	collection_name := os.Getenv(MONGODB_USER_DB)
	collection := ur.databaseConnection.Collection(collection_name)

	userIdHex, _ := bson.ObjectIDFromHex(userId)

	filter := bson.D{{Key: "_id", Value: userIdHex}}

	_, err := collection.DeleteOne(context.Background(), filter)
	if err != nil {
		logger.Error("Error trying to updater user",
			err,
			zap.String("journey", "DeleteUser"))
		return rest_err.NewInternalServerError(err.Error())
	}

	logger.Info(
		"DeleteUser repository executed successfully",
		zap.String("userId", userId),
		zap.String("journey", "DeleteUser"))

	return nil
}
