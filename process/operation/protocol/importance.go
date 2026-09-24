/* For license and copyright information please see the LEGAL file in the code repository */

package operation_p

type Importance interface {
	// We need both of below items to queue in more specific manner.
	Field_Priority
	Field_Weight
}
