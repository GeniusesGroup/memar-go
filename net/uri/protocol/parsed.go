/* For license and copyright information please see the LEGAL file in the code repository */

package uri_p

type Field_Parsed interface {
	Field_Scheme
	Field_Authority
	Field_Path
	Field_Query
	Field_Fragment
}
