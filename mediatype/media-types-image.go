/* For license and copyright information please see LEGAL file in repository */

package mediatype

var (
	GIF  = newMediaType("", "image/gif", "gif", "Graphics Interchange Format")
	JPG  = newMediaType("", "image/jpeg", "jpg", "JPEG images")
	JPEG = newMediaType("", "image/jpeg", "jpeg", "JPEG images")
	PNG  = newMediaType("", "image/png", "png", "Portable Network Graphics")
	SVG  = newMediaType("", "image/svg+xml", "svg", " Scalable Vector Graphics")
	ICO  = newMediaType("", "image/x-icon", "ico", "Icon format")
	WEBP = newMediaType("", "image/webp", "webp", "WEBP image")
	TIF  = newMediaType("", "image/tiff", "tif", "Tagged Image File Format")
	TIFF = newMediaType("", "image/tiff", "tiff", "Tagged Image File Format")
)
