/* For license and copyright information please see the LEGAL file in the code repository */

package new

import (
	"memar/computer/datatype"
	error_p "memar/process/error/protocol"
	cmd "memar/process/command"
	command_p "memar/process/command/protocol"
	operation_p "memar/process/operation/protocol"
	"memar/process/service"
	"memar/protocol"
	user_p "memar/user/protocol"
)

var Service serv

type serv struct {
	datatype.DataType
	service.Service
	cmd.Command
}

// Init use to initialize all services and register them in the given parent.
//
//memar:impl memar/computer/capsule/protocol.LifeCycle
func (s *serv) Init(parent command_p.Command) (err error_p.Error) {
	err = s.DataType.Init("domain/memar.scm.geniuses.group; module=module; type=service; name=new")
	err = s.Command.Init(parent)
	return
}

//memar:impl memar/datatype/protocol.DataType_Details
func (s *serv) LifeCycle() datatype_p.LifeCycle      { return datatype_p.LifeCycle_PreAlpha }
func (s *serv) ReferenceURI() string                 { return "" }
func (s *serv) IssueDate() string                    { return "unix/1684654692" }
func (s *serv) ExpiryDate() string                   { return "" }
func (s *serv) ExpireInFavorOf() datatype_p.DataType { return nil }

//memar:impl memar/process/service/protocol.Service
func (s *serv) URI() string                        { return "" }
func (s *serv) ActionType() operation_p.ActionType { return operation_p.ActionType_Create }
func (s *serv) UserType() user_p.Type              { return user_p.Type_Unset }

//memar:impl memar/process/operation/protocol.Importance
func (s *serv) Priority() operation_p.Priority { return operation_p.Priority_Lowest }
func (s *serv) Weight() operation_p.Weight     { return operation_p.Weight_Lowest }

//memar:impl memar/protocol.ServiceDetails
func (s *serv) Request() protocol.DataType  { return nil }
func (s *serv) Response() protocol.DataType { return nil }

//memar:impl memar/protocol.Quiddity
func (s *serv) Name() string         { return "new" }
func (s *serv) Abbreviation() string { return "" }
func (s *serv) Aliases() []string    { return []string{} }

//memar:impl memar/process/command/protocol.Command
func (s *serv) RunnableCommand() bool { return true }

type slice struct {
	buf uintptr
}

func (s *slice) iterate() {}
