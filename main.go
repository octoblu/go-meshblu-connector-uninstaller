package main

import (
	"fmt"
	"log"
	"os"

	"github.com/coreos/go-semver/semver"
	"github.com/fatih/color"
	"github.com/octoblu/go-meshblu-connector-service/manage"
	"github.com/urfave/cli"
	De "github.com/visionmedia/go-debug"
)

var debug = De.Debug("meshblu-connector-uninstaller:main")

func main() {
	app := cli.NewApp()
	app.Name = "meshblu-connector-uninstaller"
	app.UsageText = "meshblu-connector-uninstaller <command> [options] <uuid>"
	app.Version = version()
	app.Commands = []cli.Command{
		{
			Name:      "service",
			Usage:     "uninstall a connector installed as a Service",
			ArgsUsage: "<uuid>",
			Action:    Service,
			Flags:     []cli.Flag{},
		},
		{
			Name:      "user-service",
			Usage:     "uninstall a connector installed as a UserService",
			ArgsUsage: "<uuid>",
			Action:    UserService,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:   "home-dir, H",
					Usage:  "Home directory of the user.",
					EnvVar: "HOME",
				},
				cli.StringFlag{
					Name:   "username, u",
					Usage:  "System Username",
					EnvVar: "USER",
				},
			},
		},
	}
	app.Flags = []cli.Flag{}
	app.Run(os.Args)
}

// Service uninstalls a connector insalled as a Service
func Service(context *cli.Context) {
	uuid := getServiceOpts(context)
	err := manage.UninstallService(&manage.UninstallServiceOptions{UUID: uuid})
	die(err)
}

// UserService uninstalls a connector installed as a UserService
func UserService(context *cli.Context) {
	homeDir, username, uuid := getUserServiceOpts(context)
	err := manage.UninstallUserService(&manage.UninstallUserServiceOptions{
		UUID:            uuid,
		HomeDir:         homeDir,
		ServiceUsername: username,
	})
	die(err)
}

func anyFlagsMissing(context *cli.Context, flagNames ...string) bool {
	for _, flagName := range flagNames {
		if !context.IsSet(flagName) {
			return false
		}
	}
	return true
}

func die(err error) {
	if err != nil {
		log.Fatalln(err.Error())
	}
	os.Exit(0)
}

func getServiceOpts(context *cli.Context) string {
	uuid := context.Args().Get(0)
	if uuid == "" {
		cli.ShowAppHelp(context)

		fmt.Println()

		if uuid == "" {
			color.Red("Missing required argument <uuid>")
		}
		os.Exit(1)
	}
	return uuid
}

func getUserServiceOpts(context *cli.Context) (string, string, string) {
	homeDir := context.String("home-dir")
	username := context.String("username")
	uuid := context.Args().Get(0)

	if uuid == "" || anyFlagsMissing(context, "home-dir", "username") {
		cli.ShowAppHelp(context)

		fmt.Println()

		if !context.IsSet("home-dir") {
			color.Red("Missing required option --home-dir, -H, env: HOME")
		}
		if !context.IsSet("username") {
			color.Red("Missing required option --username, -u, env: USER")
		}
		if uuid == "" {
			color.Red("Missing required argument <uuid>")
		}
		os.Exit(1)
	}

	return homeDir, username, uuid
}

func version() string {
	version, err := semver.NewVersion(VERSION)
	if err != nil {
		errorMessage := fmt.Sprintf("Error with version number: %v", VERSION)
		log.Panicln(errorMessage, err.Error())
	}
	return version.String()
}
