/* For license and copyright information please see the LEGAL file in the code repository */

package uri

import (
	datatype_p "memar/computer/datatype/protocol"
)

//memar:impl memar/identifier/mediatype/protocol.Field_MediaType
func (u *URI[BUF]) MediaType() string { return "application/uri" } // application/x-www-form-urlencoded

//memar:impl memar/protocol.FileExtension
func (u *URI[BUF]) FileExtension() string { return "uri" }

//memar:impl memar/datatype/protocol.DataType_Details
func (u *URI[BUF]) LifeCycle() datatype_p.LifeCycle { return datatype_p.LifeCycle_PreAlpha }
func (u *URI[BUF]) ReferenceURI() string {
	return "https://www.iana.org/assignments/media-types/application/http"
}
func (u *URI[BUF]) IssueDate() string                    { return "" }
func (u *URI[BUF]) ExpiryDate() string                   { return "" }
func (u *URI[BUF]) ExpireInFavorOf() datatype_p.DataType { return nil }
