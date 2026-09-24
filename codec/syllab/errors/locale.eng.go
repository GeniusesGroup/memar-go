//go:build lang_eng

/* For license and copyright information please see the LEGAL file in the code repository */

package syllab_errs

const domainEnglish = "Syllab"

func init() {
	NeededTypeNotExist.SetDetail(lang_p.LanguageEnglish, domainEnglish,
		"Needed Type Not Exist",
		"Custom struct type exist in upper struct type that generator can't access it to know its fields",
		"",
		"",
		nil)
	FieldType.SetDetail(lang_p.LanguageEnglish, domainEnglish, "Field Type",
		"Requested type may include function, interface, int, uint, ... field type that can't encode||decode",
		"",
		"",
		nil)
	ArrayLen.SetDetail(lang_p.LanguageEnglish, domainEnglish, "Array Len",
		"Length of array larger than 32 bit space that syllab can encode||decode",
		"",
		"",
		nil)
	ShortArrayDecode.SetDetail(lang_p.LanguageEnglish, domainEnglish, "Short Array Decode",
		"Given array shorter than expected to decode data from it",
		"",
		"Problem can in Dynamic_Array that pass to encoder",
		nil)
	DecodeHeapOverFlow.SetDetail(lang_p.LanguageEnglish, domainEnglish, "Decode Heap OverFlow",
		"Encoded syllab want to access to out of array.",
		"",
		"",
		nil)
}

func init() {
	MediaType.SetDetail(lang_p.LanguageEnglish, "Syllab codec protocol", "", "", "", "", []string{})
}
