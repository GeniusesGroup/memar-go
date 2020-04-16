/* For license and copyright information please see LEGAL file in repository */

package http

// StandardMimeTypes : All standard mime types.
var StandardMimeTypes = struct {
	Application string
	Audio       string
	Font        string
	Image       string
	Message     string
	Text        string
	Video       string
}{
	Application: "application",
	Audio:       "audio",
	Font:        "font",
	Image:       "image",
	Message:     "message",
	Text:        "text",
	Video:       "video"}

// MimeTypes : All standard mime types.
var MimeTypes = map[string]string{
	"ARC":   "application/octet-stream",                        //Archive document (multiple files embedded)
	"JSON":  "application/json",                                //JSON format
	"XML":   "application/xml",                                 //XML format
	"XHTML": "application/xhtml+xml",                           //XHTML
	"JS":    "application/javascript",                          //JavaScript (ECMAScript)
	"PDF":   "application/pdf",                                 //Adobe Portable Document Format
	"GZ":    "application/gzip",                                // strictly speaking under MIME gzip would only be used as an encoding, not a content-type, but it's common to have .gz files
	"TAR":   "application/tar",                                 //Tape Archive
	"ABW":   "application/x-abiword",                           //AbiWord document
	"AZW":   "application/vnd.amazon.ebook",                    //Amazon Kindle eBook format
	"BIN":   "application/octet-stream",                        //Any kind of binary data
	"BZ":    "application/x-bzip",                              //BZip archive
	"BZ2":   "application/x-bzip2",                             //BZip2 archive
	"CSH":   "application/x-csh",                               //C-Shell script
	"DOC":   "application/msword",                              //Microsoft Word
	"EPUB":  "application/epub+zip",                            //Electronic publication
	"MPKG":  "application/vnd.apple.installer+xml",             //Apple Installer Package
	"ODP":   "application/vnd.oasis.opendocument.presentation", //OpenDocuemnt presentation document
	"ODS":   "application/vnd.oasis.opendocument.spreadsheet",  //OpenDocuemnt spreadsheet document
	"ODT":   "application/vnd.oasis.opendocument.text",         //OpenDocument text document
	"OGX":   "application/ogg",                                 //OGG
	"PPT":   "application/vnd.ms-powerpoint",                   //Microsoft PowerPoint
	"RAR":   "application/x-rar-compressed",                    //RAR archive
	"RTF":   "application/rtf",                                 //Rich Text Format
	"SH":    "application/x-sh",                                //Bourne shell script
	"SWF":   "application/x-shockwave-flash",                   //Small web format (SWF) or Adobe Flash document
	"ZIP":   "application/zip",                                 //ZIP archive
	"JAR":   "application/java-archive",                        //Java Archive
	"VSD":   "application/vnd.visio",                           //Microsoft Visio
	"XLS":   "application/vnd.ms-excel",                        //Microsoft Excel
	"XUL":   "application/vnd.mozilla.xul+xml",                 //XUL
	"7Z":    "application/x-7z-compressed",                     //7-zip archive

	"TEXT": "text/plain", //
	"TXT":  "text/plain", //
	"HTM":  "text/html",  //HyperText Markup Language
	"HTML": "text/html",  //HyperText Markup Language
	"CSS":  "text/css",   //Cascading Style Sheets
	"CSV":  "text/csv",   //Comma-separated values

	"GIF":  "image/gif",     //Graphics Interchange Format
	"JPG":  "image/jpeg",    //JPEG images
	"JPEG": "image/jpeg",    //JPEG images
	"PNG":  "image/png",     //Portable Network Graphics
	"SVG":  "image/svg+xml", //Scalable Vector Graphics
	"ICO":  "image/x-icon",  //Icon format
	"WEBP": "image/webp",    //WEBP image
	"TIF":  "image/tiff",    //Tagged Image File Format
	"TIFF": "image/tiff",    //Tagged Image File Format

	"ASC": "text/plain",                //
	"TSV": "text/tab-separated-values", //
	"XSL": "text/xml",                  //
	"XSD": "text/xml",                  //
	"ICS": "text/calendar",             //iCalendar format

	"AAC":  "audio/aac",   //AAC audio file
	"WAV":  "audio/x-wav", //Waveform Audio Format
	"WEBA": "audio/webm",  // WEBM audio
	"OGA":  "audio/ogg",   //OGG audio
	"MID":  "audio/midi",  //Musical Instrument Digital Interface
	"MIDI": "audio/midi",  //Musical Instrument Digital Interface

	"WEBM": "video/webm",      //WEBM video
	"AVI":  "video/x-msvideo", //Audio Video Interleave
	"MPEG": "video/mpeg",      //MPEG Video
	"OGV":  "video/ogg",       //OGG video
	"3GP":  "video/3gpp",      //3GPP audio/video container	||   audio/3gpp if it doesn't contain video
	"3G2":  "video/3gpp2",     //3GPP2 audio/video container	 ||  audio/3gpp2 if it doesn't contain video

	"WOFF":  "font/woff",  //Web Open Font Format
	"WOFF2": "font/woff2", //Web Open Font Format
	"TTF":   "font/ttf",   //TrueType Font

	"EML": "message/rfc822", //
}
