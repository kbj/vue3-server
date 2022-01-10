package service

import "vue3-server/service/system"

type Service struct {
	System system.ServiceGroup
}

var ServiceApp = new(Service)
