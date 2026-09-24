/* For license and copyright information please see the LEGAL file in the code repository */

package integer

// ElementIndex is related to `NumberOfElement`
// ElementIndex can refer to any location of memory blocks in byte or 8bit number.
type ElementIndex int

func (ei *ElementIndex) Increment() { *ei++ }
func (ei *ElementIndex) Decrement() { *ei-- }
