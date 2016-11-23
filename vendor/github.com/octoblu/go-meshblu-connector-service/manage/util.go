package manage

import (
	"os"
	"path/filepath"

	"github.com/kardianos/service"
	De "github.com/tj/go-debug"
)

var debug = De.Debug("meshblu-connector-service:manage")

func connectorUserServiceDirectory(homeDir, uuid string) string {
	return filepath.Join(userServiceDirectory(homeDir), uuid)
}

func connectorServiceDirectory(uuid string) string {
	return filepath.Join(serviceDirectory(), uuid)
}

func deregisterService(uuid string) error {
	config := &service.Config{
		Name: serviceName(uuid),
		Option: service.KeyValue{
			"UserService": false,
		},
	}

	program := &Program{}
	svc, err := service.New(program, config)
	if err != nil {
		return err
	}

	return svc.Uninstall()
}

func deregisterUserService(uuid, serviceUsername, servicePassword string) error {
	config := &service.Config{
		Name: serviceName(uuid),
		Option: service.KeyValue{
			"UserService": true,
			"UserName":    serviceUsername,
			"Password":    servicePassword,
		},
	}

	program := &Program{}
	svc, err := service.New(program, config)
	if err != nil {
		return err
	}

	return svc.Uninstall()
}

func removeServiceDirectories(uuid string) error {
	return os.RemoveAll(connectorServiceDirectory(uuid))
}

func removeUserServiceDirectories(homeDir, uuid string) error {
	return os.RemoveAll(connectorUserServiceDirectory(homeDir, uuid))
}
