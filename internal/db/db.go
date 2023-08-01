// mongo.go

package db

import (
	"context"
	"fmt"
	"lixIQ/backend/internal/controllers"
	"lixIQ/backend/internal/routes"
	"lixIQ/backend/internal/services"
	"lixIQ/backend/internal/utils"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var (
	mongoClient *mongo.Client
	database    *mongo.Database
	ctx         context.Context

	userService         services.UserService
	UserController      controllers.UserController
	UserRouteController routes.UserRouteController

	authCollection      *mongo.Collection
	authService         services.AuthService
	AuthController      controllers.AuthController
	AuthRouteController routes.AuthRouteController
)

func init() {

	var appConfig = utils.LoadConfig()

	ctx = context.TODO()

	mongoconn := options.Client().ApplyURI(appConfig.MongoUri)
	mongoclient, err := mongo.Connect(ctx, mongoconn)

	if err != nil {
		panic(err)
	}

	if err := mongoclient.Ping(ctx, readpref.Primary()); err != nil {
		panic(err)
	}

	fmt.Println("MongoDB successfully connected...")

	// Collections
	authCollection = mongoclient.Database(appConfig.MongoDbName).Collection("users")
	userService = services.NewUserServiceImpl(authCollection, ctx)
	authService = services.NewAuthService(authCollection, ctx)
	AuthController = controllers.NewAuthController(authService, userService)
	AuthRouteController = routes.NewAuthRouteController(AuthController)

	UserController = controllers.NewUserController(userService)
	UserRouteController = routes.NewRouteUserController(UserController)

}

func Close() {
	if mongoClient != nil {
		mongoClient.Disconnect(context.Background())
	}
}

// Database variable as a function
func GetDatabase() *mongo.Database {
	return database
}
