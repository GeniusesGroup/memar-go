/* For license and copyright information please see LEGAL file in repository */

package error

import (
	"fmt"
	"hash/crc32"
	"strconv"
	"unsafe"

	lang "../language"
	"../log"
)

var errPool = map[uint32]*Error{}

// Error is a extended implementation of error.
// Never change english name due to it adds unnecessary complicated troubleshooting errors on SDK!
type Error struct {
	id          uint32
	idAsString  string
	name        map[lang.Language]string
	description map[lang.Language]string
	information interface{}
	JSON        []byte
	Syllab      []byte
}

// New returns a new error!
func New() *Error {
	var err = Error{
		name:        make(map[lang.Language]string),
		description: make(map[lang.Language]string),
	}
	return &err
}

// GetErrByCode returns desire error if exist or nil!
func GetErrByCode(id uint32) (err *Error) {
	err = errPool[id]
	if err == nil {
		return ErrErrorNotFound
	}
	return err
}

// GetCode return id of error if err id exist.
func GetCode(err error) uint32 {
	if err == nil {
		return 0
	}
	var exErr *Error
	exErr = err.(*Error)
	if exErr != nil {
		return exErr.id
	}
	// if error not nil but not Error, pass biggest number!
	return 4294967295
}

// Save finalize needed logic on given error and save to errPool.
func (e *Error) Save() *Error {
	var englishName = e.name[lang.EnglishLanguage]
	if englishName == "" {
		log.Fatal("Error must have english name >> ", e)
	}

	e.id = crc32.ChecksumIEEE(*(*[]byte)(unsafe.Pointer(&englishName)))
	e.idAsString = strconv.FormatUint(uint64(e.id), 10)
	e.Syllab = e.syllabEncoder()
	e.JSON = e.jsonEncoder()

	if errPool[e.id] != nil {
		log.Warn("Duplicate Error id exist")
		log.Fatal("Exiting error >> ", errPool[e.id].name[lang.EnglishLanguage], " New error >> ", englishName)
	}
	errPool[e.id] = e
	return e
}

// SetName add name to existing error and return it.
func (e *Error) SetName(lang lang.Language, name string) *Error {
	e.name[lang] = name
	return e
}

// SetDescription add description text to existing error and return it.
func (e *Error) SetDescription(lang lang.Language, text string) *Error {
	e.description[lang] = text
	return e
}

// Error returns id of error.
func (e *Error) Error() string {
	if e == nil {
		return "0"
	}
	return e.idAsString
}

// Text return full details of error in text.
func (e *Error) Text(lang lang.Language) string {
	if e == nil {
		return "Error is empty"
	}
	return fmt.Sprintf("Error Code: %v\n Error Name: %v\n Error text: %v\n Error Additional information: %v\n", e.id, e.name[lang], e.description[lang], e.information)
}

// AddInformation add to existing error and return it as new error(pointer)!
func (e *Error) AddInformation(information interface{}) error {
	return &Error{
		name:        e.name,
		description: e.description,
		id:          e.id,
		information: information,
	}
}

// IsEqual compare two error.
func (e *Error) IsEqual(err error) bool {
	var exErr = err.(*Error)
	if exErr != nil && e.id == exErr.id {
		return true
	}

	return false
}

func (e *Error) syllabEncoder() (buf []byte) {
	buf = make([]byte, 8)

	buf[4] = byte(e.id)
	buf[5] = byte(e.id >> 8)
	buf[6] = byte(e.id >> 16)
	buf[7] = byte(e.id >> 24)

	return
}

func (e *Error) jsonEncoder() (buf []byte) {
	buf = make([]byte, 0, e.jsonLen())

	buf = append(buf, `{"ID":`...)
	buf = strconv.AppendUint(buf, uint64(e.id), 10)

	buf = append(buf, `,"Name":{`...)
	if e.name != nil {
		for key, value := range e.name {
			buf = append(buf, '"')
			buf = strconv.AppendUint(buf, uint64(key), 10)
			buf = append(buf, `":"`...)
			buf = append(buf, value...)
			buf = append(buf, `",`...)
		}
		buf = buf[:len(buf)-1] // Remove trailing comma
	}

	buf = append(buf, `},"Description":{`...)
	if e.description != nil {
		for key, value := range e.description {
			buf = append(buf, '"')
			buf = strconv.AppendUint(buf, uint64(key), 10)
			buf = append(buf, `":"`...)
			buf = append(buf, value...)
			buf = append(buf, `",`...)
		}
		buf = buf[:len(buf)-1] // Remove trailing comma
	}

	buf = append(buf, "}}"...)
	return
}

func (e *Error) jsonLen() (ln int) {
	ln = 33 // len(`{"ID":,"Name":{},"Description":{}`)
	if e.name != nil {
		ln += len(e.name) * 15 // 15 = 5(len('"":""')) + 10(len(uint32))
		for _, value := range e.name {
			ln += len(value)
		}
	}
	if e.description != nil {
		ln += len(e.name) * 15 // 15 = 5(len('"":""')) + 10(len(uint32))
		for _, value := range e.description {
			ln += len(value)
		}
	}
	return ln
}

/*
We choose first style for error declaration!

	var Err--- = errorr.New().
		SetName(lang.EnglishLanguage, "nnnn").
		SetDescription(lang.EnglishLanguage, "dddd").Save()

	var Err--- = errorr.Error{
		Name: map[lang.Language]string{
			lang.EnglishLanguage: "nnnn",
		},
		Description: map[lang.Language]string{
			lang.EnglishLanguage: "nnnn",
		},
	}

*/
