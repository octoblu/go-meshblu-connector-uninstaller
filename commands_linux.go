package main

import (
	"fmt"

	"github.com/urfave/cli"
)

// UserLogin uninstalls a connector installed as UserLogin
func UserLogin(context *cli.Context) {
	die(fmt.Errorf("UserLogin uninstall is Windows only, and is not supported on Linux."))
}
