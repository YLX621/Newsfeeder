package handler

import(
	"github.com/gin-gonic/gin"
	"net/http"
)

func NewfeedPost() gin.HandlerFunc{
	return func(c *gin.Context){
		c.JSON(http.StatusOK,map[string]string{
			"hello":"found me",
		})
	}
}