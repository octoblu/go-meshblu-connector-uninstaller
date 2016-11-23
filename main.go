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

var debug = De.Debug("go-meshblu-connector-uninstaller:main")

func main() {
	app := cli.NewApp()
	app.Name = "go-meshblu-connector-uninstaller"
	app.UsageText = "go-meshblu-connector-uninstaller [options] <uuid>"
	app.Version = version()
	app.Action = run
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:   "home-dir, H",
			Usage:  "Home directory of the user",
			EnvVar: "HOME",
		},
		cli.StringFlag{
			Name:   "username, u",
			Usage:  "System Username",
			EnvVar: "USER",
		},
	}
	app.Run(os.Args)
}

func run(context *cli.Context) {
	homeDir, username, uuid := getOpts(context)
	err := manage.Uninstall(&manage.UninstallOptions{
		UUID:            uuid,
		HomeDir:         homeDir,
		ServiceUsername: username,
	})
	if err != nil {
		log.Fatalln(err.Error())
	}
	os.Exit(0)
}

func getOpts(context *cli.Context) (string, string, string) {
	homeDir := context.String("home-dir")
	username := context.String("username")
	uuid := context.Args().Get(0)

	if homeDir == "" || username == "" || uuid == "" {
		cli.ShowAppHelp(context)

		if homeDir == "" {
			color.Red("  Missing required option --home-dir, -H, env: HOME")
		}
		if username == "" {
			color.Red("  Missing required option --username, -u, env: USER")
		}
		if uuid == "" {
			color.Red("  Missing required argument <uuid>")
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
