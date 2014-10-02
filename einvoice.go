package main

import "github.com/gin-gonic/gin"

func HandlerEinvoiceGet(c *gin.Context) {
	obj := gin.H{
		"PageTitle":    "E-invoice",
		"PageSubTitle": "E-invoice API PoC",
	}

	c.HTML(200, "body", obj)
}

func HandlerEinvoicePost(c *gin.Context) {
	c.HTML(200, "body", nil)
}
