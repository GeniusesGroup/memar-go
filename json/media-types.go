/* For license and copyright information please see the LEGAL file in the code repository */

package json

import (
	"libgo/detail"
	"libgo/mediatype"
	"libgo/protocol"
)

var MediaType mediaType

func init() {
	MediaType.Init("application/json")
}

type mediaType struct {
	detail.DS
	mediatype.MT
}

//libgo:impl libgo/protocol.MediaType
func (s *mediaType) FileExtension() string               { return "json" }
func (s *mediaType) Status() protocol.SoftwareStatus     { return protocol.Software_StableRelease }
func (s *mediaType) ReferenceURI() string                { return "" }
func (s *mediaType) IssueDate() protocol.Time            { return nil }
func (s *mediaType) ExpiryDate() protocol.Time           { return nil }
func (s *mediaType) ExpireInFavorOf() protocol.MediaType { return nil }

//libgo:impl libgo/protocol.Object
func (s *mediaType) Fields() []protocol.Object_Member_Field   { return nil }
func (s *mediaType) Methods() []protocol.Object_Member_Method { return nil }
