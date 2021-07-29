package main

import (
	"github.com/architagr/golang-microservice-tutorial/employee/add"
	"github.com/architagr/golang-microservice-tutorial/employee/data"
	"github.com/architagr/golang-microservice-tutorial/employee/delete"
	"github.com/architagr/golang-microservice-tutorial/employee/get"
	"github.com/architagr/golang-microservice-tutorial/employee/persistance"
	"github.com/architagr/golang-microservice-tutorial/employee/update"

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
	ginR.Use(CORSMiddleware())
	group := ginR.Group("api/employees")

	repo := getPersistanceObj()
	registerGetRoutes(group, repo)
	registerPutRoutes(group, repo)
	registerAddRoutes(group, repo)
	registerDeleteRoutes(group, repo)

	ginR.Run(*url)
}
func CORSMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {

        c.Header("Access-Control-Allow-Origin", "*")
        c.Header("Access-Control-Allow-Credentials", "true")
        c.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
        c.Header("Access-Control-Allow-Methods", "POST,HEAD,PATCH, OPTIONS, GET, PUT")

        if c.Request.Method == "OPTIONS" {
            c.AbortWithStatus(204)
            return
        }

        c.Next()
    }
}
func getPersistanceObj() persistance.IEmployeeDbContext {
	return persistance.InitMongoDb("", "")
}

func registerGetRoutes(group *gin.RouterGroup, repo persistance.IEmployeeDbContext) {
	service := get.InitService(repo)
	handler := get.InitHandler(service)
	getRouter := get.InitRouter(handler)
	getRouter.RegisterRoutes(group)
}

func registerPutRoutes(group *gin.RouterGroup, repo persistance.IEmployeeDbContext) {
	service := update.InitService(repo)
	handler := update.InitHandler(service)
	putRouter := update.InitRouter(handler)
	putRouter.RegisterRoutes(group)
}

func registerAddRoutes(group *gin.RouterGroup, repo persistance.IEmployeeDbContext) {
	service := add.InitService(repo)
	handler := add.InitHandler(service)
	addRouter := add.InitRouter(handler)
	addRouter.RegisterRoutes(group)
}

func registerDeleteRoutes(group *gin.RouterGroup, repo persistance.IEmployeeDbContext) {
	service := delete.InitService(repo)
	handler := delete.InitHandler(service)
	deleteRouter := delete.InitRouter(handler)
	deleteRouter.RegisterRoutes(group)
}
