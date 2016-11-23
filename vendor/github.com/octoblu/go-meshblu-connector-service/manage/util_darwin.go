package manage

import (
	"fmt"
	"path/filepath"
)

func deregisterUserLogin(localAppData, uuid string) error {
	return fmt.Errorf("UserLogin is windows only, not supported on macOS")
}

func removeUserLoginDirectories(localAppData, uuid string) error {
	return fmt.Errorf("UserLogin is windows only, not supported on macOS")
}

func serviceName(uuid string) string {
	return fmt.Sprintf("com.octoblu.%s", uuid)
}

func serviceDirectory() string {
	return filepath.Join("/Library", "MeshbluConnectors")
}

func userServiceDirectory(homeDir string) string {
	return filepath.Join(homeDir, "Library", "MeshbluConnectors")
}
