package main

import (
	"fmt"
	"os"

	"github.com/dokku/dokku/plugins/common"
	flag "github.com/spf13/pflag"
)

func main() {
	quiet := flag.Bool("quiet", false, "--quiet: set DOKKU_QUIET_OUTPUT=1")
	global := flag.Bool("global", false, "--global: Whether global or app-specific")
	flag.Parse()
	cmd := flag.Arg(0)

	if *quiet {
		os.Setenv("DOKKU_QUIET_OUTPUT", "1")
	}

	var err error
	switch cmd {
	case "compose-up":
		projectName := flag.Arg(1)
		composeFile := flag.Arg(2)
		err = common.ComposeUp(common.ComposeUpInput{
			ProjectName: projectName,
			ComposeFile: composeFile,
		})
	case "compose-down":
		projectName := flag.Arg(1)
		composeFile := flag.Arg(2)
		err = common.ComposeDown(common.ComposeDownInput{
			ProjectName: projectName,
			ComposeFile: composeFile,
		})
	case "copy-dir-from-image":
		appName := flag.Arg(1)
		image := flag.Arg(2)
		source := flag.Arg(3)
		destination := flag.Arg(4)
		err = common.CopyDirFromImage(appName, image, source, destination)
	case "copy-from-image":
		appName := flag.Arg(1)
		image := flag.Arg(2)
		source := flag.Arg(3)
		destination := flag.Arg(4)
		err = common.CopyFromImage(appName, image, source, destination)
	case "docker-cleanup":
		appName := flag.Arg(1)
		force := common.ToBool(flag.Arg(2))
		if *global {
			appName = "--global"
		}
		err = common.DockerCleanup(appName, force)
	case "is-deployed":
		appName := flag.Arg(1)
		if !common.IsDeployed(appName) {
			err = fmt.Errorf("App %v not deployed", appName)
		}
	case "image-is-cnb-based":
		image := flag.Arg(1)
		if common.IsImageCnbBased(image) {
			fmt.Print("true")
		} else {
			fmt.Print("false")
		}
	case "image-is-herokuish-based":
		image := flag.Arg(1)
		appName := flag.Arg(2)
		if common.IsImageHerokuishBased(image, appName) {
			fmt.Print("true")
		} else {
			fmt.Print("false")
		}
	case "scheduler-detect":
		appName := flag.Arg(1)
		if *global {
			appName = "--global"
		}
		fmt.Print(common.GetAppScheduler(appName))
	case "plugn-trigger-exists":
		triggerName := flag.Arg(1)
		if common.PlugnTriggerExists(triggerName) {
			fmt.Print("true")
		} else {
			fmt.Print("false")
		}
	case "verify-app-name":
		appName := flag.Arg(1)
		err = common.VerifyAppName(appName)
	default:
		err = fmt.Errorf("Invalid common command call: %v", cmd)
	}

	if err != nil {
		common.LogFailWithErrorQuiet(err)
	}
}
