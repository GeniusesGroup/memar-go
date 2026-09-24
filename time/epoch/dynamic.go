/* For license and copyright information please see the LEGAL file in the code repository */

package epoch

type Field_DynamicEpoch interface {
	DynamicEpoch() Dynamic
}

type Dynamic uint8

const (
	DynamicEpoch_Unset Dynamic = iota
	DynamicEpoch_IssueDate

	// question, ...
	DynamicEpoch_StartPoint
	DynamicEpoch_LastUse

	// product modules
	DynamicEpoch_FirstUse
	AddToInvoice
	DynamicEpoch_PurchaseTime

	// task modules
	DynamicEpoch_LastInterval
	DynamicEpoch_Continuously

	// Order
	DynamicEpoch_LastSuggest
	DynamicEpoch_LastDeal
)
