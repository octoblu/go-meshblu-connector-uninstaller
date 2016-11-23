package manage

import (
	"fmt"
	"path/filepath"
)

func serviceName(uuid string) string {
	return fmt.Sprintf("com.octoblu.%s", uuid)
}

func serviceDirectory() string {
	return filepath.Join("/Library", "MeshbluConnectors")
}

func userServiceDirectory(homeDir string) string {
	return filepath.Join(homeDir, "Library", "MeshbluConnectors")
}
