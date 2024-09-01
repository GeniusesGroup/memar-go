/* For license and copyright information please see the LEGAL file in the code repository */

package math_p

// https://docs.microsoft.com/en-us/dotnet/api/system.decimal?view=net-6.0#methods
// https://github.com/shopspring/decimal
type Decimal /*[I Integer, D Signed]*/ interface {
	Integer() int64  // "-123456" of the number "-123456.235689"
	Decimal() uint64 // "235689" of the number "-123456.235689"

	Mul() // multiple
	Add()
	Sub()
	Div()
}
