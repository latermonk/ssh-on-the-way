package main

import (

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/takumines/vue-gin-file-uploader/server/handler"
)

func main()  {
	r := gin.Default()

	r.Use(cors.New(cors.Config {
		AllowOrigins: []string{"http://localhost:8080"},
		AllowMethods: []string{"GET", "POST", "DELETE", "OPTIONS"},
		AllowHeaders: []string{"*"},
	}))

	r.POST("/images", handler.Upload)
	r.Run(":8888")
}
