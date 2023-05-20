//go:build lang_eng

/* For license and copyright information please see the LEGAL file in the code repository */

package json

import (
	"libgo/protocol"
)

const domainEnglish = "JSON"

func init() {
	ErrEncodedIncludeNotDefinedKey.SetDetail(protocol.LanguageEnglish, domainEnglish,
		"Encoded Include Not Deffiend Key",
		"Given encoded json string include a key that must not be in the encoded string",
		"",
		"",
		nil)

	ErrEncodedCorrupted.SetDetail(protocol.LanguageEnglish, domainEnglish,
		"Encoded Corrupted",
		"Given encoded json string corrupted and not encode in the way that can decode",
		"",
		"",
		nil)

	ErrEncodedIntegerCorrupted.SetDetail(protocol.LanguageEnglish, domainEnglish,
		"Encoded Integer Corrupted",
		"Given encoded json in Integer part corrupted and not encode in the way that can decode",
		"",
		"",
		nil)

	ErrEncodedStringCorrupted.SetDetail(protocol.LanguageEnglish, domainEnglish,
		"Encoded String Corrupted",
		"Given encoded json in string part corrupted and not encode in the way that can decode",
		"",
		"",
		nil)

	ErrEncodedArrayCorrupted.SetDetail(protocol.LanguageEnglish, domainEnglish,
		"Encoded Array Corrupted",
		"Given encoded json in array part corrupted and not encode in the way that can decode",
		"",
		"",
		nil)

	ErrEncodedSliceCorrupted.SetDetail(protocol.LanguageEnglish, domainEnglish,
		"Encoded Slice Corrupted",
		"Given encoded json in slice part corrupted and not encode in the way that can decode",
		"",
		"",
		nil)
}

func init() {
	MediaType.SetDetail(protocol.LanguageEnglish,
		"JavaScript Object Notation format",
		"",
		"",
		"",
		"",
		[]string{})
}
