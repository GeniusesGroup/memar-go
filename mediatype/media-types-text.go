/* For license and copyright information please see LEGAL file in repository */

package mediatype

import "../protocol"

var (
	ASC = New("text/plain", "asc").
		SetDetail(protocol.LanguageEnglish, "", "", []string{})

	CSS = New("text/css", "css").
		SetDetail(protocol.LanguageEnglish, "Cascading Style Sheets", "", []string{})
	CSV = New("text/csv", "csv").
		SetDetail(protocol.LanguageEnglish, "Comma-separated values", "", []string{})

	TEXT = New("text/plain", "text").
		SetDetail(protocol.LanguageEnglish, "", "", []string{})
	TXT = New("text/plain", "txt").
		SetDetail(protocol.LanguageEnglish, "", "", []string{})

	HTM = New("text/html", "htm").
		SetDetail(protocol.LanguageEnglish, "HyperText Markup Language", "", []string{})
	HTML = New("text/html", "html").
		SetDetail(protocol.LanguageEnglish, "HyperText Markup Language", "", []string{})

	TSV = New("text/tab-separated-values", "tsv").
		SetDetail(protocol.LanguageEnglish, "", "", []string{})

	XSL = New("text/xml", "xsl").
		SetDetail(protocol.LanguageEnglish, "", "", []string{})
	XSD = New("text/xml", "xsd").
		SetDetail(protocol.LanguageEnglish, "", "", []string{})

	ICS = New("text/calendar", "ics").
		SetDetail(protocol.LanguageEnglish, "iCalendar format", "", []string{})
)
