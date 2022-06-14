/* For license and copyright information please see LEGAL file in repository */

package mediatypes

import (
	"../mediatype"
	"../protocol"
)

var (
	AVI mediatype.MediaType

	MPEG mediatype.MediaType

	OGV mediatype.MediaType

	ThreeGPVideo mediatype.MediaType
	ThreeG2Video mediatype.MediaType

	WEBM mediatype.MediaType
)

func init() {
	AVI.Init("video/x-msvideo")
	AVI.SetFileExtension("avi")
	AVI.SetDetail(protocol.LanguageEnglish, "Audio Video Interleave", "", "", "", "", []string{})

	MPEG.Init("video/mpeg")
	MPEG.SetFileExtension("mpeg")
	MPEG.SetDetail(protocol.LanguageEnglish, "MPEG Video", "", "", "", "", []string{})

	OGV.Init("video/ogg")
	OGV.SetFileExtension("ogg")
	OGV.SetDetail(protocol.LanguageEnglish, "OGG video", "", "", "", "", []string{})

	ThreeGPVideo.Init("video/3gpp")
	ThreeGPVideo.SetFileExtension("3gp")
	ThreeGPVideo.SetDetail(protocol.LanguageEnglish, "3GPP video container", "", "", "", "", []string{})

	ThreeG2Video.Init("video/3gpp2")
	ThreeG2Video.SetFileExtension("3g2")
	ThreeG2Video.SetDetail(protocol.LanguageEnglish, "3GPP2 video container", "", "", "", "", []string{})

	WEBM.Init("video/webm")
	WEBM.SetFileExtension("webm")
	WEBM.SetDetail(protocol.LanguageEnglish, "WEBM video", "", "", "", "", []string{})
}
