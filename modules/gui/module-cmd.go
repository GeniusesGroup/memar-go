/* For license and copyright information please see the LEGAL file in the code repository */

package module

import (
	function_p "memar/computer/function/protocol"
	"memar/computer/datatype"
	datatype_p "memar/computer/datatype/protocol"
	error_p "memar/process/error/protocol"
	cmd "memar/process/command"
	command_p "memar/process/command/protocol"
	operation_p "memar/process/operation/protocol"
	"memar/process/service"
	"memar/protocol"

	"memar/modules/module/services/new"
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
	// err = s.MT.Init("html.spec.whatwg.org")

	// new.Service.Init(s)

	s.Command.Init(parent,
		&new.Service,
	)

	return
}

//memar:impl memar/protocol.MediaType
func (s *command) FileExtension() string                  { return "" }
func (s *command) LifeCycle() datatype_p.LifeCycle        { return datatype_p.LifeCycle_Unset }
func (s *command) ReferenceURI() string                   { return "" }
func (s *command) IssueDate() string                      { return "" }
func (s *command) ExpiryDate() string                     { return "" }
func (s *command) ExpireInFavorOf() mediatype_p.MediaType { return nil }

//memar:impl memar/process/command/protocol.Object
func (s *command) Fields() []protocol.DataType  { return nil }
func (s *command) Methods() []function_p.Method { return nil }

//memar:impl memar/protocol.Service
func (s *command) URI() string                        { return "" }
func (s *command) Priority() operation_p.Priority     { return operation_p.Priority_Unset }
func (s *command) Weight() operation_p.Weight         { return operation_p.Weight_Unset }
func (s *command) ActionType() operation_p.ActionType { return operation_p.ActionType_None }

//memar:impl memar/protocol.ServiceDetails
func (s *command) Request() protocol.DataType  { return nil }
func (s *command) Response() protocol.DataType { return nil }

//memar:impl memar/protocol.Quiddity
func (s *command) Name() string         { return "gui" }
func (c *command) Abbreviation() string { return "" }
func (s *command) Aliases() []string    { return []string{"ui"} }

//memar:impl memar/process/command/protocol.Command
func (s *command) RunnableCommand() bool { return false }
func (s *command) ServeCLA(arguments []string) (err error_p.Error) {
	// err = cmd.ServeCLA(s, arguments)
	return
}
