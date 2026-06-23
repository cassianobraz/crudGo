package routes

import (
	"github.com/cassianobraz/crudGo/src/controller"
	"github.com/cassianobraz/crudGo/src/model"
	"github.com/gin-gonic/gin"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func InitRoutes(
	r *gin.RouterGroup,
	userController controller.UserControllerInterface) {

	r.GET("/getUserById/:userId", model.VerifyTokenMiddleware, userController.FindUserById)
	r.GET("/getUserByEmail/:userEmail", model.VerifyTokenMiddleware, userController.FindUserByEmail)
	r.POST("/createUser", userController.CreateUser)
	r.PUT("/updateUser/:userId", userController.UpdateUser)
	r.DELETE("/deleteUser/:userId", userController.DeleteUser)

	r.POST("/login", userController.LoginUser)

	r.GET("swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
