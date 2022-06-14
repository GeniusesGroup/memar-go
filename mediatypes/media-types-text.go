/* For license and copyright information please see LEGAL file in repository */

package mediatypes

import (
	"../mediatype"
	"../protocol"
)

var (
	ASC mediatype.MediaType

	CSS mediatype.MediaType
	CSV mediatype.MediaType

	ICS mediatype.MediaType

	TEXT mediatype.MediaType
	TXT  mediatype.MediaType

	HTM  mediatype.MediaType
	HTML mediatype.MediaType

	TSV mediatype.MediaType

	XSL mediatype.MediaType
	XSD mediatype.MediaType
)

func init() {
	ASC.Init("text/plain")
	ASC.SetFileExtension("asc")
	ASC.SetDetail(protocol.LanguageEnglish, "", "", "", "", "", []string{})

	CSS.Init("text/css")
	CSS.SetFileExtension("css")
	CSS.SetDetail(protocol.LanguageEnglish, "Cascading Style Sheets", "", "", "", "", []string{})

	CSV.Init("text/csv")
	CSV.SetFileExtension("csv")
	CSV.SetDetail(protocol.LanguageEnglish, "Comma-separated values", "", "", "", "", []string{})

	ICS.Init("text/calendar")
	ICS.SetFileExtension("ics")
	ICS.SetDetail(protocol.LanguageEnglish, "iCalendar format", "", "", "", "", []string{})

	TEXT.Init("text/plain")
	TEXT.SetFileExtension("text")
	TEXT.SetDetail(protocol.LanguageEnglish, "", "", "", "", "", []string{})

	TXT.Init("text/plain")
	TXT.SetFileExtension("txt")
	TXT.SetDetail(protocol.LanguageEnglish, "", "", "", "", "", []string{})

	HTM.Init("text/html")
	HTM.SetFileExtension("htm")
	HTM.SetDetail(protocol.LanguageEnglish, "HyperText Markup Language", "", "", "", "", []string{})

	HTML.Init("text/html")
	HTML.SetFileExtension("html")
	HTML.SetDetail(protocol.LanguageEnglish, "HyperText Markup Language", "", "", "", "", []string{})

	TSV.Init("text/tab-separated-values")
	TSV.SetFileExtension("tsv")
	TSV.SetDetail(protocol.LanguageEnglish, "", "", "", "", "", []string{})

	XSL.Init("text/xml")
	XSL.SetFileExtension("xsl")
	XSL.SetDetail(protocol.LanguageEnglish, "", "", "", "", "", []string{})

	XSD.Init("text/xml")
	XSD.SetFileExtension("xsd")
	XSD.SetDetail(protocol.LanguageEnglish, "", "", "", "", "", []string{})

}
