/* For license and copyright information please see LEGAL file in repository */

package mediatypes

import (
	"../mediatype"
	"../protocol"
)

var (
	GIF mediatype.MediaType

	JPG  mediatype.MediaType
	JPEG mediatype.MediaType

	PNG mediatype.MediaType

	SVG mediatype.MediaType

	ICO mediatype.MediaType

	WEBP mediatype.MediaType

	TIF  mediatype.MediaType
	TIFF mediatype.MediaType
)

func init() {
	GIF.Init("image/gif")
	GIF.SetFileExtension("gif")
	GIF.SetDetail(protocol.LanguageEnglish, "Graphics Interchange Format", "", "", "", "", []string{})

	JPG.Init("image/jpeg")
	JPG.SetFileExtension("jpg")
	JPG.SetDetail(protocol.LanguageEnglish, "JPEG images", "", "", "", "", []string{})

	JPEG.Init("image/jpeg")
	JPEG.SetFileExtension("jpeg")
	JPEG.SetDetail(protocol.LanguageEnglish, "JPEG images", "", "", "", "", []string{})

	PNG.Init("image/png")
	PNG.SetFileExtension("png")
	PNG.SetDetail(protocol.LanguageEnglish, "Portable Network Graphics", "", "", "", "", []string{})

	SVG.Init("image/svg+xml")
	SVG.SetFileExtension("svg")
	SVG.SetDetail(protocol.LanguageEnglish, " Scalable Vector Graphics", "", "", "", "", []string{})

	ICO.Init("image/x-icon")
	ICO.SetFileExtension("ico")
	ICO.SetDetail(protocol.LanguageEnglish, "Icon format", "", "", "", "", []string{})

	WEBP.Init("image/webp")
	WEBP.SetFileExtension("webp")
	WEBP.SetDetail(protocol.LanguageEnglish, "WEBP image", "", "", "", "", []string{})

	TIF.Init("image/tiff")
	TIF.SetFileExtension("tif")
	TIF.SetDetail(protocol.LanguageEnglish, "Tagged Image File Format", "", "", "", "", []string{})

	TIFF.Init("image/tiff")
	TIFF.SetFileExtension("tiff")
	TIFF.SetDetail(protocol.LanguageEnglish, "Tagged Image File Format", "", "", "", "", []string{})
}
