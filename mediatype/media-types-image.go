/* For license and copyright information please see LEGAL file in repository */

package mediatype

import "../protocol"

var (
	GIF = New("image/gif", "gif").
		SetDetail(protocol.LanguageEnglish, "Graphics Interchange Format", "", []string{})

	JPG = New("image/jpeg", "jpg").
		SetDetail(protocol.LanguageEnglish, "JPEG images", "", []string{})
	JPEG = New("image/jpeg", "jpeg").
		SetDetail(protocol.LanguageEnglish, "JPEG images", "", []string{})

	PNG = New("image/png", "png").
		SetDetail(protocol.LanguageEnglish, "Portable Network Graphics", "", []string{})

	SVG = New("image/svg+xml", "svg").
		SetDetail(protocol.LanguageEnglish, " Scalable Vector Graphics", "", []string{})

	ICO = New("image/x-icon", "ico").
		SetDetail(protocol.LanguageEnglish, "Icon format", "", []string{})

	WEBP = New("image/webp", "webp").
		SetDetail(protocol.LanguageEnglish, "WEBP image", "", []string{})

	TIF = New("image/tiff", "tif").
		SetDetail(protocol.LanguageEnglish, "Tagged Image File Format", "", []string{})
	TIFF = New("image/tiff", "tiff").
		SetDetail(protocol.LanguageEnglish, "Tagged Image File Format", "", []string{})
)
