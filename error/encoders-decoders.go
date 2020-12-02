/* For license and copyright information please see LEGAL file in repository */

package error

import (
	"strconv"
)

// SyllabDecoder decode syllab to given Error
func (e *Error) SyllabDecoder(buf []byte) {

}
func (e *Error) syllabEncoder() (buf []byte) {
	buf = make([]byte, 8)

	buf[4] = byte(e.id)
	buf[5] = byte(e.id >> 8)
	buf[6] = byte(e.id >> 16)
	buf[7] = byte(e.id >> 24)

	return
}

// JSONDecoder decode json to given Error
func (e *Error) JSONDecoder(buf []byte) {

}

func (e *Error) jsonEncoder() (buf []byte) {
	buf = make([]byte, 0, e.jsonLen())

	buf = append(buf, `{"ID":`...)
	buf = append(buf, e.idAsString...)

	buf = append(buf, `,"Detail":{`...)
	if e.detail != nil {
		for key, value := range e.detail {
			buf = append(buf, '"')
			buf = strconv.AppendUint(buf, uint64(key), 10)
			buf = append(buf, `":{"Domain":"`...)
			buf = append(buf, value.Domain...)
			buf = append(buf, `","Short":"`...)
			buf = append(buf, value.Short...)
			buf = append(buf, `","Long":"`...)
			buf = append(buf, value.Long...)
			buf = append(buf, `"},`...)
		}
		buf = buf[:len(buf)-1] // Remove trailing comma
	}

	buf = append(buf, "}}"...)
	return
}

func (e *Error) jsonLen() (ln int) {
	ln = 28 // len(`{"ID":0000000000,"Detail":{}`)
	if e.detail != nil {
		ln += len(e.detail) * 45 // 45 = 10(len(uint32)) + 35(len('{"Domain":"","Short":"","Long":"",}'))
		for _, value := range e.detail {
			ln += len(value.Domain)
			ln += len(value.Short)
			ln += len(value.Long)
		}
	}
	return ln
}
