package routes

import (
	"bankdetails/controller"

	"github.com/gin-gonic/gin"
)

func Routes(r *gin.Engine) {
	r.POST("/api/accounts", controller.CreateAccount)
	r.POST("/api/transactions", controller.Transaction)
}
