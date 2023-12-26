package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"

	// User defined folders
	"github.com/Darshanbennur/gin_API/controllers"
	"github.com/Darshanbennur/gin_API/services"
)

var (
	server         *gin.Engine
	userservice    services.UserService
	usercontroller controllers.UserController
	ctx            context.Context
	usercollection *mongo.Collection
	mongoClient    *mongo.Client
	err            error
)

func init() {
	ctx = context.TODO()

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	mongoURL := os.Getenv("MONGODB_URL")
	if mongoURL == "" {
		log.Fatal("MongoDB URL is not set")
	}
	mongoconn := options.Client().ApplyURI(mongoURL)

	mongoclient, err := mongo.Connect(ctx, mongoconn)
	if err != nil {
		log.Fatal(err)
		log.Fatal("Error in connecting with database")
	}
	err = mongoclient.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Mongo connection is established...")

	usercollection = (*mongo.Collection)(mongoclient.Database("userdb").Collection("users"))
	userservice = services.NewUserService(usercollection, ctx)
	usercontroller = controllers.New(userservice)
	server = gin.Default()
}

func main() {
	defer mongoClient.Disconnect(ctx)

	basepath := server.Group("/v1")
	usercontroller.RegisterRoutes(basepath)

	fmt.Println("Initializing the server ...")
	log.Fatal(server.Run(":3000"))
}
