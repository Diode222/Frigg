package manager

import (
	"github.com/Diode222/etcd_service_discovery/etcdservice"
	"sync"
)

var serviceManager *etcdservice.ServiceManager
var serviceMangerOnce sync.Once

func GetServiceManger(etcdAddr string) *etcdservice.ServiceManager {
	serviceMangerOnce.Do(func() {
		serviceManager = etcdservice.NewServiceManager(etcdAddr)
	})
	return serviceManager
}
