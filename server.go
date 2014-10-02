package main

import (
	"os"

	"github.com/gin-gonic/gin"
)

const (
	NEW_RELIC_LICENSE = ""
)

func main() {
	err := InitNewrelicAgent(NEW_RELIC_LICENSE, "missmagi", true)
	if err != nil {
		panic(err)
	}

	gin.SetMode("release")

	r := gin.Default()
	r.LoadHTMLGlob("templates/*.tmpl")
	r.Static("/assets", "assets")

	r.GET("/", func(c *gin.Context) {
		c.HTML(200, "body", nil)
	})

	r.GET("/einvoice", HandlerEinvoiceGet)
	r.POST("/einvoice", HandlerEinvoicePost)

	// Listen and server on 0.0.0.0:8080
	port := "80"
	if os.Getenv("PORT") != "" {
		port = os.Getenv("PORT")
	}

	r.Use(HandlerNewRelic)
	r.Run(":" + port)
}
