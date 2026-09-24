/* For license and copyright information please see the LEGAL file in the code repository */

package modules

import (
	"fmt"
	"os"
	"time"

	"memar/computer/datatype"
	datatype_p "memar/computer/datatype/protocol"
	error_p "memar/process/error/protocol"
	cmd "memar/process/command"
	"memar/process/service"

	// "memar/modules/app"
	// "memar/modules/json"
	"memar/modules/module"
)

var RootCommand rootCommand

type rootCommand struct {
	datatype.DataType
	service.Service
	cmd.Command
}

// Init use to initialize all services and register them in the RootCommand.
//
//memar:impl memar/computer/capsule/protocol.LifeCycle
func (s *rootCommand) Init() (err error_p.Error) {
	// app.Command.Init(s)
	// json.Command.Init(s)
	module.Command.Init(s)

	s.Command.Init(nil,
		// &app.Command,
		&module.Command,
		// &json.Command
	)
	return
}
func (s *rootCommand) Reinit() (err error_p.Error) {
	return
}
func (s *rootCommand) Deinit() (err error_p.Error) {
	return
}

//memar:impl memar/datatype/protocol.DataType_Details
func (s *rootCommand) LifeCycle() datatype_p.LifeCycle      { return datatype_p.LifeCycle_PreAlpha }
func (s *rootCommand) ReferenceURI() string                 { return "" }
func (s *rootCommand) IssueDate() string                    { return "" }
func (s *rootCommand) ExpiryDate() string                   { return "" }
func (s *rootCommand) ExpireInFavorOf() datatype_p.DataType { return nil }

//memar:impl memar/protocol.ServiceDetails
func (s *rootCommand) Request() datatype_p.DataType  { return nil }
func (s *rootCommand) Response() datatype_p.DataType { return nil }

//memar:impl memar/protocol.Quiddity
func (s *rootCommand) Name() string         { return "memar" }
func (c *rootCommand) Abbreviation() string { return "" }
func (s *rootCommand) Aliases() []string    { return []string{} }

//memar:impl memar/process/command/protocol.Command
func (s *rootCommand) RunnableCommand() bool { return false }

//memar:impl memar/process/command/protocol.CommandHandler
func (s *rootCommand) ServeCLA(arguments []string) (err error_p.Error) {
	var start = time.Now()

	err = cmd.ServeCLA(s, arguments)
	if err != nil {
		fmt.Fprintln(os.Stderr, "\n\nError:\n\n", err)
	} else {
		fmt.Println("*************************************** Memar (Golang) ***************************************")
		fmt.Println("Version: ", version)
		fmt.Println("Start at: ", start)
		fmt.Println("Working directory: ", cmd.Cwd())
		fmt.Println("Running duration: ", time.Since(start))
		fmt.Println("You may help us and create issue:\n	https://github.com/GeniusesGroup/memar/issues/new")
		fmt.Println("For more information, see:\n	https://Geniuses.Group/scm?repo=memar-go")
		fmt.Println("See you soon!")
		fmt.Println("-------------------------------------------------------------------------------------")
	}
	return
}
