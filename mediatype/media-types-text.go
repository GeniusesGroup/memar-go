/* For license and copyright information please see LEGAL file in repository */

package mediatype

var (
	TEXT = newMediaType("", "text/plain", "text", "")
	TXT  = newMediaType("", "text/plain", "txt", "")
	HTM  = newMediaType("", "text/html", "htm", "HyperText Markup Language")
	HTML = newMediaType("", "text/html", "html", "HyperText Markup Language")
	CSS  = newMediaType("", "text/css", "css", "Cascading Style Sheets")
	CSV  = newMediaType("", "text/csv", "csv", "Comma-separated values")
	ASC  = newMediaType("", "text/plain", "asc", "")
	TSV  = newMediaType("", "text/tab-separated-values", "tsv", "")
	XSL  = newMediaType("", "text/xml", "xsl", "")
	XSD  = newMediaType("", "text/xml", "xsd", "")
	ICS  = newMediaType("", "text/calendar", "ics", "iCalendar format")
)
