package system

// 内容serivce
import (
	"backend/api"
	"fmt"
	"github.com/gin-gonic/gin"
)

type RouterGroup struct {
	ContentRouter
	BaseRouter
	InitRouter
}

var (
	contentApi = api.ApiGroupApp.SystemApiGroup.ContentApi
	dbApi      = api.ApiGroupApp.SystemApiGroup.DBApi
)

type BaseRouter struct{}

func (s *BaseRouter) InitBaseRouter(Router *gin.RouterGroup) (R gin.IRoutes) {
	baseRouter := Router.Group("base")
	{
		baseRouter.GET("login", func(context *gin.Context) {
			fmt.Println("loginloginlogin")
		})
	}
	return baseRouter
}

type InitRouter struct{}

func (s *InitRouter) InitInitRouter(Router *gin.RouterGroup) {
	initRouter := Router.Group("init")
	{
		initRouter.POST("initdb", dbApi.InitDB)   // 初始化数据库
		initRouter.POST("checkdb", dbApi.CheckDB) // 检测是否需要初始化数据库
	}
}
