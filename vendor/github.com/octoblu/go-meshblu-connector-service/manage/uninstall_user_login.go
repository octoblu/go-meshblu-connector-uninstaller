package manage

// UninstallUserLoginOptions are used to call UninstallUserLogin
type UninstallUserLoginOptions struct {
	// UUID is the UUID of the connector to uninstall
	UUID string

	// LocalAppData is the user's ~/AppData/Local directory
	LocalAppData string
}

// UninstallUserLogin uninstalls the meshblu-connector
// indicated by the uuid
// * Stop/deregister the service
// * Remove all the directories
func UninstallUserLogin(options *UninstallUserLoginOptions) error {
	err := deregisterUserLogin(options.LocalAppData, options.UUID)
	if err != nil {
		debug("Service deregistration failed, probably wasn't installed: %v", err.Error())
	}

	return removeUserLoginDirectories(options.LocalAppData, options.UUID)
}
