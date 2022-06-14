/* For license and copyright information please see LEGAL file in repository */

package mediatypes

import (
	"../mediatype"
	"../protocol"
)

var (
	ABW mediatype.MediaType
	ARC mediatype.MediaType
	BIN mediatype.MediaType
	BZ  mediatype.MediaType
	BZ2 mediatype.MediaType
	CSH mediatype.MediaType

	DOC mediatype.MediaType

	EPUB mediatype.MediaType

	GZ mediatype.MediaType

	HTTP         mediatype.MediaType
	HTTPRequest  mediatype.MediaType
	HTTPResponse mediatype.MediaType

	JAR        mediatype.MediaType
	JavaScript mediatype.MediaType
	JSON       mediatype.MediaType

	OGX mediatype.MediaType

	PDF mediatype.MediaType
	// TODO::: Can it be a mediatype?? Due it can't represent its data without schema
	Protobuf mediatype.MediaType

	RAR mediatype.MediaType
	RTF mediatype.MediaType

	SRPC   mediatype.MediaType
	Syllab mediatype.MediaType
	SevenZ mediatype.MediaType
	SH     mediatype.MediaType
	SWF    mediatype.MediaType

	TAR mediatype.MediaType

	URI mediatype.MediaType

	WASM mediatype.MediaType

	XML   mediatype.MediaType
	XHTML mediatype.MediaType

	ZIP mediatype.MediaType
	ZZ  mediatype.MediaType
)

// Vendor specific application mediatype
var (
	AZW mediatype.MediaType

	LOG mediatype.MediaType

	MPKG mediatype.MediaType

	ODP mediatype.MediaType
	ODS mediatype.MediaType
	ODT mediatype.MediaType

	PPT mediatype.MediaType

	VSD mediatype.MediaType

	XLS mediatype.MediaType
	XUL mediatype.MediaType
)

func init() {
	ABW.Init("application/x-abiword")
	ABW.SetFileExtension("abw")
	ABW.SetDetail(protocol.LanguageEnglish, "AbiWord document", "", "", "", "", []string{})

	ARC.Init("application/octet-stream")
	ARC.SetFileExtension("arc")
	ARC.SetDetail(protocol.LanguageEnglish, "Archive document (multiple files embedded)", "", "", "", "", []string{})

	BIN.Init("application/octet-stream")
	BIN.SetFileExtension("bin")
	BIN.SetDetail(protocol.LanguageEnglish, "Any kind of binary data", "", "", "", "", []string{})

	BZ.Init("application/x-bzip")
	BZ.SetFileExtension("bz")
	BZ.SetDetail(protocol.LanguageEnglish, "BZip archive", "", "", "", "", []string{})

	BZ2.Init("application/x-bzip2")
	BZ.SetFileExtension("bz2")
	BZ.SetDetail(protocol.LanguageEnglish, "BZip2 archive", "", "", "", "", []string{})

	CSH.Init("application/x-csh")
	CSH.SetFileExtension("csh")
	CSH.SetDetail(protocol.LanguageEnglish, "C-Shell script", "", "", "", "", []string{})

	DOC.Init("application/msword")
	DOC.SetFileExtension("doc")
	DOC.SetDetail(protocol.LanguageEnglish, "Microsoft Word", "", "", "", "", []string{})

	EPUB.Init("application/epub+zip")
	EPUB.SetFileExtension("epub")
	EPUB.SetDetail(protocol.LanguageEnglish, "Electronic publication", "", "", "", "", []string{})

	GZ.Init("application/gzip")
	GZ.SetFileExtension("gz")
	GZ.SetDetail(protocol.LanguageEnglish, "",
		"strictly speaking under MIME gzip would only be used as an encoding, not a content-type, but it's common to have .gz files",
		"",
		"",
		"",
		[]string{})

	HTTP.Init("application/http")
	HTTP.SetInfo(protocol.Software_PreAlpha, 0, "https://www.iana.org/assignments/media-types/application/http")
	HTTP.SetFileExtension("http")
	HTTP.SetDetail(protocol.LanguageEnglish, "Hypertext Transfer Protocol",
		"An application layer protocol in the Internet protocol suite model for distributed, collaborative, hypermedia information",
		"",
		"",
		"",
		[]string{})

	HTTPRequest.Init("application/http; request")
	HTTPRequest.SetFileExtension("reqhttp")
	HTTPRequest.SetDetail(protocol.LanguageEnglish, "Hypertext Transfer Protocol Request", "", "", "", "", []string{})

	HTTPResponse.Init("application/http; response")
	HTTPResponse.SetFileExtension("reshttp")
	HTTPResponse.SetDetail(protocol.LanguageEnglish, "Hypertext Transfer Protocol Response", "", "", "", "", []string{})

	JAR.Init("application/java-archive")
	JAR.SetFileExtension("jar")
	JAR.SetDetail(protocol.LanguageEnglish, "Java Archive", "", "", "", "", []string{})

	JavaScript.Init("application/javascript")
	JavaScript.SetFileExtension("js")
	JavaScript.SetDetail(protocol.LanguageEnglish, "JavaScript (ECMAScript) programming language", "", "", "", "", []string{})

	JSON.Init("application/json")
	JSON.SetFileExtension("json")
	JSON.SetDetail(protocol.LanguageEnglish, "JavaScript Object Notation format", "", "", "", "", []string{})

	OGX.Init("application/ogg")
	OGX.SetFileExtension("ogx")
	OGX.SetDetail(protocol.LanguageEnglish, "OGG", "", "", "", "", []string{})

	PDF.Init("application/pdf")
	PDF.SetFileExtension("pdf")
	PDF.SetDetail(protocol.LanguageEnglish, "Adobe Portable Document Format", "", "", "", "", []string{})

	Protobuf.Init("application/protobuf")
	Protobuf.SetFileExtension("protobuf")
	Protobuf.SetInfo(protocol.Software_PreAlpha, 0, "https://datatracker.ietf.org/doc/html/draft-rfernando-protocol-buffers-00")
	Protobuf.SetDetail(protocol.LanguageEnglish, "", "", "", "", "", []string{})

	RAR.Init("application/x-rar-compressed")
	RAR.SetFileExtension("rar")
	RAR.SetDetail(protocol.LanguageEnglish, "RAR archive", "", "", "", "", []string{})

	RTF.Init("application/rtf")
	RTF.SetFileExtension("rtf")
	RTF.SetDetail(protocol.LanguageEnglish, "Rich Text Format", "", "", "", "", []string{})

	SRPC.Init("application/srpc")
	SRPC.SetFileExtension("srpc")
	SRPC.SetDetail(protocol.LanguageEnglish, "Syllab Remote procedure call protocol", "", "", "", "", []string{})

	Syllab.Init("application/syllab")
	Syllab.SetFileExtension("syllab")
	Syllab.SetDetail(protocol.LanguageEnglish, "Syllab codec protocol", "", "", "", "", []string{})

	SevenZ.Init("application/x-7z-compressed")
	SevenZ.SetFileExtension("7z")
	SevenZ.SetDetail(protocol.LanguageEnglish, "7-zip archive", "", "", "", "", []string{})

	SH.Init("application/x-sh")
	SH.SetFileExtension("sh")
	SH.SetDetail(protocol.LanguageEnglish, "Bourne shell script", "", "", "", "", []string{})

	SWF.Init("application/x-shockwave-flash")
	SWF.SetFileExtension("swf")
	SWF.SetDetail(protocol.LanguageEnglish, "Small web format (SWF) or Adobe Flash document", "", "", "", "", []string{})

	TAR.Init("application/tar")
	TAR.SetFileExtension("tar")
	TAR.SetDetail(protocol.LanguageEnglish, "Tape Archive", "", "", "", "", []string{})

	URI.Init("application/uri")
	URI.SetFileExtension("uri")
	URI.SetDetail(protocol.LanguageEnglish, "URI", "", "", "", "", []string{})

	WASM.Init("application/wasm")
	WASM.SetFileExtension("wasm")
	WASM.SetDetail(protocol.LanguageEnglish, "WebAssembly",
		"WebAssembly (abbreviated Wasm) is a binary instruction format for a stack-based virtual machine",
		"", "", "",
		[]string{})

	XML.Init("application/xml")
	XML.SetFileExtension("xml")
	XML.SetDetail(protocol.LanguageEnglish, "XML format", "", "", "", "", []string{})

	XHTML.Init("application/xhtml+xml")
	XHTML.SetFileExtension("xhtml")
	XHTML.SetDetail(protocol.LanguageEnglish, "", "", "", "", "", []string{})

	ZIP.Init("application/zip")
	ZIP.SetFileExtension("zip")
	ZIP.SetDetail(protocol.LanguageEnglish, "ZIP archive", "", "", "", "", []string{})

	ZZ.Init("application/defalate")
	ZZ.SetFileExtension("zz") // https://fileinfo.com/extension/zz
	ZZ.SetDetail(protocol.LanguageEnglish, "zlib archive", "", "", "", "", []string{})

	// Vendor specific application mediatype

	AZW.Init("application/vnd.amazon.ebook")
	AZW.SetFileExtension("azw")
	AZW.SetDetail(protocol.LanguageEnglish, "Amazon Kindle eBook format", "", "", "", "", []string{})

	LOG.Init("application/vnd.log.protocol+syllab")
	LOG.SetFileExtension("log")

	MPKG.Init("application/vnd.apple.installer+xml")
	MPKG.SetFileExtension("mpkg")
	MPKG.SetDetail(protocol.LanguageEnglish, "Apple Installer Package", "", "", "", "", []string{})

	ODP.Init("application/vnd.oasis.opendocument.presentation")
	ODP.SetFileExtension("odp")
	ODP.SetDetail(protocol.LanguageEnglish, "OpenDocument presentation document", "", "", "", "", []string{})

	ODS.Init("application/vnd.oasis.opendocument.spreadsheet")
	ODS.SetFileExtension("ods")
	ODS.SetDetail(protocol.LanguageEnglish, "OpenDocument spreadsheet document", "", "", "", "", []string{})

	ODT.Init("application/vnd.oasis.opendocument.text")
	ODT.SetFileExtension("odt")
	ODT.SetDetail(protocol.LanguageEnglish, "OpenDocument text document", "", "", "", "", []string{})

	PPT.Init("application/vnd.ms-powerpoint")
	PPT.SetFileExtension("ppt")
	PPT.SetDetail(protocol.LanguageEnglish, "Microsoft PowerPoint", "", "", "", "", []string{})

	VSD.Init("application/vnd.visio")
	VSD.SetFileExtension("vsd")
	VSD.SetDetail(protocol.LanguageEnglish, "Microsoft Visio", "", "", "", "", []string{})

	XLS.Init("application/vnd.ms-excel")
	XLS.SetFileExtension("xls")
	XLS.SetDetail(protocol.LanguageEnglish, "Microsoft Excel", "", "", "", "", []string{})

	XUL.Init("application/vnd.mozilla.xul+xml")
	XUL.SetFileExtension("xul")
}
