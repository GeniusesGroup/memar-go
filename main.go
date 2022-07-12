/* For license and copyright information please see LEGAL file in repository */

package main

import (
	"flag"
	"fmt"
	"os"
	"time"
	// libgo "./services"
)

func main() {
	var start = time.Now()

	var serviceName string
	if len(os.Args) > 1 {
		serviceName = os.Args[1]
	}

	// https://wails.app/reference/cli/
	switch serviceName {
	case "":
		helpMessage()
		os.Exit(2)
	case "version":
		fmt.Println("Version:", version, `
You may help us and create issue:
https://github.com/GeniusesGroup/libgo/issues/new
For more information, see:
https://github.com/GeniusesGroup/libgo/`)
		os.Exit(0)
	case "help":
		helpMessage()
		os.Exit(0)
	case "update":
		// TODO::: update libgo
	case "interactive":
		// TODO::: cli interactive mode
	case "new": // || "init" || "create"
		// TODO::: make new project by tempelate address e.g. github url or defaults one:
		// - raw(with little predefined codes)
		// - society(government)
		// - org(organization)
		// - aggregator-org
		// - game(or sub type of org??)
		// Ask version control: libgo || git

		// # Make new project base on libgo with git as version control
		// # - Initialize new project by `git init`
		// # - Clone exiting repo by `git clone ${repository path}`.
		// # - Add libgo to project as submodule by `git submodule add -b master https://github.com/GeniusesGroup/libgo`
		// #  `//go:generate libgo -destination=../mock.go -package=store_mock  -self_package=github.com// github.com/// UserStore`
	case "build":
		// TODO::: generate and build new project as OS image(Unikernel) can run on PersiaOS, KVM, XEN, ...
	case "deploy":
		// TODO::: build and deploy new project to a first server just if it is new project by not raw template
	case "init":
		// TODO:::
	case "issue":
		// TODO::: This command speeds up the process for submitting an issue
	case "generate":
		// TODO::: run generators on the projects without build
	case "gui":
		// TODO::: GUI services like make new page
	case "???":
		// libgo.CompleteJson()
		//go:generate libgo -service=complete-json -force_update -destination=../service.go -package=services
		//go:generate libgo -service=gui -os ios -certificate "Apple Distribution" -profile "My App" -appID "com.example.myapp"
	default:
		fmt.Fprintf(os.Stderr, "libgo %s: unknown command\nRun 'libgo help' for usage.\n", serviceName)
		os.Exit(2)
	}

	fmt.Println("*************************************** Libgo ***************************************")
	fmt.Println("Version: ", version)
	fmt.Println("Start at: ", start)
	fmt.Println("Running duration: ", time.Since(start))
	fmt.Println("See you soon!")
	fmt.Println("-------------------------------------------------------------------------------------")
}

func helpMessage() {
	fmt.Fprintf(flag.CommandLine.Output(), "Libgo is a tool to develope software in Go by the full details Geniuses.Group architecture\n\n")
	fmt.Printf("If no command is provided %s will start the runner with the provided flags\n\n", os.Args[0])
	fmt.Println("Commands:")
	fmt.Println("	init	creates")

	fmt.Println("Flags:")
	flag.PrintDefaults()
}
