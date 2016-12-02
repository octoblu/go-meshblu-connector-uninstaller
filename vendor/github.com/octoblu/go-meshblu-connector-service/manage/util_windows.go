package manage

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"time"

	"golang.org/x/sys/windows/registry"
)

func deregisterUserLogin(localAppData, uuid string) error {
	err := killUserLoginProcess(localAppData, uuid)
	if err != nil {
		return err
	}

	key, err := registry.OpenKey(registry.CURRENT_USER, `Software\Microsoft\Windows\CurrentVersion\Run`, registry.WRITE)
	if err != nil {
		return err
	}

	err = key.DeleteValue(serviceName(uuid))
	if err != nil {
		return err
	}

	return key.Close()
}

func killUserLoginProcess(localAppData, uuid string) error {
	pid, err := userLoginPID(localAppData, uuid)
	if err != nil {
		return err
	}

	if pid == -1 {
		return nil
	}

	process, err := os.FindProcess(pid)
	if err != nil {
		return err
	}

	return process.Kill()
}

func removeUserLoginDirectories(localAppData, uuid string) error {
	var err error

	for i := 0; i < 10; i++ {
		err = os.RemoveAll(connectorUserLoginDirectory(localAppData, uuid))
		debug("removeUserLoginDirectories attempt %v/10: %v", i+1, err)
		if _, isPathError := err.(*os.PathError); isPathError {
			debug("is a PathError, will try again in 10s")
			time.Sleep(10 * time.Second)
			continue
		}

		debug("Am not trying again because the error was: %v", err)
		return err
	}

	return fmt.Errorf("Failed to remove directories: %v", err.Error())
}

func serviceName(uuid string) string {
	return fmt.Sprintf("MeshbluConnector-%s", uuid)
}

func serviceDirectory() string {
	programFilesDir := os.Getenv("PROGRAMFILESX86")
	if programFilesDir == "" {
		programFilesDir = os.Getenv("PROGRAMFILES")
	}
	return filepath.Join(programFilesDir, "MeshbluConnectors")
}

func userLoginPID(localAppData, uuid string) (int, error) {
	data, err := ioutil.ReadFile(filepath.Join(connectorUserLoginDirectory(localAppData, uuid), "update.json"))
	if err != nil && os.IsNotExist(err) {
		return -1, nil
	}
	if err != nil {
		return -1, err
	}

	parsed := struct{ Pid int }{Pid: -1}
	err = json.Unmarshal(data, &parsed)
	if err != nil {
		return -1, err
	}

	return parsed.Pid, nil
}

func userServiceDirectory(homeDir string) string {
	return filepath.Join(os.Getenv("LOCALAPPDATA"), "MeshbluConnectors")
}

func connectorUserLoginDirectory(localAppData, uuid string) string {
	return filepath.Join(localAppData, "MeshbluConnectors", uuid)
}
