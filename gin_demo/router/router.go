package router

import (
	"github.com/gin-gonic/gin"
	"sgx-multichain/gin_demo/handlers"
	"sgx-multichain/gin_demo/middleware"
)

func Init() {
	r := gin.Default()
	r.Use(middleware.Cors())
	v1 := r.Group("/v1")
	{
		v1.GET("/onchain", handlers.PutOnChain)
	}
	r.Run(":8000")
}