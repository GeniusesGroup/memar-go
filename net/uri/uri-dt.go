/* For license and copyright information please see the LEGAL file in the code repository */

package uri

import (
	datatype_p "memar/computer/datatype/protocol"
)

//memar:impl memar/identifier/mediatype/protocol.Field_MediaType
func (u *URI) MediaType() string { return "application/uri" } // application/x-www-form-urlencoded

//memar:impl memar/protocol.FileExtension
func (u *URI) FileExtension() string { return "uri" }

//memar:impl memar/datatype/protocol.DataType_Details
func (u *URI) LifeCycle() datatype_p.LifeCycle { return datatype_p.LifeCycle_PreAlpha }
func (u *URI) ReferenceURI() string {
	return "https://www.iana.org/assignments/media-types/application/http"
}
func (u *URI) IssueDate() string                    { return "" }
func (u *URI) ExpiryDate() string                   { return "" }
func (u *URI) ExpireInFavorOf() datatype_p.DataType { return nil }
