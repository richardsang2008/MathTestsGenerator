package controllers

import "github.com/gin-gonic/gin"

type PinController struct {
}

func (r *PinController) Pinhandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}
