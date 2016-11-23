package manage

// UninstallServiceOptions are used to call UninstallService
type UninstallServiceOptions struct {
	// UUID is the UUID of the connector to uninstall
	UUID string
}

// UninstallService uninstalls the meshblu-connector
// indicated by the uuid
// * Stop/deregister the service
// * Remove all the directories
func UninstallService(options *UninstallServiceOptions) error {
	err := deregisterService(options.UUID)
	if err != nil {
		debug("Service deregistration failed, probably wasn't installed: %v", err.Error())
	}

	return removeServiceDirectories(options.UUID)
}
