package api

import (
	"backend/api/system"
)

var ApiGroupApp = new(ApiGroup)

type ApiGroup struct {
	SystemApiGroup system.ApiGroup
}
