package service

import "backend/service/system"

var ServiceGroupApp = new(ServiceGroup)

type ServiceGroup struct {
	SystemServiceGroup system.ServiceGroup
}
