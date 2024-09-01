/* For license and copyright information please see the LEGAL file in the code repository */

package operation_p

type Importance interface {
	// TODO::: need both of below items??!!
	Priority() Priority // Use to queue requests by its priority
	Weight() Weight     // Use to queue requests by its weights in the same priority
}
