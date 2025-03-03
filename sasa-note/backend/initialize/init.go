package initialize

import (
	"backend/constant"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"log"
	"os"
)

// inital global

func InitConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	// 读取配置文件
	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}
	BandingConstants()
	fmt.Println("Config loaded successfully.")
}

func BandingConstants() {
	constant.ServerNotePort = viper.GetString("server.note.port")
}

// 初始化日志
func InitLogger() *log.Logger {
	logFile, err := os.OpenFile("gin.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		panic(fmt.Errorf("failed to open log file: %w", err))
	}
	return log.New(logFile, "[GIN] ", log.LstdFlags|log.Lshortfile)

}
func printHello() {
	fmt.Println("Hello	this my way")
}

// 初始化gin 路由
func InitRouter(routeName string) *gin.Engine {
	r := gin.Default()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	//r.GET("/test")
	if routeName == constant.NOTE_SERVER_NAME {
		InitNoteRouter(r)
	} else if routeName == constant.STORAGE_SERVER_NAME {

	} else {
		return nil
	}
	return r
}
