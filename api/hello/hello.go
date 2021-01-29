package hello

import (
	"example.com/someservice555/internal/routing"
	"github.com/gin-gonic/gin"
	"net/http"
)

const Path = "/hello"
const Method = routing.GET

func ProcessRequest(c *gin.Context) {
	name := c.Query("name")
	c.JSON(http.StatusOK, gin.H{"msg": fmt.Sprintf("Hello %v!", name)})
}