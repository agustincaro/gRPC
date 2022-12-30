package main

import (
	"github.com/agustincaro/gRPC/cmd/server/routes"
	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()

	router := routes.NewRouter(r)
	router.MapRoutes()

	if err := r.Run(":8084"); err != nil {
		panic(err)
	}
}
