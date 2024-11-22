package main

import (
	"github.com/gin-gonic/gin"
	"github.com/rhc07/simple-go-service/routes"
)

func main() {
	r := gin.Default()
	routes.GetVacationRouter(r)
	r.Run()
}
