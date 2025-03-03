package main

import (
	"backend/constant"
	"backend/initialize"
)

func main() {
	initialize.InitConfig()
	initialize.InitDB()
	initialize.InitLogger()
	initialize.InitRouter(constant.NOTE_SERVER_NAME)
}
