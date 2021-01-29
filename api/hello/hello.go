package hello

import (
	"example.com/someservice555/internal/routing"
	"github.com/gin-gonic/gin"
	"net/http"
)

const Path = "/hello"
const Method = routing.GET

func ProcessRequest(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"msg": "Hello!"})
}