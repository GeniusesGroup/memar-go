/* For license and copyright information please see LEGAL file in repository */

package mediatype

import "../protocol"

var (
	GIF = New("image/gif").SetFileExtension("gif").
		SetDetail(protocol.LanguageEnglish, "Graphics Interchange Format", "", "", "", "", []string{})

	JPG = New("image/jpeg").SetFileExtension("jpg").
		SetDetail(protocol.LanguageEnglish, "JPEG images", "", "", "", "", []string{})
	JPEG = New("image/jpeg").SetFileExtension("jpeg").
		SetDetail(protocol.LanguageEnglish, "JPEG images", "", "", "", "", []string{})

	PNG = New("image/png").SetFileExtension("png").
		SetDetail(protocol.LanguageEnglish, "Portable Network Graphics", "", "", "", "", []string{})

	SVG = New("image/svg+xml").SetFileExtension("svg").
		SetDetail(protocol.LanguageEnglish, " Scalable Vector Graphics", "", "", "", "", []string{})

	ICO = New("image/x-icon").SetFileExtension("ico").
		SetDetail(protocol.LanguageEnglish, "Icon format", "", "", "", "", []string{})

	WEBP = New("image/webp").SetFileExtension("webp").
		SetDetail(protocol.LanguageEnglish, "WEBP image", "", "", "", "", []string{})

	TIF = New("image/tiff").SetFileExtension("tif").
		SetDetail(protocol.LanguageEnglish, "Tagged Image File Format", "", "", "", "", []string{})
	TIFF = New("image/tiff").SetFileExtension("tiff").
		SetDetail(protocol.LanguageEnglish, "Tagged Image File Format", "", "", "", "", []string{})
)
