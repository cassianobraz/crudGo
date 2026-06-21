package repository

import (
	"context"
	"testing"

	"github.com/cassianobraz/crudGo/src/model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"go.mongodb.org/mongo-driver/v2/x/mongo/driver/drivertest"
)

func TestUserRepository_CreateUser(t *testing.T) {
	databaseName := "user_database_test"
	collectionName := "user_collection_test"

	t.Setenv("MONGODB_USER_DB", collectionName)

	t.Run("when_sending_a_valid_domain_returns_success", func(t *testing.T) {
		mockDeployment := drivertest.NewMockDeployment(
			bson.D{
				{Key: "ok", Value: 1},
				{Key: "n", Value: 1},
				{Key: "acknowledged", Value: true},
			},
		)

		clientOptions := options.Client()
		clientOptions.Deployment = mockDeployment

		client, err := mongo.Connect(clientOptions)
		require.NoError(t, err)

		t.Cleanup(func() {
			err := client.Disconnect(context.Background())
			assert.NoError(t, err)
		})

		databaseMock := client.Database(databaseName)

		repo := NewUserRepository(databaseMock)

		domain := model.NewUserDomain("teste@teste.com", "test", "teste", 90)

		userDomain, err := repo.CreateUser(domain)

		_, errID := bson.ObjectIDFromHex(userDomain.GetID())

		assert.Nil(t, err)
		assert.Nil(t, errID)
		assert.Equal(t, userDomain.GetEmail(), domain.GetEmail())
		assert.Equal(t, userDomain.GetName(), domain.GetName())
		assert.Equal(t, userDomain.GetAge(), domain.GetAge())
		assert.Equal(t, userDomain.GetPassword(), domain.GetPassword())
	})

	t.Run("return_error_from_database", func(t *testing.T) {
		mockDeployment := drivertest.NewMockDeployment(
			bson.D{
				{Key: "ok", Value: 0},
			},
		)

		clientOptions := options.Client()
		clientOptions.Deployment = mockDeployment

		client, err := mongo.Connect(clientOptions)
		require.NoError(t, err)

		t.Cleanup(func() {
			err := client.Disconnect(context.Background())
			assert.NoError(t, err)
		})

		databaseMock := client.Database(databaseName)

		repo := NewUserRepository(databaseMock)

		domain := model.NewUserDomain("teste@teste.com", "test", "teste", 90)

		userDomain, err := repo.CreateUser(domain)

		assert.NotNil(t, err)
		assert.Nil(t, userDomain)
	})
}
