package main

import (
	"backend/constant"
	"backend/initialize"
)

func main() {
	initialize.InitConfig()
	initialize.InitLogger()
	initialize.InitRouter(constant.USER_SERVER_NAME)

}
