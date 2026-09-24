/* For license and copyright information please see the LEGAL file in the code repository */

package decimal_p

type Field_Quantity interface {
	Quantity() Quantity
}

// Quantity or Amount
type Quantity Decimal
