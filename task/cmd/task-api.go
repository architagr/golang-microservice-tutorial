package main

import (
	"github.com/architagr/golang-microservice-tutorial/task/add"
	"github.com/architagr/golang-microservice-tutorial/task/data"
	"github.com/architagr/golang-microservice-tutorial/task/delete"
	"github.com/architagr/golang-microservice-tutorial/task/get"
	"github.com/architagr/golang-microservice-tutorial/task/persistance"
	"github.com/architagr/golang-microservice-tutorial/task/update"

	"flag"

	"github.com/gin-gonic/gin"
)

var (
	port = flag.String("port", "8080", "port to be used")
	ip   = flag.String("ip", "localhost", "ip to be used")
)

func main() {
	flag.Parse()
	flags := data.NewFlags(*ip, *port)
	url, _ := flags.GetApplicationUrl()
	ginR := gin.Default()
	group := ginR.Group("api/task")

	repo := getPersistanceObj()
	registerGetRoutes(group, repo)
	registerPutRoutes(group, repo)
	registerAddRoutes(group, repo)
	registerDeleteRoutes(group, repo)

	ginR.Run(*url)
}

func getPersistanceObj() persistance.ITaskDbContext {
	return persistance.InitMongoDb("", "")
}

func registerGetRoutes(group *gin.RouterGroup, repo persistance.ITaskDbContext) {
	service := get.InitService(repo)
	handler := get.InitHandler(service)
	getRouter := get.InitRouter(handler)
	getRouter.RegisterRoutes(group)
}

func registerPutRoutes(group *gin.RouterGroup, repo persistance.ITaskDbContext) {
	service := update.InitService(repo)
	handler := update.InitHandler(service)
	putRouter := update.InitRouter(handler)
	putRouter.RegisterRoutes(group)
}

func registerAddRoutes(group *gin.RouterGroup, repo persistance.ITaskDbContext) {
	service := add.InitService(repo)
	handler := add.InitHandler(service)
	addRouter := add.InitRouter(handler)
	addRouter.RegisterRoutes(group)
}

func registerDeleteRoutes(group *gin.RouterGroup, repo persistance.ITaskDbContext) {
	service := delete.InitService(repo)
	handler := delete.InitHandler(service)
	deleteRouter := delete.InitRouter(handler)
	deleteRouter.RegisterRoutes(group)
}
