/* For license and copyright information please see LEGAL file in repository */

package error

import (
	"../log"
	"../protocol"
)

// Errors store
type Errors struct {
	poolByID  map[uint64]protocol.Error
	poolByURN map[string]protocol.Error
	jsSDK     map[protocol.LanguageID][]byte
}

func (e *Errors) Init() {
	e.poolByID = make(map[uint64]protocol.Error, 512)
	e.poolByURN = make(map[string]protocol.Error, 512)
	e.jsSDK = map[protocol.LanguageID][]byte{}
}

func (e *Errors) RegisterError(err protocol.Error) {
	var errID = err.URN().ID()
	var exitingError = e.poolByID[errID]
	if exitingError != nil {
		log.Warn("Duplicate Error id exist, Check it now for bad urn set or collision occurred!")
		log.Warn("Exiting error >> ", exitingError.URN(), " New error >> ", err.URN())
		return
	}

	e.poolByID[errID] = err
	e.poolByURN[err.URN().URI()] = err
	e.updateJsSDK(err)
}

func (e *Errors) UnRegisterError(err protocol.Error) {
	delete(e.poolByID, err.URN().ID())
	delete(e.poolByURN, err.URN().URI())
}

// GetErrorByID returns desire error if exist or ErrNotFound!
func (e *Errors) GetErrorByID(id uint64) (err protocol.Error) {
	if id == 0 {
		return
	}
	var ok bool
	err, ok = e.poolByID[id]
	if !ok {
		err = ErrNotFound
	}
	return
}

// GetErrorByID returns desire error if exist or ErrNotFound!
func (e *Errors) GetErrorByURN(urn string) (err protocol.Error) {
	var ok bool
	err, ok = e.poolByURN[urn]
	if !ok {
		err = ErrNotFound
	}
	return
}

func (e *Errors) updateJsSDK(err protocol.Error) {
	for _, detail := range err.Details() {
		var lang = detail.Language()
		e.jsSDK[lang] = append(e.jsSDK[lang], "GitiError.New(\""+err.IDasString()+"\",\""+err.URN().URI()+"\").SetDetail(\""+detail.Domain()+"\",\""+detail.Short()+"\",\""+detail.Long()+"\",\""+detail.UserAction()+"\",\""+detail.DevAction()+"\")\n"...)
	}
}

// ErrorsSDK can return nil slice if request not supported!
func (e *Errors) ErrorsSDK(humanLanguage protocol.LanguageID, machineLanguage protocol.MediaType) (sdk []byte, err protocol.Error) {
	switch machineLanguage.Extension() {
	case "js":
		sdk = e.jsSDK[humanLanguage]
	}
	if sdk == nil {
		err = ErrSDKNotFound
	}
	return
}
