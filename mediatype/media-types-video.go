/* For license and copyright information please see LEGAL file in repository */

package mediatype

import "../protocol"

var (
	AVI = New("video/x-msvideo", "avi").
		SetDetail(protocol.LanguageEnglish, "Audio Video Interleave", "", []string{})

	MPEG = New("video/mpeg", "mpeg").
		SetDetail(protocol.LanguageEnglish, "MPEG Video", "", []string{})

	OGV = New("video/ogg", "ogg").
		SetDetail(protocol.LanguageEnglish, "OGG video", "", []string{})

	ThreeGPVideo = New("video/3gpp", "3gp").
			SetDetail(protocol.LanguageEnglish, "3GPP video container", "", []string{})
	ThreeG2Video = New("video/3gpp2", "3g2").
			SetDetail(protocol.LanguageEnglish, "3GPP2 video container", "", []string{})

	WEBM = New("video/webm", "webm").
		SetDetail(protocol.LanguageEnglish, "WEBM video", "", []string{})
)
