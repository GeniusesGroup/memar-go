/* For license and copyright information please see LEGAL file in repository */

package authorization

import (
	etime "../earth-time"
	er "../error"
	"../json"
	"../syllab"
)

// SyllabDecoder decode syllab to given AccessControl
func (ac *AccessControl) SyllabDecoder(buf []byte, stackIndex uint32) {
	ac.AllowSocieties = syllab.GetUInt32Array(buf, stackIndex)
	ac.DenySocieties = syllab.GetUInt32Array(buf, stackIndex+8)
	ac.AllowRouters = syllab.GetUInt32Array(buf, stackIndex+16)
	ac.DenyRouters = syllab.GetUInt32Array(buf, stackIndex+24)

	ac.AllowWeekdays = etime.Weekdays(syllab.GetUInt8(buf, stackIndex+32))
	ac.DenyWeekdays = etime.Weekdays(syllab.GetUInt8(buf, stackIndex+33))
	ac.AllowDayhours = etime.Dayhours(syllab.GetUInt32(buf, stackIndex+34))
	ac.DenyDayhours = etime.Dayhours(syllab.GetUInt32(buf, stackIndex+38))

	ac.AllowServices = syllab.GetUInt32Array(buf, stackIndex+42)
	ac.DenyServices = syllab.GetUInt32Array(buf, stackIndex+50)
	ac.AllowCRUD = CRUD(syllab.GetUInt8(buf, stackIndex+58))
	ac.DenyCRUD = CRUD(syllab.GetUInt8(buf, stackIndex+59))
	return
}

// SyllabEncoder encode given AccessControl to syllab format
func (ac *AccessControl) SyllabEncoder(buf []byte, stackIndex, heapIndex uint32) (nextHeapAddr uint32) {
	heapIndex = syllab.SetUInt32Array(buf, ac.AllowSocieties, stackIndex, heapIndex)
	heapIndex = syllab.SetUInt32Array(buf, ac.DenySocieties, stackIndex+8, heapIndex)
	heapIndex = syllab.SetUInt32Array(buf, ac.AllowRouters, stackIndex+16, heapIndex)
	heapIndex = syllab.SetUInt32Array(buf, ac.DenyRouters, stackIndex+24, heapIndex)

	syllab.SetUInt8(buf, stackIndex+32, uint8(ac.AllowWeekdays))
	syllab.SetUInt8(buf, stackIndex+33, uint8(ac.DenyWeekdays))
	syllab.SetUInt32(buf, stackIndex+34, uint32(ac.AllowDayhours))
	syllab.SetUInt32(buf, stackIndex+38, uint32(ac.DenyDayhours))

	heapIndex = syllab.SetUInt32Array(buf, ac.AllowServices, stackIndex+42, heapIndex)
	heapIndex = syllab.SetUInt32Array(buf, ac.DenyServices, stackIndex+50, heapIndex)
	syllab.SetUInt8(buf, stackIndex+58, uint8(ac.AllowCRUD))
	syllab.SetUInt8(buf, stackIndex+59, uint8(ac.DenyCRUD))
	return heapIndex
}

// SyllabStackLen return stack length of AccessControl
func (ac *AccessControl) SyllabStackLen() (ln uint32) {
	return 60 // 8+8+8+8+ 1+1+4+4+ 8+8+1+1
}

// SyllabHeapLen return heap length of AccessControl
func (ac *AccessControl) SyllabHeapLen() (ln uint32) {
	ln += uint32(len(ac.AllowSocieties) * 4)
	ln += uint32(len(ac.DenySocieties) * 4)
	ln += uint32(len(ac.AllowRouters) * 4)
	ln += uint32(len(ac.DenyRouters) * 4)
	ln += uint32(len(ac.AllowServices) * 4)
	ln += uint32(len(ac.DenyServices) * 4)
	return
}

// SyllabLen return whole length of AccessControl
func (ac *AccessControl) SyllabLen() (ln int) {
	return int(ac.SyllabStackLen() + ac.SyllabHeapLen())
}

// JSONDecoder decode json to given AccessControl
func (ac *AccessControl) JSONDecoder(buf []byte) (remainBuf []byte, err *er.Error) {
	var decoder = json.DecoderUnsafeMinifed{
		Buf: buf,
	}
	for err == nil {
		var keyName = decoder.DecodeKey()
		switch keyName {
		case "AllowSocieties":
			ac.AllowSocieties, err = decoder.DecodeUInt32SliceAsNumber()
		case "DenySocieties":
			ac.DenySocieties, err = decoder.DecodeUInt32SliceAsNumber()
		case "AllowRouters":
			ac.AllowRouters, err = decoder.DecodeUInt32SliceAsNumber()
		case "DenyRouters":
			ac.DenyRouters, err = decoder.DecodeUInt32SliceAsNumber()

		case "AllowWeekdays":
			var num uint8
			num, err = decoder.DecodeUInt8()
			ac.AllowWeekdays = etime.Weekdays(num)
		case "DenyWeekdays":
			var num uint8
			num, err = decoder.DecodeUInt8()
			ac.DenyWeekdays = etime.Weekdays(num)
		case "AllowDayhours":
			var num uint8
			num, err = decoder.DecodeUInt8()
			ac.AllowDayhours = etime.Dayhours(num)
		case "DenyDayhours":
			var num uint8
			num, err = decoder.DecodeUInt8()
			ac.DenyDayhours = etime.Dayhours(num)

		case "AllowServices":
			ac.AllowServices, err = decoder.DecodeUInt32SliceAsNumber()
		case "DenyServices":
			ac.DenyServices, err = decoder.DecodeUInt32SliceAsNumber()
		case "AllowCRUD":
			var num uint8
			num, err = decoder.DecodeUInt8()
			ac.AllowCRUD = CRUD(num)
		case "DenyCRUD":
			var num uint8
			num, err = decoder.DecodeUInt8()
			ac.DenyCRUD = CRUD(num)

		default:
			err = decoder.NotFoundKey()
		}

		if len(decoder.Buf) < 3 {
			return
		}
	}
	remainBuf = decoder.Buf
	return
}

// JSONEncoder encode given AccessControl to json format.
func (ac *AccessControl) JSONEncoder(buf []byte) (addBuf []byte) {
	var encoder = json.Encoder{
		Buf: buf,
	}

	encoder.EncodeString(`{"AllowSocieties":[`)
	encoder.EncodeUInt32SliceAsNumber(ac.AllowSocieties)

	encoder.EncodeString(`],"DenySocieties":[`)
	encoder.EncodeUInt32SliceAsNumber(ac.DenySocieties)

	encoder.EncodeString(`],"AllowRouters":[`)
	encoder.EncodeUInt32SliceAsNumber(ac.AllowRouters)

	encoder.EncodeString(`],"DenyRouters":[`)
	encoder.EncodeUInt32SliceAsNumber(ac.DenyRouters)

	encoder.EncodeString(`],"AllowWeekdays":`)
	encoder.EncodeUInt8(uint8(ac.AllowWeekdays))

	encoder.EncodeString(`,"DenyWeekdays":`)
	encoder.EncodeUInt8(uint8(ac.DenyWeekdays))

	encoder.EncodeString(`,"AllowDayhours":`)
	encoder.EncodeUInt64(uint64(ac.AllowDayhours))

	encoder.EncodeString(`,"DenyDayhours":`)
	encoder.EncodeUInt64(uint64(ac.DenyDayhours))

	encoder.EncodeString(`,"AllowServices":[`)
	encoder.EncodeUInt32SliceAsNumber(ac.AllowServices)

	encoder.EncodeString(`],"DenyServices":[`)
	encoder.EncodeUInt32SliceAsNumber(ac.DenyServices)

	encoder.EncodeString(`],"AllowCRUD":`)
	encoder.EncodeUInt8(uint8(ac.AllowCRUD))

	encoder.EncodeString(`,"DenyCRUD":`)
	encoder.EncodeUInt8(uint8(ac.DenyCRUD))

	encoder.EncodeByte('}')
	return encoder.Buf
}

// JSONLen return json needed len to encode!
func (ac *AccessControl) JSONLen() (ln int) {
	ln = len(ac.AllowSocieties) * 11
	ln += len(ac.DenySocieties) * 11
	ln += len(ac.AllowRouters) * 11
	ln += len(ac.DenyRouters) * 11
	ln += len(ac.AllowServices) * 11
	ln += len(ac.DenyServices) * 11
	ln += 247
	return
}
