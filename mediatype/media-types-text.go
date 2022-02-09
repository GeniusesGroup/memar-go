/* For license and copyright information please see LEGAL file in repository */

package mediatype

import "../protocol"

var (
	ASC = New("text/plain").SetFileExtension("asc").
		SetDetail(protocol.LanguageEnglish, "", "", "", "", "", []string{})

	CSS = New("text/css").SetFileExtension("css").
		SetDetail(protocol.LanguageEnglish, "Cascading Style Sheets", "", "", "", "", []string{})
	CSV = New("text/csv").SetFileExtension("csv").
		SetDetail(protocol.LanguageEnglish, "Comma-separated values", "", "", "", "", []string{})

	TEXT = New("text/plain").SetFileExtension("text").
		SetDetail(protocol.LanguageEnglish, "", "", "", "", "", []string{})
	TXT = New("text/plain").SetFileExtension("txt").
		SetDetail(protocol.LanguageEnglish, "", "", "", "", "", []string{})

	HTM = New("text/html").SetFileExtension("htm").
		SetDetail(protocol.LanguageEnglish, "HyperText Markup Language", "", "", "", "", []string{})
	HTML = New("text/html").SetFileExtension("html").
		SetDetail(protocol.LanguageEnglish, "HyperText Markup Language", "", "", "", "", []string{})

	TSV = New("text/tab-separated-values").SetFileExtension("tsv").
		SetDetail(protocol.LanguageEnglish, "", "", "", "", "", []string{})

	XSL = New("text/xml").SetFileExtension("xsl").
		SetDetail(protocol.LanguageEnglish, "", "", "", "", "", []string{})
	XSD = New("text/xml").SetFileExtension("xsd").
		SetDetail(protocol.LanguageEnglish, "", "", "", "", "", []string{})

	ICS = New("text/calendar").SetFileExtension("ics").
		SetDetail(protocol.LanguageEnglish, "iCalendar format", "", "", "", "", []string{})
)
