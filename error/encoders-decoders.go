/* For license and copyright information please see LEGAL file in repository */

package error

import (
	"strconv"

	"../protocol"
)

// FromSyllab decode syllab to given Error
func (e *Error) FromSyllab(buf []byte) {

}

func (e *Error) ToSyllab() (buf []byte) {
	buf = make([]byte, 8)

	return
}

// JSONDecoder decode json to given Error
func (e *Error) FromJSON(buf protocol.Buffer) {

}

func (e *Error) ToJSON() (buf []byte) {
	buf = make([]byte, 0, e.LenAsJSON())

	buf = append(buf, `{"ID":`...)
	buf = append(buf, e.idAsString...)

	buf = append(buf, `,"Detail":{`...)
	if e.detail != nil {
		for key, value := range e.detail {
			buf = append(buf, '"')
			buf = strconv.AppendUint(buf, uint64(key), 10)
			buf = append(buf, `":{"Domain":"`...)
			buf = append(buf, value.domain...)
			buf = append(buf, `","Short":"`...)
			buf = append(buf, value.short...)
			buf = append(buf, `","Long":"`...)
			buf = append(buf, value.long...)
			buf = append(buf, `","UserAction":"`...)
			buf = append(buf, value.userAction...)
			buf = append(buf, `","DevAction":"`...)
			buf = append(buf, value.devAction...)
			buf = append(buf, `"},`...)
		}
		buf = buf[:len(buf)-1] // Remove trailing comma
	}

	buf = append(buf, "}}"...)
	return
}

func (e *Error) LenAsJSON() (ln int) {
	ln = 38 // len(`{"ID":18446744073709551615,"Detail":{}`)
	if len(e.detail) != 0 {
		ln += len(e.detail) * 76 // 76 = 10(len(uint32)) + 66(len('{"Domain":"","Short":"","Long":"","UserAction":"","DevAction":"",}'))
		for _, value := range e.detail {
			ln += len(value.domain)
			ln += len(value.short)
			ln += len(value.long)
			ln += len(value.userAction)
			ln += len(value.devAction)
		}
	}
	return ln
}
