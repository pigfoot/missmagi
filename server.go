package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	r.LoadHTMLTemplates("templates/*")

	r.GET("/einvoice", func(c *gin.Context) {
		c.String(200, "einvoice.tmpl")
	})

	// Listen and server on 0.0.0.0:8080
	r.Run(":80")
}
