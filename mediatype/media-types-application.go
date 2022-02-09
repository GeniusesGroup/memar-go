/* For license and copyright information please see LEGAL file in repository */

package mediatype

import "../protocol"

var (
	ABW = New("application/x-abiword").SetFileExtension("abw").
		SetDetail(protocol.LanguageEnglish, "AbiWord document", "", "", "", "", []string{})
	ARC = New("application/octet-stream").SetFileExtension("arc").
		SetDetail(protocol.LanguageEnglish, "Archive document (multiple files embedded)", "", "", "", "", []string{})

	BIN = New("application/octet-stream").SetFileExtension("bin").
		SetDetail(protocol.LanguageEnglish, "Any kind of binary data", "", "", "", "", []string{})
	BZ = New("application/x-bzip").SetFileExtension("bz").
		SetDetail(protocol.LanguageEnglish, "BZip archive", "", "", "", "", []string{})
	BZ2 = New("application/x-bzip2").SetFileExtension("bz2").
		SetDetail(protocol.LanguageEnglish, "BZip2 archive", "", "", "", "", []string{})

	CSH = New("application/x-csh").SetFileExtension("csh").
		SetDetail(protocol.LanguageEnglish, "C-Shell script", "", "", "", "", []string{})

	DOC = New("application/msword").SetFileExtension("doc").
		SetDetail(protocol.LanguageEnglish, "Microsoft Word", "", "", "", "", []string{})

	EPUB = New("application/epub+zip").SetFileExtension("epub").
		SetDetail(protocol.LanguageEnglish, "Electronic publication", "", "", "", "", []string{})

	GZ = New("application/gzip").SetFileExtension("gz").
		SetDetail(protocol.LanguageEnglish, "",
			"strictly speaking under MIME gzip would only be used as an encoding, not a content-type, but it's common to have .gz files",
			"",
			"",
			"",
			[]string{})

	HTTP = New("application/http").SetFileExtension("http"). // , "https://www.iana.org/assignments/media-types/application/http")
		SetDetail(protocol.LanguageEnglish, "Hypertext Transfer Protocol",
			"An application layer protocol in the Internet protocol suite model for distributed, collaborative, hypermedia information",
			"",
			"",
			"",
			[]string{})
	HTTPRequest = New("application/http; request").SetFileExtension("reqhttp").
			SetDetail(protocol.LanguageEnglish, "Hypertext Transfer Protocol Request", "", "", "", "", []string{})
	HTTPResponse = New("application/http; response").SetFileExtension("reshttp").
			SetDetail(protocol.LanguageEnglish, "Hypertext Transfer Protocol Response", "", "", "", "", []string{})

	JAR = New("application/java-archive").SetFileExtension("jar").
		SetDetail(protocol.LanguageEnglish, "Java Archive", "", "", "", "", []string{})
	JavaScript = New("application/javascript").SetFileExtension("js").
			SetDetail(protocol.LanguageEnglish, "JavaScript (ECMAScript) programming language", "", "", "", "", []string{})
	JSON = New("application/json").SetFileExtension("json").
		SetDetail(protocol.LanguageEnglish, "JavaScript Object Notation format", "", "", "", "", []string{})

	OGX = New("application/ogg").SetFileExtension("ogx").
		SetDetail(protocol.LanguageEnglish, "OGG", "", "", "", "", []string{})

	PDF = New("application/pdf").SetFileExtension("pdf").
		SetDetail(protocol.LanguageEnglish, "Adobe Portable Document Format", "", "", "", "", []string{})
	// TODO::: Can it be a mediatype?? Due it can't represent its data without schema
	Protobuf = New("application/protobuf").SetFileExtension("protobuf").
			SetInfo(protocol.Software_PreAlpha, 0, "https://datatracker.ietf.org/doc/html/draft-rfernando-protocol-buffers-00").
			SetDetail(protocol.LanguageEnglish, "", "", "", "", "", []string{})

	RAR = New("application/x-rar-compressed").SetFileExtension("rar").
		SetDetail(protocol.LanguageEnglish, "RAR archive", "", "", "", "", []string{})
	RTF = New("application/rtf").SetFileExtension("rtf").
		SetDetail(protocol.LanguageEnglish, "Rich Text Format", "", "", "", "", []string{})

	SRPC = New("application/srpc").SetFileExtension("srpc").
		SetDetail(protocol.LanguageEnglish, "Syllab Remote procedure call protocol", "", "", "", "", []string{})
	Syllab = New("application/syllab").SetFileExtension("syllab").
		SetDetail(protocol.LanguageEnglish, "Syllab codec protocol", "", "", "", "", []string{})
	SevenZ = New("application/x-7z-compressed").SetFileExtension("7z").
		SetDetail(protocol.LanguageEnglish, "7-zip archive", "", "", "", "", []string{})
	SH = New("application/x-sh").SetFileExtension("sh").
		SetDetail(protocol.LanguageEnglish, "Bourne shell script", "", "", "", "", []string{})
	SWF = New("application/x-shockwave-flash").SetFileExtension("swf").
		SetDetail(protocol.LanguageEnglish, "Small web format (SWF) or Adobe Flash document", "", "", "", "", []string{})

	TAR = New("application/tar").SetFileExtension("tar").
		SetDetail(protocol.LanguageEnglish, "Tape Archive", "", "", "", "", []string{})

	URI = New("application/uri").SetFileExtension("uri").
		SetDetail(protocol.LanguageEnglish, "URI", "", "", "", "", []string{})

	WASM = New("application/wasm").SetFileExtension("wasm").
		SetDetail(protocol.LanguageEnglish, "WebAssembly",
			"WebAssembly (abbreviated Wasm) is a binary instruction format for a stack-based virtual machine",
			"", "", "",
			[]string{})

	XML = New("application/xml").SetFileExtension("xml").
		SetDetail(protocol.LanguageEnglish, "XML format", "", "", "", "", []string{})
	XHTML = New("application/xhtml+xml").SetFileExtension("xhtml").
		SetDetail(protocol.LanguageEnglish, "", "", "", "", "", []string{})

	ZIP = New("application/zip").SetFileExtension("zip").
		SetDetail(protocol.LanguageEnglish, "ZIP archive", "", "", "", "", []string{})
	ZZ = New("application/defalate").SetFileExtension("zz"). // https://fileinfo.com/extension/zz
		SetDetail(protocol.LanguageEnglish, "zlib archive", "", "", "", "", []string{})
)

// Vendor specific application mediatype
var (
	AZW = New("application/vnd.amazon.ebook").SetFileExtension("azw").
		SetDetail(protocol.LanguageEnglish, "Amazon Kindle eBook format", "", "", "", "", []string{})

	LOG = New("application/vnd.log.protocol+syllab").SetFileExtension("log")

	MPKG = New("application/vnd.apple.installer+xml").SetFileExtension("mpkg").
		SetDetail(protocol.LanguageEnglish, "Apple Installer Package", "", "", "", "", []string{})

	ODP = New("application/vnd.oasis.opendocument.presentation").SetFileExtension("odp").
		SetDetail(protocol.LanguageEnglish, "OpenDocument presentation document", "", "", "", "", []string{})
	ODS = New("application/vnd.oasis.opendocument.spreadsheet").SetFileExtension("ods").
		SetDetail(protocol.LanguageEnglish, "OpenDocument spreadsheet document", "", "", "", "", []string{})
	ODT = New("application/vnd.oasis.opendocument.text").SetFileExtension("odt").
		SetDetail(protocol.LanguageEnglish, "OpenDocument text document", "", "", "", "", []string{})

	PPT = New("application/vnd.ms-powerpoint").SetFileExtension("ppt").
		SetDetail(protocol.LanguageEnglish, "Microsoft PowerPoint", "", "", "", "", []string{})

	VSD = New("application/vnd.visio").SetFileExtension("vsd").
		SetDetail(protocol.LanguageEnglish, "Microsoft Visio", "", "", "", "", []string{})

	XLS = New("application/vnd.ms-excel").SetFileExtension("xls").
		SetDetail(protocol.LanguageEnglish, "Microsoft Excel", "", "", "", "", []string{})
	XUL = New("application/vnd.mozilla.xul+xml").SetFileExtension("xul")
)
