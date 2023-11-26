/* For license and copyright information please see the LEGAL file in the code repository */

package module

import (
	cmd "memar/command"
	"memar/mediatype"
	"memar/protocol"
	"memar/service"

	"memar/modules/module/services/new"
)

var Command command

type command struct {
	mediatype.MT
	service.Service
	cmd.Command
}

// Init use to initialize all services and register them in the given parent.
//
//libgo:impl memar/protocol.ObjectLifeCycle
func (s *command) Init(parent protocol.Command) (err protocol.Error) {
	// err = s.MT.Init("html.spec.whatwg.org")

	// new.Service.Init(s)

	s.Command.Init(parent,
		&new.Service,
	)

	return
}

//libgo:impl memar/protocol.MediaType
func (s *command) FileExtension() string               { return "" }
func (s *command) Status() protocol.SoftwareStatus     { return protocol.Software_Unset }
func (s *command) ReferenceURI() string                { return "" }
func (s *command) IssueDate() protocol.Time            { return nil }
func (s *command) ExpiryDate() protocol.Time           { return nil }
func (s *command) ExpireInFavorOf() protocol.MediaType { return nil }

//libgo:impl memar/protocol.Object
func (s *command) Fields() []protocol.DataType         { return nil }
func (s *command) Methods() []protocol.DataType_Method { return nil }

//libgo:impl memar/protocol.Service
func (s *command) URI() string                 { return "" }
func (s *command) Priority() protocol.Priority { return protocol.Priority_Unset }
func (s *command) Weight() protocol.Weight     { return protocol.Weight_Unset }
func (s *command) CRUDType() protocol.CRUD     { return protocol.CRUD_None }
func (s *command) UserType() protocol.UserType { return protocol.UserType_Unset }

//libgo:impl memar/protocol.ServiceDetails
func (s *command) Request() protocol.DataType  { return nil }
func (s *command) Response() protocol.DataType { return nil }

//libgo:impl memar/protocol.Quiddity
func (s *command) Name() string         { return "gui" }
func (c *command) Abbreviation() string { return "" }
func (s *command) Aliases() []string    { return []string{"ui"} }

//libgo:impl memar/protocol.Command
func (s *command) Runnable() bool { return false }
func (s *command) ServeCLA(arguments []string) (err protocol.Error) {
	// err = cmd.ServeCLA(s, arguments)
	return
}
