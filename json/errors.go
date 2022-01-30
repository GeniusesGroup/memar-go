/* For license and copyright information please see LEGAL file in repository */

package json

import (
	er "../error"
	"../protocol"
)

const domainEnglish = "JSON"
const domainPersian = "جیسون"

// Declare Errors Details
// TODO::: use json.ietf.org or ??
var (
	ErrEncodedIncludeNotDeffiendKey = er.New("urn:giti:json.ecma-international.org:error:encoded-include-not-deffiend-key").SetDetail(protocol.LanguageEnglish, domainEnglish, "Encoded Include Not Deffiend Key",
		"Given encoded json string include a key that must not be in the encoded string",
		"",
		"").Save()

	ErrEncodedCorrupted = er.New("urn:giti:json.ecma-international.org:error:encoded-corrupted").SetDetail(protocol.LanguageEnglish, domainEnglish, "Encoded Corrupted",
		"Given encoded json string corruputed and not encode in the way that can decode",
		"",
		"").Save()

	ErrEncodedIntegerCorrupted = er.New("urn:giti:json.ecma-international.org:error:encoded-integer-corrupted").SetDetail(protocol.LanguageEnglish, domainEnglish, "Encoded Integer Corrupted",
		"Given encoded json in Integer part corruputed and not encode in the way that can decode",
		"",
		"").Save()

	ErrEncodedStringCorrupted = er.New("urn:giti:json.ecma-international.org:error:encoded-string-corrupted").SetDetail(protocol.LanguageEnglish, domainEnglish, "Encoded String Corrupted",
		"Given encoded json in string part corruputed and not encode in the way that can decode",
		"",
		"").Save()

	ErrEncodedArrayCorrupted = er.New("urn:giti:json.ecma-international.org:error:encoded-array-corrupted").SetDetail(protocol.LanguageEnglish, domainEnglish, "Encoded Array Corrupted",
		"Given encoded json in array part corruputed and not encode in the way that can decode",
		"",
		"").Save()

	ErrEncodedSliceCorrupted = er.New("urn:giti:json.ecma-international.org:error:encoded-slice-corrupted").SetDetail(protocol.LanguageEnglish, domainEnglish, "Encoded Slice Corrupted",
		"Given encoded json in slice part corruputed and not encode in the way that can decode",
		"",
		"").Save()
)
