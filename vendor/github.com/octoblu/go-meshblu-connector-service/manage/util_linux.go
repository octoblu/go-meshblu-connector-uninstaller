package manage

import (
	"fmt"
	"log"
	"path/filepath"
)

func deregisterUserLogin(localAppData, uuid string) error {
	return fmt.Errorf("UserLogin is windows only, not supported on linux")
}

func removeUserLoginDirectories(localAppData, uuid string) error {
	return fmt.Errorf("UserLogin is windows only, not supported on Linux")
}

func serviceName(uuid string) string {
	return fmt.Sprintf("MeshbluConnector-%s", uuid)
}

func serviceDirectory() string {
	return filepath.Join("/opt", "MeshbluConnectors")
}

func userServiceDirectory(homeDir string) string {
	log.Fatalln("userServiceDirectory not available in Linux, only Windows/macOS")
	return ""
}
