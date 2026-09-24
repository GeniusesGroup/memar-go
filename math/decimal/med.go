/* For license and copyright information please see the LEGAL file in the code repository */

package dec

import (
	error_p "memar/process/error/protocol"
)

// 1500.10 * 9.99% = [(1500*9/100) + (1500*99/10000)] + [(10*9/100) + (10*99/10000)]
// = 135 + 14.85 = 149.85
// 150010*999/10000 = 1498510
// DEC(2,2)
type Med struct {
	integer int64  //9
	decimal uint64 //99
}

//memar:impl memar/computer/capsule/protocol.LifeCycle
func (d *Med) Init(integer int64, decimal uint64) (err error_p.Error) {
	d.integer = integer
	d.decimal = decimal
	return
}
func (d *Med) Reinit(integer int64, decimal uint64) (err error_p.Error) {
	d.integer = integer
	d.decimal = decimal
	return
}
func (d *Med) Deinit() (err error_p.Error) {
	return
}

func (d Med) Integer() int64  { return d.integer }
func (d Med) Decimal() uint64 { return d.decimal }
