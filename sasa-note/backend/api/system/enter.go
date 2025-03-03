package system

import (
	"backend/service"
)

type ApiGroup struct {
	ContentApi
	DBApi
}

var (
	contentSerice = service.ServiceGroupApp.SystemServiceGroup.ContentService
	//initDBService = service.ServiceGroupApp.SystemServiceGroup.InitDBService
)
