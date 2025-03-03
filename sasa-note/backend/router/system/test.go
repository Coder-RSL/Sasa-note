package system

import (
	"backend/middleware"
	"github.com/gin-gonic/gin"
)

type TestRouter struct{}

func (s *TestRouter) InitContentRouter(Router *gin.RouterGroup) {
	contentRouter := Router.Group("test").Use(middleware.OperationRecod())
	{
		contentRouter.GET("hello", contentApi.AddContent) // 新建笔记
	}

}
