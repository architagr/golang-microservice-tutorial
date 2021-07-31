module github.com/architagr/golang-microservice-tutorial/testGrpc

go 1.16

replace github.com/architagr/golang-microservice-tutorial/rpc => ../rpc

require (
	github.com/architagr/golang-microservice-tutorial/rpc v0.0.0-00010101000000-000000000000
	google.golang.org/grpc v1.39.0
)
