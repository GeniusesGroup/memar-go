/* For license and copyright information please see the LEGAL file in the code repository */

package generator

import (
	"text/template"
)

var temp_file_ModuleCMD = template.Must(template.New("temp_file_ModuleCMD").Parse(
	`//go:build cmd

/* For license and copyright information please see the LEGAL file in the code repository */

package {{.PackageName}}

import (
	function_p "memar/computer/function/protocol"
	cmd "memar/process/command"
	"memar/computer/datatype"
	"memar/protocol"
	"memar/process/service"
)

var Command command

type command struct {
	datatype.DataType
	service.Service
	cmd.Command
}

// Init use to initialize all services and register them in the given parent.
//
//memar:impl memar/computer/capsule/protocol.LifeCycle
func (s *command) Init(parent command_p.Command) (err error_p.Error) {
	// TODO::: add new module service here

	s.Command.Init(parent,
		// TODO::: add new module service here
	)

	return
}

//memar:impl memar/protocol.MediaType
func (s *command) MediaType() string { return "" }

//memar:impl memar/datatype/protocol.DataType_Details
func (s *command) LifeCycle() datatype_p.LifeCycle    { return datatype_p.LifeCycle_Unset }
func (s *command) ReferenceURI() string               { return "" }
func (s *command) IssueDate() string                  { return "" }
func (s *command) ExpiryDate() string                 { return "" }
func (s *command) ExpireInFavorOf() datatype_p.DataType { return nil }

//memar:impl memar/process/command/protocol.Object
func (s *command) Fields() []protocol.DataType   { return nil }
func (s *command) Methods() []function_p.Method { return nil }

//memar:impl memar/protocol.Service
func (s *command) URI() string                         { return "" }
func (s *command) Priority() operation_p.Priority      { return operation_p.Priority_Unset }
func (s *command) Weight() operation_p.Weight          { return operation_p.Weight_Unset }
func (s *command) ActionType() operation_p.ActionType  { return operation_p.ActionType_None }
func (s *command) UserType() user_p.Type               { return user_p.Type_Unset }

//memar:impl memar/protocol.ServiceDetails
func (s *command) Request() protocol.DataType  { return nil }
func (s *command) Response() protocol.DataType { return nil }

//memar:impl memar/protocol.Quiddity
func (s *command) Name() string         { return "" }
func (c *command) Abbreviation() string { return "" }
func (s *command) Aliases() []string    { return []string{} }

//memar:impl memar/process/command/protocol.Command
func (s *command) RunnableCommand() bool { return false }
func (s *command) ServeCLA(arguments []string) (err error_p.Error) {
	err = cmd.ServeCLA(s, arguments)
	return
}	
`))

var temp_file_ModuleGUI = template.Must(template.New("temp_file_ModuleGUI").Parse(
	`//go:build gui

/* For license and copyright information please see the LEGAL file in the code repository */

package {{.PackageName}}

import (
	"memar/protocol"
)

// RegisterServices will register services in the application.
func RegisterServices() {
	services.Register(&)
}

func RegisterPages() {
	pages.Register(&)
}

func RegisterWidgets() {
	widgets.Register(&)
}
`))

var temp_file_ModuleServer = template.Must(template.New("temp_file_ModuleServer").Parse(
	`//go:build server

/* For license and copyright information please see the LEGAL file in the code repository */

package {{.PackageName}}

import (
	"memar/protocol"
)

// RegisterServices will register services in the application.
func RegisterServices() {
	services.Register(&)
}
`))

var temp_file_Module = template.Must(template.New("temp_file_Module").Parse(
	`/* For license and copyright information please see the LEGAL file in the code repository */

package {{.PackageName}}

import (
	"memar/module"
	"memar/protocol"
)

var ModuleName = moduleName{
	Name: "",
}

type moduleName struct {
	module.Module
}

func init() {
	modules.Register(&ModuleName)
}
`))

var temp_file_LocaleENG = template.Must(template.New("temp_file_LocaleENG").Parse(
	`//go:build lang_eng

/* For license and copyright information please see the LEGAL file in the code repository */

package {{.PackageName}}

import (
	"memar/protocol"
)

const (
	domainEnglish = "module"
)

func init() {
	Command.Service.SetDetail(lang_p.LanguageEnglish,
		domainEnglish,
		"",
		"",
		"",
		"",
		nil)
}
`))

var temp_file_Protocol_DataModel = template.Must(template.New("temp_file_Protocol_DataModel").Parse(
	`/* For license and copyright information please see the LEGAL file in the code repository */

package {{.PackageName}}

import (
	"memar/protocol"
)

// A [DataModel](https://en.wikipedia.org/wiki/Data_model) is an abstract model that organizes elements of data (storage fields) and
// standardizes how they relate to one another modules and to the properties of real-world entities.
// It is related to [data modeling](https://en.wikipedia.org/wiki/Data_modeling) process in software engineering
// It is not whole [Domain_model](https://en.wikipedia.org/wiki/Domain_model) because incorporate just data, not behavior.
// It is not a [storage model](https://en.wikipedia.org/wiki/Storage_model) that is a model that captures key physical aspects of data structure in a data store.
type DataModel interface {
	SaveTime() time_p.Time
}
`))
