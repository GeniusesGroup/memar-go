/* For license and copyright information please see LEGAL file in repository */

package authorization

// SyllabDecoder decode syllab to given AccessControl
func (ac *AccessControl) SyllabDecoder(buf []byte, stackIndex int) {
	return
}

// SyllabEncoder encode given AccessControl to syllab format
func (ac *AccessControl) SyllabEncoder(buf []byte, stackIndex, heapIndex uint32) (hi uint32) {

	return heapIndex
}

// SyllabStackLen return stack length of AccessControl
func (ac *AccessControl) SyllabStackLen() (ln uint32) {
	return 68 // 8+8+8+8+ 1+1+8+8+ 8+8+1+1
}

// SyllabHeapLen return heap length of AccessControl
func (ac *AccessControl) SyllabHeapLen() (ln uint32) {
	ln += uint32(len(ac.AllowSocieties) * 4)
	ln += uint32(len(ac.DenySocieties) * 4)
	ln += uint32(len(ac.AllowRouters) * 4)
	ln += uint32(len(ac.DenyRouters) * 4)
	ln += uint32(len(ac.AllowTime) * 2)
	ln += uint32(len(ac.DenyTime) * 2)
	ln += uint32(len(ac.AllowServices) * 4)
	ln += uint32(len(ac.DenyServices) * 4)
	return
}

// SyllabLen return whole length of AccessControl
func (ac *AccessControl) SyllabLen() (ln int) {
	return int(ac.SyllabStackLen() + ac.SyllabHeapLen())
}
