package initialize

import (
	"backend/constant"
	"backend/router"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func InitNoteRouter(r *gin.Engine) *gin.Engine {
	systemRouter := router.RouterGroupApp.System

	PublicGroup := r.Group("")
	PrivateGroup := r.Group("")

	// health check
	PublicGroup.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, "ok")
	})

	systemRouter.InitBaseRouter(PublicGroup)
	systemRouter.InitInitRouter(PublicGroup) // 自动初始化相关
	systemRouter.InitContentRouter(PrivateGroup)
	err := r.Run(":" + constant.ServerNotePort) // listen and serve on 0.0.0.0:8080

	if err != nil {
		log.Fatalf("server run error: %v", err)
		panic(err)
	}

	return r
}
