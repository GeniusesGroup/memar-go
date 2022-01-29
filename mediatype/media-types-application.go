/* For license and copyright information please see LEGAL file in repository */

package mediatype

import "../protocol"

var (
	ABW = New("application/x-abiword", "abw").
		SetDetail(protocol.LanguageEnglish, "AbiWord document", "", []string{})
	ARC = New("application/octet-stream", "arc").
		SetDetail(protocol.LanguageEnglish, "Archive document (multiple files embedded)", "", []string{})
	AZW = New("application/vnd.amazon.ebook", "azw").
		SetDetail(protocol.LanguageEnglish, "Amazon Kindle eBook format", "", []string{})

	BIN = New("application/octet-stream", "bin").
		SetDetail(protocol.LanguageEnglish, "Any kind of binary data", "", []string{})
	BZ = New("application/x-bzip", "bz").
		SetDetail(protocol.LanguageEnglish, "BZip archive", "", []string{})
	BZ2 = New("application/x-bzip2", "bz2").
		SetDetail(protocol.LanguageEnglish, "BZip2 archive", "", []string{})

	CSH = New("application/x-csh", "csh").
		SetDetail(protocol.LanguageEnglish, "C-Shell script", "", []string{})

	DOC = New("application/msword", "doc").
		SetDetail(protocol.LanguageEnglish, "Microsoft Word", "", []string{})

	EPUB = New("application/epub+zip", "epub").
		SetDetail(protocol.LanguageEnglish, "Electronic publication", "", []string{})

	GZ = New("application/gzip", "gz").
		SetDetail(protocol.LanguageEnglish, "", "strictly speaking under MIME gzip would only be used as an encoding, not a content-type, but it's common to have .gz files", []string{})

	HTTP = New("application/http", "http"). // , "https://www.iana.org/assignments/media-types/application/http")
		SetDetail(protocol.LanguageEnglish, "Hypertext Transfer Protocol", "An application layer protocol in the Internet protocol suite model for distributed, collaborative, hypermedia information", []string{})
	HTTPRequest = New("application/http; request", "http").
			SetDetail(protocol.LanguageEnglish, "Hypertext Transfer Protocol Request", "", []string{})
	HTTPResponse = New("application/http; response", "http").
			SetDetail(protocol.LanguageEnglish, "Hypertext Transfer Protocol Response", "", []string{})

	JAR = New("application/java-archive", "jar").
		SetDetail(protocol.LanguageEnglish, "Java Archive", "", []string{})
	JavaScript = New("application/javascript", "js").
			SetDetail(protocol.LanguageEnglish, "JavaScript (ECMAScript) programming language", "", []string{})
	JSON = New("application/json", "json").
		SetDetail(protocol.LanguageEnglish, "JavaScript Object Notation format", "", []string{})

	LOG = New("application/vnd.log.protocol+syllab", "log")

	MPKG = New("application/vnd.apple.installer+xml", "mpkg").
		SetDetail(protocol.LanguageEnglish, "Apple Installer Package", "", []string{})

	ODP = New("application/vnd.oasis.opendocument.presentation", "odp").
		SetDetail(protocol.LanguageEnglish, "OpenDocument presentation document", "", []string{})
	ODS = New("application/vnd.oasis.opendocument.spreadsheet", "ods").
		SetDetail(protocol.LanguageEnglish, "OpenDocument spreadsheet document", "", []string{})
	ODT = New("application/vnd.oasis.opendocument.text", "odt").
		SetDetail(protocol.LanguageEnglish, "OpenDocument text document", "", []string{})
	OGX = New("application/ogg", "ogx").
		SetDetail(protocol.LanguageEnglish, "OGG", "", []string{})

	PDF = New("application/pdf", "pdf").
		SetDetail(protocol.LanguageEnglish, "Adobe Portable Document Format", "", []string{})
	PPT = New("application/vnd.ms-powerpoint", "ppt").
		SetDetail(protocol.LanguageEnglish, "Microsoft PowerPoint", "", []string{})
	// TODO::: Can it be a mediatype?? Due it can't represent its data without schema
	Protobuf = New("application/protobuf", "protobuf").
			SetInfo(protocol.Software_PreAlpha, 0, "https://datatracker.ietf.org/doc/html/draft-rfernando-protocol-buffers-00").
			SetDetail(protocol.LanguageEnglish, "", "", []string{})

	RAR = New("application/x-rar-compressed", "rar").
		SetDetail(protocol.LanguageEnglish, "RAR archive", "", []string{})
	RTF = New("application/rtf", "rtf").
		SetDetail(protocol.LanguageEnglish, "Rich Text Format", "", []string{})

	SRPC = New("application/srpc", "srpc").
		SetDetail(protocol.LanguageEnglish, "Syllab Remote procedure call protocol", "", []string{})
	Syllab = New("application/syllab", "syllab").
		SetDetail(protocol.LanguageEnglish, "Syllab codec protocol", "", []string{})
	SevenZ = New("application/x-7z-compressed", "7z").
		SetDetail(protocol.LanguageEnglish, "7-zip archive", "", []string{})
	SH = New("application/x-sh", "sh").
		SetDetail(protocol.LanguageEnglish, "Bourne shell script", "", []string{})
	SWF = New("application/x-shockwave-flash", "swf").
		SetDetail(protocol.LanguageEnglish, "Small web format (SWF) or Adobe Flash document", "", []string{})

	TAR = New("application/tar", "tar").
		SetDetail(protocol.LanguageEnglish, "Tape Archive", "", []string{})

	VSD = New("application/vnd.visio", "vsd").
		SetDetail(protocol.LanguageEnglish, "Microsoft Visio", "", []string{})

	URI = New("application/uri", "uri").
		SetDetail(protocol.LanguageEnglish, "URI", "", []string{})

	WASM = New("application/wasm", "wasm").
		SetDetail(protocol.LanguageEnglish, "WebAssembly", "WebAssembly (abbreviated Wasm) is a binary instruction format for a stack-based virtual machine", []string{})

	XLS = New("application/vnd.ms-excel", "xls").
		SetDetail(protocol.LanguageEnglish, "Microsoft Excel", "", []string{})
	XML = New("application/xml", "xml").
		SetDetail(protocol.LanguageEnglish, "XML format", "", []string{})
	XHTML = New("application/xhtml+xml", "xhtml").
		SetDetail(protocol.LanguageEnglish, "", "", []string{})
	XUL = New("application/vnd.mozilla.xul+xml", "xul").
		SetDetail(protocol.LanguageEnglish, "", "", []string{})

	ZIP = New("application/zip", "zip").
		SetDetail(protocol.LanguageEnglish, "ZIP archive", "", []string{})
	ZZ = New("application/defalate", "zz"). // https://fileinfo.com/extension/zz
		SetDetail(protocol.LanguageEnglish, "zlib archive", "", []string{})
)
