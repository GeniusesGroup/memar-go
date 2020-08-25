/* For license and copyright information please see LEGAL file in repository */

package json

import (
	"../assets"
)

// https://github.com/pquerna/ffjson
// https://github.com/mailru/easyjson

/*
Before pass file to safe||unsafe function, dev must add needed methods to desire type by below template!
Otherwise panic may occur due to improve performance we don't check some bad situation!!

func ({{DesireName}} *{{DesireType}}) jsonDecoder(buf []byte) (err error) {
	return
}

func ({{DesireName}} *{{DesireType}}) jsonEncoder(offset int) (buf []byte) {
	return
}
*/

// CompleteMethodsSafe use to update given go files and complete json encoder&&decoder to any struct type in it!
// It will overwrite given file methods! If you need it clone it before pass it here!
func CompleteMethodsSafe(file *assets.File) (err error) {

	return
}
