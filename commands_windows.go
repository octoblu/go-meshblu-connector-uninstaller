package main

import (
	"fmt"
	"os"

	"github.com/fatih/color"
	"github.com/octoblu/go-meshblu-connector-service/manage"
	"github.com/urfave/cli"
)

func userLoginOpts(context *cli.Context) (string, string) {
	uuid := context.Args().Get(0)
	localAppData := context.String("local-app-data")

	if uuid == "" || localAppData == "" {
		cli.ShowCommandHelp(context, "user-service")

		fmt.Println()

		if localAppData == "" {
			color.Red("Missing required option --local-app-data, -l, env: LOCALAPPDATA")
		}
		if uuid == "" {
			color.Red("Missing required argument <uuid>")
		}
		os.Exit(1)
	}

	return uuid, localAppData
}

// UserLogin uninstalls a connector installed as UserLogin
func UserLogin(context *cli.Context) {
	uuid, localAppData := userLoginOpts(context)
	err := manage.UninstallUserLogin(&manage.UninstallUserLoginOptions{UUID: uuid, LocalAppData: localAppData})
	die(err)
}
