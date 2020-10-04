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
	detail      map[lang.Language]Detail
	Chain       *Error
	information interface{}
	JSON        []byte
	Syllab      []byte
}

// Detail store detail about an error
type Detail struct {
	Short string
	Long  string
}

// New returns a new error!
func New() *Error {
	var err = Error{
		detail: make(map[lang.Language]Detail),
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
	var englishName = e.detail[lang.EnglishLanguage].Short
	if englishName == "" {
		log.Fatal("Error must have english name >> ", e)
	}

	e.id = crc32.ChecksumIEEE(*(*[]byte)(unsafe.Pointer(&englishName)))
	e.idAsString = strconv.FormatUint(uint64(e.id), 10)
	e.Syllab = e.syllabEncoder()
	e.JSON = e.jsonEncoder()

	if errPool[e.id] != nil {
		log.Warn("Duplicate Error id exist, Check it now!!!!!!!!!!!!!!!!!")
		log.Warn("Exiting error >> ", errPool[e.id].detail[lang.EnglishLanguage].Short, " New error >> ", englishName)
	}
	errPool[e.id] = e
	return e
}

// SetDetail add short and long text detail to existing error and return it.
func (e *Error) SetDetail(lang lang.Language, short, long string) *Error {
	e.detail[lang] = Detail{
		Short: short,
		Long:  long,
	}

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
		return ErrErrorIsEmpty.detail[lang].Long
	}
	return fmt.Sprintf("Error Code: %v\n Short detail: %v\n Long detail: %v\n Error Additional information: %v\n", e.id, e.detail[lang].Short, e.detail[lang].Long, e.information)
}

// AddInformation add to existing error and return it as new error(pointer)!
func (e *Error) AddInformation(information interface{}) *Error {
	return &Error{
		Chain:       e,
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
