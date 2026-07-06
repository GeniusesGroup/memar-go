//go:build lang_eng

/* For license and copyright information please see the LEGAL file in the code repository */

package json_errs

const domainEnglish = "JSON"

func init() {
	EncodedIncludeNotDefinedKey.SetDetail(lang_p.LanguageEnglish, domainEnglish,
		"Encoded Include Not Deffiend Key",
		"Given encoded json string include a key that must not be in the encoded string",
		"",
		"",
		nil)

	EncodedCorrupted.SetDetail(lang_p.LanguageEnglish, domainEnglish,
		"Encoded Corrupted",
		"Given encoded json string corrupted and not encode in the way that can decode",
		"",
		"",
		nil)

	EncodedIntegerCorrupted.SetDetail(lang_p.LanguageEnglish, domainEnglish,
		"Encoded Integer Corrupted",
		"Given encoded json in Integer part corrupted and not encode in the way that can decode",
		"",
		"",
		nil)

	EncodedStringCorrupted.SetDetail(lang_p.LanguageEnglish, domainEnglish,
		"Encoded String Corrupted",
		"Given encoded json in string part corrupted and not encode in the way that can decode",
		"",
		"",
		nil)

	EncodedArrayCorrupted.SetDetail(lang_p.LanguageEnglish, domainEnglish,
		"Encoded Array Corrupted",
		"Given encoded json in array part corrupted and not encode in the way that can decode",
		"",
		"",
		nil)

	EncodedSliceCorrupted.SetDetail(lang_p.LanguageEnglish, domainEnglish,
		"Encoded Slice Corrupted",
		"Given encoded json in slice part corrupted and not encode in the way that can decode",
		"",
		"",
		nil)
}

func init() {
	MediaType.SetDetail(lang_p.LanguageEnglish,
		"JavaScript Object Notation format",
		"",
		"",
		"",
		"",
		[]string{})
}
