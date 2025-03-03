package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func OperationRecod() gin.HandlerFunc {
	return func(c *gin.Context) {
		// record operation
		fmt.Println(c.Errors)
	}
}
