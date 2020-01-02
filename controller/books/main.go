package books

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetAll(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"books": "you got books",})
}

func Insert(c *gin.Context)  {
}
