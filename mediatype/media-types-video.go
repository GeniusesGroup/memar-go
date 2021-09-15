/* For license and copyright information please see LEGAL file in repository */

package mediatype

var (
	WEBM         = newMediaType("", "video/webm", "webm", "WEBM video")
	AVI          = newMediaType("", "video/x-msvideo", "avi", "Audio Video Interleave")
	MPEG         = newMediaType("", "video/mpeg", "mpeg", "MPEG Video")
	OGV          = newMediaType("", "video/ogg", "ogg", "OGG video")
	ThreeGPVideo = newMediaType("", "video/3gpp", "3gp", "3GPP video container")
	ThreeG2Video = newMediaType("", "video/3gpp2", "3g2", "3GPP2 video container")
)
