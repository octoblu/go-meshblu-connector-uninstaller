package manage

import (
	"fmt"
	"log"
	"path/filepath"
)

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
