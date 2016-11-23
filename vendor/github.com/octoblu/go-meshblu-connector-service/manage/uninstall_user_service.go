package manage

// UninstallUserServiceOptions are used to call UninstallUserService
type UninstallUserServiceOptions struct {
	// HomeDir is the HomeDir of user
	HomeDir string

	// UUID is the UUID of the connector to uninstall
	UUID string

	// ServiceUsername is the name of the user the service is registered
	// under
	ServiceUsername string

	// ServicePassword is a password of some kind. I think it's only required if you set the
	// user to someone other than the one executing this process, but who knows.
	ServicePassword string
}

// UninstallUserService uninstalls the meshblu-connector
// indicated by the uuid
// * Stop/deregister the service
// * Remove all the directories
func UninstallUserService(options *UninstallUserServiceOptions) error {
	err := deregisterUserService(options.UUID, options.ServiceUsername, options.ServicePassword)
	if err != nil {
		debug("UserService deregistration failed, probably wasn't installed: %v", err.Error())
	}

	return removeUserServiceDirectories(options.HomeDir, options.UUID)
}
