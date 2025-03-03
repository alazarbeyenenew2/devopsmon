package handler

import "github.com/gin-gonic/gin"

type Scanner interface {
	PortScanner(c *gin.Context)
}
