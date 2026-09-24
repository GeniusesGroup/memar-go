/* For license and copyright information please see the LEGAL file in the code repository */

package face

type Value[C, I any] struct {
	concrete C
	abstract I
}

func (v *Value[C, I]) Init(concrete C, abstract I) {
	v.concrete = concrete
	v.abstract = abstract
}

func (v *Value[C, I]) Interface() I { return v.abstract }

// To call s.String(), the Go compiler generates code that does the equivalent of the C expression
// s.tab->fun[0](s.data): it calls the appropriate function pointer from the itable,
