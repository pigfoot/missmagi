package main

import (
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode("release")

	r := gin.Default()
	r.LoadHTMLGlob("templates/*.tmpl")
	r.Static("/assets", "assets")

	r.GET("/", func(c *gin.Context) {
		c.Redirect(301, "/einvoice")
	})

	r.GET("/einvoice", func(c *gin.Context) {
		obj := gin.H{"PageTitle": "E-invoice"}
		c.HTML(200, "einvoice", obj)
	})

	// Listen and server on 0.0.0.0:8080
	port := "80"
	if os.Getenv("PORT") != "" {
		port = os.Getenv("PORT")
	}
	r.Run(":" + port)
}
