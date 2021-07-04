package main

import (
	"github.com/architagr/golang-microservice-tutorial/authentication/models"
	"github.com/architagr/golang-microservice-tutorial/authentication/routers"
	"flag"
	"fmt"
	"log"
	"os"
)

var (
	port = flag.String("port", "8080", "port to be used")
	ip   = flag.String("ip", "localhost", "ip to be used")
)

func main() {
	flag.Parse()
	flags := models.NewFlags(*ip, *port)

	fmt.Println("Starting Api")

	logger := log.New(os.Stdout, "auth", 1)
	route := routers.NewRoute(logger, flags)
	engine := route.RegisterRoutes()

	url, _ := flags.GetApplicationUrl()
	engine.Run(*url)
}
