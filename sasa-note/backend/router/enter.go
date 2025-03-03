package router

import (
	"backend/router/system"
)

var RouterGroupApp = new(RouterGroup)

type RouterGroup struct {
	System system.RouterGroup
}
