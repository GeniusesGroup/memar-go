/* For license and copyright information please see LEGAL file in repository */

package error

import (
	"fmt"
	"hash/crc32"
	"strconv"

	lang "../language"
	"../log"
)

var errPool = map[uint32]*Error{}

// ERRPoolSlice store to access all errors in order
var ERRPoolSlice = []*Error{}

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
	Domain string
	Short  string
	Long   string
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
	return
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
	var englishDetail = e.detail[lang.EnglishLanguage]
	if englishDetail.Short == "" {
		log.Fatal("Error must have english name >> ", *e)
	}
	if englishDetail.Domain == "" {
		log.Fatal("Error must have Domain >> ", *e)
	}

	var buf = make([]byte, len(englishDetail.Domain)+len(englishDetail.Short))
	copy(buf, englishDetail.Domain)
	copy(buf[len(englishDetail.Domain):], englishDetail.Short)

	e.id = crc32.ChecksumIEEE(buf)
	e.idAsString = strconv.FormatUint(uint64(e.id), 10)
	e.Syllab = e.syllabEncoder()
	e.JSON = e.jsonEncoder()

	if errPool[e.id] != nil {
		log.Warn("Duplicate Error id exist, Check it now!!!!!!!!!!!!!!!!!")
		log.Warn("Exiting error >> ", *errPool[e.id], " New error >> ", englishDetail)
	}

	errPool[e.id] = e
	ERRPoolSlice = append(ERRPoolSlice, e)
	return e
}

// SetDetail add short and long text detail to existing error and return it.
func (e *Error) SetDetail(lang lang.Language, domain, short, long string) *Error {
	e.detail[lang] = Detail{
		Domain: domain,
		Short:  short,
		Long:   long,
	}

	return e
}

// GetDetail return short and long text detail to existing error and return it.
func (e *Error) GetDetail(lang lang.Language) (domain, short, long, idAsString string) {
	if e == nil {
		return
	}
	var detail = e.detail[lang]
	return detail.Domain, detail.Short, detail.Long, e.idAsString
}

// ID return id of error if err exist.
func (e *Error) ID() uint32 {
	// if e == nil {
	// 	return 0
	// }
	return e.id
}

// IDasString return id of error as string if err exist.
func (e *Error) IDasString() string {
	// if e == nil {
	// 	return "0"
	// }
	return e.idAsString
}

// Error return full details of error in text.
func (e *Error) Error() string {
	// if e == nil {
	// 	return ErrErrorIsEmpty.detail[log.Language].Long
	// }
	return fmt.Sprintf("Error Code: %v\n Short detail: %v\n Long detail: %v\n Error Additional information: %v\n", e.id, e.detail[log.Language].Short, e.detail[log.Language].Long, e.information)
}

// AddInformation add to existing error and return it as new error(pointer)!
func (e *Error) AddInformation(information interface{}) *Error {
	return &Error{
		Chain:       e,
		information: information,
	}
}

// Equal compare two Error.
func (e *Error) Equal(err *Error) bool {
	if e == nil && err == nil {
		return true
	}
	if e != nil && err != nil && e.id == err.id {
		return true
	}
	return false
}

// IsEqual compare two error.
func (e *Error) IsEqual(err error) bool {
	var exErr = err.(*Error)
	if exErr != nil && e.id == exErr.id {
		return true
	}
	return false
}
