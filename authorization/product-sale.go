/* For license and copyright information please see LEGAL file in repository */

package authorization

import (
	etime "../earth-time"
	er "../error"
	"../json"
	"../syllab"
)

// Product store needed data to authorize product as auction or any other order type
type Product struct {
	AllowUserID   [32]byte `index-hash:"ID" json:",string"` // If 0 means this sale is not just for specific UserID! can be any OrgID e.g. DistributionCenterID
	AllowUserType UserType
	AllowWeekdays etime.Weekdays
	AllowDayhours etime.Dayhours
	GroupID       [32]byte `index-hash:"ID" json:",string"` // it can be 0 and means sale is global!
	MinNumBuy     uint64   // Minimum number to buy in this auction to use for sale-off(Discount)
	StockNumber   uint64   // 0 for unlimited until related product exist to sell!
	LiveUntil     etime.Time
}

// SyllabDecoder decode syllab to given Product
func (p *Product) SyllabDecoder(buf []byte, stackIndex uint32) (err *er.Error) {
	// var add, ln uint32
	// var tempSlice []byte

	if uint32(len(buf)) < p.SyllabStackLen() {
		err = syllab.ErrSyllabDecodeSmallSlice
		return
	}

	copy(p.AllowUserID[:], buf[0:])
	p.AllowUserType = UserType(syllab.GetUInt8(buf, 32))
	p.AllowWeekdays = etime.Weekdays(syllab.GetUInt8(buf, 33))
	p.AllowDayhours = etime.Dayhours(syllab.GetUInt32(buf, 34))
	copy(p.GroupID[:], buf[38:])
	p.MinNumBuy = syllab.GetUInt64(buf, 70)
	p.StockNumber = syllab.GetUInt64(buf, 78)
	p.LiveUntil = etime.Time(syllab.GetInt64(buf, 86))
	return
}

// SyllabEncoder encode given Product to syllab format
func (p *Product) SyllabEncoder(buf []byte, stackIndex, heapIndex uint32) (nextHeapAddr uint32) {
	copy(buf[0:], p.AllowUserID[:])
	syllab.SetUInt8(buf, 32, uint8(p.AllowUserType))
	syllab.SetUInt8(buf, 33, uint8(p.AllowWeekdays))
	syllab.SetUInt32(buf, 34, uint32(p.AllowDayhours))
	copy(buf[38:], p.GroupID[:])
	syllab.SetUInt64(buf, 70, p.MinNumBuy)
	syllab.SetUInt64(buf, 78, p.StockNumber)
	syllab.SetInt64(buf, 86, int64(p.LiveUntil))
	return
}

// SyllabStackLen return stack length of Product
func (p *Product) SyllabStackLen() (ln uint32) {
	return 94
}

// SyllabHeapLen return heap length of Product
func (p *Product) SyllabHeapLen() (ln uint32) {
	return
}

// SyllabLen return whole length of Product
func (p *Product) SyllabLen() (ln int) {
	return int(p.SyllabStackLen() + p.SyllabHeapLen())
}

// JSONDecoder decode json to given Product
func (p *Product) JSONDecoder(decoder json.DecoderUnsafeMinifed) (err *er.Error) {
	for err == nil {
		var keyName = decoder.DecodeKey()
		switch keyName {
		case "AllowUserID":
			err = decoder.DecodeByteArrayAsBase64(p.AllowUserID[:])
		case "AllowUserType":
			var num uint8
			num, err = decoder.DecodeUInt8()
			p.AllowUserType = UserType(num)
		case "AllowWeekdays":
			var num uint8
			num, err = decoder.DecodeUInt8()
			p.AllowWeekdays = etime.Weekdays(num)
		case "AllowDayhours":
			var num uint32
			num, err = decoder.DecodeUInt32()
			p.AllowDayhours = etime.Dayhours(num)
		case "GroupID":
			err = decoder.DecodeByteArrayAsBase64(p.GroupID[:])
		case "MinNumBuy":
			p.MinNumBuy, err = decoder.DecodeUInt64()
		case "StockNumber":
			p.StockNumber, err = decoder.DecodeUInt64()
		case "LiveUntil":
			var num int64
			num, err = decoder.DecodeInt64()
			p.LiveUntil = etime.Time(num)
		default:
			err = decoder.NotFoundKeyStrict()
		}

		if len(decoder.Buf) < 3 {
			return
		}
	}
	return
}

// JSONEncoder encode given Product to json format.
func (p *Product) JSONEncoder(encoder json.Encoder) {
	encoder.EncodeString(`{"AllowUserID":"`)
	encoder.EncodeByteSliceAsBase64(p.AllowUserID[:])

	encoder.EncodeString(`","AllowUserType":`)
	encoder.EncodeUInt8(uint8(p.AllowUserType))

	encoder.EncodeString(`,"AllowWeekdays":`)
	encoder.EncodeUInt8(uint8(p.AllowWeekdays))

	encoder.EncodeString(`,"AllowDayhours":`)
	encoder.EncodeUInt32(uint32(p.AllowDayhours))

	encoder.EncodeString(`,"GroupID":"`)
	encoder.EncodeByteSliceAsBase64(p.GroupID[:])

	encoder.EncodeString(`","MinNumBuy":`)
	encoder.EncodeUInt64(p.MinNumBuy)

	encoder.EncodeString(`,"StockNumber":`)
	encoder.EncodeUInt64(p.StockNumber)

	encoder.EncodeString(`,"LiveUntil":`)
	encoder.EncodeInt64(int64(p.LiveUntil))

	encoder.EncodeByte('}')
}

// JSONLen return json needed len to encode!
func (p *Product) JSONLen() (ln int) {
	ln = 455
	return
}
