/* For license and copyright information please see the LEGAL file in the code repository */

package capsule_p

type Implement interface {
	// Implement indicate the capsule implementation interfaces e.g. `error_p.Error`
	// It CAN also write above capsule as comment in this format:::
	// `//memar:impl memar/process/error/protocol.Error`
	// 
	// `string` MUST include full interface location like above example.
	Implement() []string
}
