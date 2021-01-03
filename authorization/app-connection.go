/* For license and copyright information please see LEGAL file in repository */

package authorization

import (
	etime "../earth-time"
	er "../error"
	"../json"
	"../syllab"
)

// AccessControl store needed data to authorize a request to a platform!
// Use add methods due to arrays must store in sort for easy read and comparison
// Some standards: ABAC, ACL, RBAC, ... features for AccessControl fields,
type AccessControl struct {
	// Authorize Where:
	AllowSocieties []uint32 // Part of GP address
	DenySocieties  []uint32 // Part of GP address
	AllowRouters   []uint32 // Part of GP address
	DenyRouters    []uint32 // Part of GP address

	// Authorize When:
	AllowWeekdays etime.Weekdays // Any day of the week
	DenyWeekdays  etime.Weekdays // Any day of the week
	AllowDayhours etime.Dayhours // Any hour of the day
	DenyDayhours  etime.Dayhours // Any hour of the day

	// Authorize Which:
	AllowServices []uint32 // ["ServiceID", "ServiceID"]
	DenyServices  []uint32 // ["ServiceID", "ServiceID"]
	AllowCRUD     CRUD     // CRUD == Create, Read, Update, Delete
	DenyCRUD      CRUD     // CRUD == Create, Read, Update, Delete

	// Authorize What:
	// Authorize How:
	// Authorize If:
}

// GiveFullAccess set some data to given AccessControl to be full access
func (ac *AccessControl) GiveFullAccess() {
	ac.AllowSocieties = nil
	ac.DenySocieties = nil
	ac.AllowRouters = nil
	ac.DenyRouters = nil

	ac.AllowWeekdays = etime.WeekdaysAll
	ac.DenyWeekdays = etime.WeekdaysNone
	ac.AllowDayhours = etime.DayhoursAll
	ac.DenyDayhours = etime.DayhoursNone

	ac.AllowServices = nil
	ac.DenyServices = nil
	ac.AllowCRUD = CRUDAll
	ac.DenyCRUD = CRUDNone
}

// AuthorizeWhich authorize by ServiceID and CRUD that allow or denied by access control table!
func (ac *AccessControl) AuthorizeWhich(serviceID uint32, crud CRUD) (err *er.Error) {
	if !ac.AllowCRUD.Check(crud) {
		return ErrCrudNotAllow
	}

	if ac.DenyCRUD.Check(crud) {
		return ErrCRUDDenied
	}

	var i int
	var notAuthorize bool

	var ln = len(ac.AllowServices)
	if ln > 0 {
		for i = 0; i < ln; i++ {
			if ac.AllowServices[i] == serviceID {
				goto DS
			} else {
				notAuthorize = true
			}
		}
		if notAuthorize {
			return ErrServiceNotAllow
		}
	}

DS:
	ln = len(ac.DenyServices)
	if ln > 0 {
		for i = 0; i < ln; i++ {
			if ac.DenyServices[i] == serviceID {
				return ErrServiceDenied
			}
		}
	}

	return
}

// AuthorizeWhen --
func (ac *AccessControl) AuthorizeWhen(day etime.Weekdays, hours etime.Dayhours) (err *er.Error) {
	if !ac.AllowWeekdays.Check(day) {
		return ErrDayNotAllow
	}

	if ac.DenyWeekdays.Check(day) {
		return ErrDayDenied
	}

	if !ac.AllowDayhours.Check(hours) {
		return ErrHourNotAllow
	}

	if ac.DenyDayhours.Check(hours) {
		return ErrHourDenied
	}
	return
}

// AuthorizeWhere --
func (ac *AccessControl) AuthorizeWhere(societyID, RouterID uint32) (err *er.Error) {
	var i int
	var notAuthorize bool

	var ln = len(ac.AllowSocieties)
	if ln > 0 {
		for i = 0; i < ln; i++ {
			if ac.AllowSocieties[i] == societyID {
				goto DS
			} else {
				notAuthorize = true
			}
		}
		if notAuthorize {
			return ErrNotAllowSociety
		}
	}

DS:
	ln = len(ac.DenySocieties)
	if ln > 0 {
		for i = 0; i < ln; i++ {
			if ac.DenySocieties[i] == societyID {
				return ErrDeniedSociety
			}
		}
	}

	ln = len(ac.AllowRouters)
	if ln > 0 {
		for i = 0; i < ln; i++ {
			if ac.AllowSocieties[i] == RouterID {
				goto DR
			} else {
				notAuthorize = true
			}
		}
		if notAuthorize {
			return ErrNotAllowRouter
		}
	}

DR:
	ln = len(ac.DenyRouters)
	if ln > 0 {
		for i = 0; i < ln; i++ {
			if ac.DenyRouters[i] == RouterID {
				return ErrDeniedRouter
			}
		}
	}

	return
}

// AuthorizeWhat --
func (ac *AccessControl) AuthorizeWhat() (err *er.Error) {
	// TODO::: can be implement?

	// AllowRecords   [][16]byte // ["RecordUUID", "RecordUUID"]
	// DenyRecords    [][16]byte // ["RecordUUID", "RecordUUID"]
	return
}

// AuthorizeHow --
func (ac *AccessControl) AuthorizeHow() (err *er.Error) {
	// TODO::: can be implement?

	// How []string
	return
}

// AuthorizeIf --
func (ac *AccessControl) AuthorizeIf() (err *er.Error) {
	// TODO::: can be implement?

	// If  []string
	return
}

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
func (ac *AccessControl) JSONDecoder(decoder json.DecoderUnsafeMinifed) (err *er.Error) {
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
	return
}

// JSONEncoder encode given AccessControl to json format.
func (ac *AccessControl) JSONEncoder(encoder json.Encoder) {
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
