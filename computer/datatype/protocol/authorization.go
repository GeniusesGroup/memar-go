/* For license and copyright information please see the LEGAL file in the code repository */

package datatype_p

type Authorization interface {
	// LimitToUseBy or ImportableBy
	LimitToUseBy() DataTypes
}
