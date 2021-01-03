/* For license and copyright information please see LEGAL file in repository */

package matn

import (
	er "../error"
)

// IndexTextGetReq is request structure of IndexTextGet()
type IndexTextGetReq struct {
	Terms      []string
	PageNumber uint64
}

// IndexTextGetRes is response structure of IndexTextGet()
type IndexTextGetRes struct {
	Tokens [10]PhraseToken
}

// IndexTextGet return index data of given terms if any exist
func IndexTextGet(req *IndexTextGetReq) (res *IndexTextGetRes, err *er.Error) {

	res = &IndexTextGetRes{}

	return
}

// PhraseTokenization uses the delimiters categorized under [Dash, Hyphen, Pattern_Syntax, Quotation_Mark, Terminal_Punctuation, White_Space]
// drop language-specific stop words (e.g. in English, the, an, a, and, etc.)
func PhraseTokenization(text string) (tokenized []string) {
	return
}
