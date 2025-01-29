package main

import (
	"bankdetails/data"
	"bankdetails/kafka"
	"bankdetails/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	data.DbConnect()
	kafka.ConnectKafka()

	go kafka.ConsumeTransactions()
	r := gin.Default()
	routes.Routes(r)
	r.Run(":8080")
}
