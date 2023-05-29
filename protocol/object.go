/* For license and copyright information please see the LEGAL file in the code repository */

package protocol

type Object interface {
	Fields() []Object_Member_Field
	Methods() []Object_Member_Method

	// ObjectLifeCycle
	// Object_Member_Len
	// Details
}
