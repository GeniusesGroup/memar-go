/* For license and copyright information please see LEGAL file in repository */

package iana

// MediaTypeByID use to retrieve MediaType data by ID
var MediaTypeByID map[uint16]MediaType

// MediaTypeByName use to retrieve MediaType data by name
var MediaTypeByName map[string]MediaType

// MediaTypeByType use to retrieve MediaType data by type
var MediaTypeByType map[string]MediaType

// MediaTypeBySubType use to retrieve MediaType data by sub type
var MediaTypeBySubType map[string]MediaType

// MediaType or MimeType standrad list can found here
// http://www.iana.org/assignments/media-types/media-types.xhtml
// https://en.wikipedia.org/wiki/Media_type
type MediaType struct {
	ID             uint16 // Use ID instead of other data to improve efficiency!
	Name           string // Use as file extension in windows, ...!
	Type           string // The currently registered types are: application, audio, example, font, image, message, model, multipart, text and video.
	SubType        string
	Description    string
	Reference      string
	RegisteredDate int64
	ApprovedDate   int64
}

func init() {
	MediaTypeByID = make(map[uint16]MediaType)
	MediaTypeByName = make(map[string]MediaType)
	MediaTypeByType = make(map[string]MediaType)
	MediaTypeBySubType = make(map[string]MediaType)

	// MediaTypeMap[0] = MediaType{0, "", "", "", "", "", 0, 0}
	MediaTypeByID[0] = MediaType{0, "TEXT", "text", "plain", "", "", 0, 0}
	MediaTypeByID[1] = MediaType{1, "TXT", "text", "plain", "", "", 0, 0}
	MediaTypeByID[2] = MediaType{2, "HTM", "text", "html", "HyperText Markup Language", "", 0, 0}
	MediaTypeByID[3] = MediaType{3, "HTML", "text", "html", "HyperText Markup Language", "", 0, 0}
	MediaTypeByID[4] = MediaType{4, "CSS", "text", "css", "Cascading Style Sheets", "", 0, 0}
	MediaTypeByID[5] = MediaType{5, "CSV", "text", "csv", "Comma-separated values", "", 0, 0}
	MediaTypeByID[6] = MediaType{6, "ASC", "text", "plain", "", "", 0, 0}
	MediaTypeByID[7] = MediaType{7, "TSV", "text", "tab-separated-values", "", "", 0, 0}
	MediaTypeByID[8] = MediaType{8, "XSL", "text", "xml", "", "", 0, 0}
	MediaTypeByID[9] = MediaType{9, "XSD", "text", "xml", "", "", 0, 0}
	MediaTypeByID[10] = MediaType{10, "ICS", "text", "calendar", "iCalendar format", "", 0, 0}
	MediaTypeByID[11] = MediaType{11, "EML", "message", "rfc822", "", "", 0, 0}
	MediaTypeByID[12] = MediaType{12, "WOFF", "font", "woff", "Web Open Font Format", "", 0, 0}
	MediaTypeByID[13] = MediaType{13, "WOFF2", "font", "woff2", "Web Open Font Format", "", 0, 0}
	MediaTypeByID[14] = MediaType{14, "TTF", "font", "ttf", "TrueType Font", "", 0, 0}
	MediaTypeByID[15] = MediaType{15, "GIF", "image", "gif", "Graphics Interchange Format", "", 0, 0}
	MediaTypeByID[16] = MediaType{16, "JPG", "image", "jpeg", "JPEG images", "", 0, 0}
	MediaTypeByID[17] = MediaType{17, "JPEG", "image", "jpeg", "JPEG images", "", 0, 0}
	MediaTypeByID[18] = MediaType{18, "PNG", "image", "png", "Portable Network Graphics", "", 0, 0}
	MediaTypeByID[19] = MediaType{19, "SVG", "image", "svg+xml", "Scalable Vector Graphics", "", 0, 0}
	MediaTypeByID[20] = MediaType{20, "ICO", "image", "x-icon", "Icon format", "", 0, 0}
	MediaTypeByID[21] = MediaType{21, "WEBP", "image", "webp", "WEBP image", "", 0, 0}
	MediaTypeByID[22] = MediaType{22, "TIF", "image", "tiff", "Tagged Image File Format", "", 0, 0}
	MediaTypeByID[23] = MediaType{23, "TIFF", "image", "tiff", "Tagged Image File Format", "", 0, 0}
	MediaTypeByID[24] = MediaType{2, "AAC", "audio", "aac", "AAC audio file", "", 0, 0}
	MediaTypeByID[25] = MediaType{2, "WAV", "audio", "x-wav", "Waveform Audio Format", "", 0, 0}
	MediaTypeByID[26] = MediaType{2, "WEBA", "audio", "webm", "WEBM audio", "", 0, 0}
	MediaTypeByID[27] = MediaType{2, "OGA", "audio", "ogg", "OGG audio", "", 0, 0}
	MediaTypeByID[28] = MediaType{2, "MID", "audio", "midi", "Musical Instrument Digital Interface", "", 0, 0}
	MediaTypeByID[29] = MediaType{2, "MIDI", "audio", "midi", "Musical Instrument Digital Interface", "", 0, 0}
	MediaTypeByID[30] = MediaType{30, "WEBM", "video", "webm", "WEBM video", "", 0, 0}
	MediaTypeByID[31] = MediaType{31, "AVI", "video", "x-msvideo", "Audio Video Interleave", "", 0, 0}
	MediaTypeByID[32] = MediaType{32, "MPEG", "video", "mpeg", "MPEG Video", "", 0, 0}
	MediaTypeByID[33] = MediaType{33, "OGV", "video", "ogg", "OGG video", "", 0, 0}
	MediaTypeByID[34] = MediaType{34, "3GP", "video", "3gpp", "3GPP audio/video container", "audio/3gpp if it doesn't contain video", 0, 0}
	MediaTypeByID[35] = MediaType{35, "3G2", "video", "3gpp2", "3GPP2 audio/video container", "audio/3gpp2 if it doesn't contain video", 0, 0}
	MediaTypeByID[36] = MediaType{36, "ARC", "application", "octet-stream", "Archive document", "multiple files embedded", 0, 0}
	MediaTypeByID[37] = MediaType{37, "JSON", "application", "json", "JSON format", "", 0, 0}
	MediaTypeByID[38] = MediaType{38, "XML", "application", "xml", "XML format", "", 0, 0}
	MediaTypeByID[39] = MediaType{39, "XHTML", "application", "xhtml+xml", "XHTML", "", 0, 0}
	MediaTypeByID[40] = MediaType{40, "JS", "application", "javascript", "JavaScript", "ECMAScript", 0, 0}
	MediaTypeByID[41] = MediaType{41, "PDF", "application", "pdf", "Adobe Portable Document Format", "", 0, 0}
	MediaTypeByID[42] = MediaType{42, "GZ", "application", "gzip", "strictly speaking under MIME gzip would only be used as an encoding, not a content-type, but it's common to have .gz files", "", 0, 0}
	MediaTypeByID[43] = MediaType{43, "TAR", "application", "tar", "Tape Archive", "", 0, 0}
	MediaTypeByID[44] = MediaType{44, "ABW", "application", "x-abiword", "AbiWord document", "", 0, 0}
	MediaTypeByID[45] = MediaType{45, "AZW", "application", "vnd.amazon.ebook", "Amazon Kindle eBook format", "", 0, 0}
	MediaTypeByID[46] = MediaType{46, "BIN", "application", "octet-stream", "Any kind of binary data", "", 0, 0}
	MediaTypeByID[47] = MediaType{47, "BZ", "application", "x-bzip", "BZip archive", "", 0, 0}
	MediaTypeByID[48] = MediaType{48, "BZ2", "application", "x-bzip2", "BZip2 archive", "", 0, 0}
	MediaTypeByID[49] = MediaType{49, "CSH", "application", "x-csh", "C-Shell script", "", 0, 0}
	MediaTypeByID[50] = MediaType{50, "DOC", "application", "msword", "Microsoft Word", "", 0, 0}
	MediaTypeByID[51] = MediaType{51, "EPUB", "application", "epub+zip", "Electronic publication", "", 0, 0}
	MediaTypeByID[52] = MediaType{52, "MPKG", "application", "vnd.apple.installer+xml", "Apple Installer Package", "", 0, 0}
	MediaTypeByID[53] = MediaType{53, "ODP", "application", "vnd.oasis.opendocument.presentation", "OpenDocuemnt presentation document", "", 0, 0}
	MediaTypeByID[54] = MediaType{54, "ODS", "application", "vnd.oasis.opendocument.spreadsheet", "OpenDocuemnt spreadsheet document", "", 0, 0}
	MediaTypeByID[55] = MediaType{55, "ODT", "application", "vnd.oasis.opendocument.text", "OpenDocument text document", "", 0, 0}
	MediaTypeByID[56] = MediaType{56, "OGX", "application", "ogg", "", "", 0, 0}
	MediaTypeByID[57] = MediaType{57, "PPT", "application", "vnd.ms-powerpoint", "Microsoft PowerPoint", "", 0, 0}
	MediaTypeByID[58] = MediaType{58, "RAR", "application", "x-rar-compressed", "RAR archive", "", 0, 0}
	MediaTypeByID[59] = MediaType{59, "RTF", "application", "rtf", "Rich Text Format", "", 0, 0}
	MediaTypeByID[60] = MediaType{60, "SH", "application", "x-sh", "Bourne shell script", "", 0, 0}
	MediaTypeByID[61] = MediaType{61, "SWF", "application", "x-shockwave-flash", "Small web format (SWF) or Adobe Flash document", "", 0, 0}
	MediaTypeByID[62] = MediaType{62, "ZIP", "application", "zip", "ZIP archive", "", 0, 0}
	MediaTypeByID[63] = MediaType{63, "JAR", "application", "java-archive", "Java Archive", "", 0, 0}
	MediaTypeByID[64] = MediaType{64, "VSD", "application", "vnd.visio", "Microsoft Visio", "", 0, 0}
	MediaTypeByID[65] = MediaType{65, "XLS", "application", "vnd.ms-excel", "Microsoft Excel", "", 0, 0}
	MediaTypeByID[66] = MediaType{66, "XUL", "application", "vnd.mozilla.xul+xml", "", "", 0, 0}
	MediaTypeByID[67] = MediaType{67, "7Z", "application", "x-7z-compressed", "7-zip archive", "", 0, 0}
	MediaTypeByID[68] = MediaType{68, "", "", "", "", "", 0, 0}
	MediaTypeByID[69] = MediaType{69, "", "", "", "", "", 0, 0}
	MediaTypeByID[70] = MediaType{70, "", "", "", "", "", 0, 0}
	MediaTypeByID[71] = MediaType{72, "", "", "", "", "", 0, 0}

	for _, mediaType := range MediaTypeByID {
		MediaTypeByName[mediaType.Name] = mediaType
		MediaTypeByType[mediaType.Type] = mediaType
		MediaTypeBySubType[mediaType.SubType] = mediaType
	}
}
