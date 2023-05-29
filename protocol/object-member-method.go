/* For license and copyright information please see the LEGAL file in the code repository */

package protocol

// Object_Member_Method
type Object_Member_Method interface {
	// We believe fields MUST always access from inside the object,
	// So we MUST have this method just in methods not fields.
	Access() Object_Access

	Blocking() bool // TODO::: be method or as type??
	// TODO::: add more

	Object_Member
}
