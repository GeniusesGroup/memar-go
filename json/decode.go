/* For license and copyright information please see LEGAL file in repository */

package json

import (
	"bytes"
	"encoding/base64"
	"strconv"
	"unsafe"
)

// Decoder store data to decode data by each method!
type Decoder struct {
	Buf     []byte
	UIntNum uint64
	IntNum  int64
	Found   bool
}

// IterationCheckStart call in start of each decode iteration.
func (d *Decoder) IterationCheckStart() {
	var loc = bytes.IndexByte(d.Buf, '"')
	d.Buf = d.Buf[loc+1:] // remove >>	'{"' 	&& 		',"'	due to don't need them
}

// IterationCheckStartMinifed call in start of each decode iteration.
func (d *Decoder) IterationCheckStartMinifed() {
	d.Buf = d.Buf[2:] // remove >>	'{"' 	&& 		',"'	due to don't need them
}

// IterationCheckEnd call in end of each decode iteration.
func (d *Decoder) IterationCheckEnd() error {
	if d.Found {
		d.Found = false
		return nil
	}
	return ErrJSONEncodedStringCorrupted
}

// DecodeSliceAsNumber convert number string slice to []byte. pass buf start from [ and receive from ]
func (d *Decoder) DecodeSliceAsNumber() (slice []byte, err error) {
	var loc int // Coma, Colon, bracket, ... location
	var num uint64
	var end bool

	loc = bytes.IndexByte(d.Buf, ']')
	d.Buf = d.Buf[loc+1:] // indicate next json coma start location

	slice = make([]byte, 0, loc/2) // TODO::: Is loc/2 efficient enough?

	var tmpUnsafe []byte
	for !end {
		loc = bytes.IndexByte(d.Buf, ',')
		if loc < 0 {
			// Reach last item!
			end = true
		}
		tmpUnsafe = d.Buf[:loc]
		num, err = strconv.ParseUint(*(*string)(unsafe.Pointer(&tmpUnsafe)), 10, 8)
		if err != nil {
			return
		}
		slice = append(slice, byte(num))
		d.Buf = d.Buf[loc+1:]
	}

	d.Found = true
	return
}

// DecodeSliceAsBase64 convert base64 string to []byte
func (d *Decoder) DecodeSliceAsBase64() (slice []byte, err error) {
	var loc int // Coma, Colon, bracket, ... location

	loc = bytes.IndexByte(d.Buf, '"')

	slice = make([]byte, base64.StdEncoding.DecodedLen(len(d.Buf[:loc])))
	_, err = base64.StdEncoding.Decode(slice, d.Buf[:loc])
	if err != nil {
		return
	}

	d.Buf = d.Buf[loc+1:]

	d.Found = true
	return
}

// Decode16ByteArrayAsNumber use to convert string to [16]byte. pass d.Buf start from [ and receive from ]
func (d *Decoder) Decode16ByteArrayAsNumber() (array [16]byte, err error) {
	var loc int // Coma, Colon, bracket, ... location
	var num uint64

	var tmpUnsafe []byte
	for i := 0; i < 15; i++ {
		loc = bytes.IndexByte(d.Buf, ',')
		tmpUnsafe = d.Buf[:loc]
		num, err = strconv.ParseUint(*(*string)(unsafe.Pointer(&tmpUnsafe)), 10, 8)
		if err != nil {
			return
		}
		array[i] = byte(num)
		d.Buf = d.Buf[loc+1:]
	}
	loc = bytes.IndexByte(d.Buf, ']')
	tmpUnsafe = d.Buf[:loc]
	num, err = strconv.ParseUint(*(*string)(unsafe.Pointer(&tmpUnsafe)), 10, 8)
	if err != nil {
		return
	}
	array[15] = byte(num)

	d.Buf = d.Buf[loc+1:] // indicate next json coma start location
	d.Found = true
	return
}

// DecodeUIntAsNumber convert number string to number. pass d.Buf start from : and receive from ,
func (d *Decoder) DecodeUIntAsNumber(bitSize int) (err error) {
	var loc int // Coma, Colon, bracket, ... location

	loc = bytes.IndexByte(d.Buf, ',')
	if loc < 0 {
		// Reach last item of d.Buf!
		loc = len(d.Buf) - 1
	}

	var tmpUnsafe = d.Buf[:loc]
	d.UIntNum, err = strconv.ParseUint(*(*string)(unsafe.Pointer(&tmpUnsafe)), 10, bitSize)
	if err != nil {
		return
	}

	d.Buf = d.Buf[loc:] // indicate next json coma start location
	d.Found = true
	return
}
