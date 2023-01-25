package main

import (
	"context"
	"fmt"
	"log"

	"github.com/Darshanbennur/gin_API/controllers"
	"github.com/Darshanbennur/gin_API/services"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
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

	mongoconn := options.Client().ApplyURI("Enter Your MongoDB URL here")
	mongoclient, err := mongo.Connect(ctx, mongoconn)
	if err != nil {
		log.Fatal(err)
	}
	err = mongoclient.Ping(ctx, readpref.Primary())
	if err != nil{
		log.Fatal(err)
	}
	fmt.Println("Mongo Connection is Established...")

	usercollection = (*mongo.Collection)(mongoclient.Database("userdb").Collection("users"))
	userservice = services.NewUserService(usercollection, ctx)
	usercontroller = controllers.New(userservice)
	server = gin.Default()
}

func main() {
	defer mongoClient.Disconnect(ctx)

	basepath := server.Group("/v1")
	usercontroller.RegisterRoutes(basepath)

	fmt.Println("Initializing the Server ...")
	log.Fatal(server.Run(":3000"))
}
