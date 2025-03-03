package system

import (
	"backend/middleware"
	"github.com/gin-gonic/gin"
)

type ContentRouter struct{}

func (s *ContentRouter) InitContentRouter(Router *gin.RouterGroup) {
	contentRouter := Router.Group("content").Use(middleware.OperationRecod())
	{
		contentRouter.POST("insert", contentApi.AddContent)   // 新建笔记
		contentRouter.GET("list", contentApi.ListContent)     // 获取笔记列表
		contentRouter.GET("getone", contentApi.GetContent)    // 获取单个笔记
		contentRouter.PUT("update", contentApi.UpdateContent) // 更新笔记
		contentRouter.DELETE("delete")                        // 删除笔记
		contentRouter.GET("share")                            // 分享笔记
		contentRouter.GET("export")                           // 导出笔记
		contentRouter.GET("import")                           // 导入笔记
	}

}
