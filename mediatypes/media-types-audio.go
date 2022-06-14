/* For license and copyright information please see LEGAL file in repository */

package mediatypes

import (
	"../mediatype"
	"../protocol"
)

var (
	AAC mediatype.MediaType

	WAV  mediatype.MediaType
	WEBA mediatype.MediaType

	OGA mediatype.MediaType

	MID  mediatype.MediaType
	MIDI mediatype.MediaType

	ThreeGPAudio mediatype.MediaType
	ThreeG2Audio mediatype.MediaType
)

func init() {
	AAC.Init("audio/aac")
	AAC.SetFileExtension("aac")
	AAC.SetDetail(protocol.LanguageEnglish, "AAC audio file", "", "", "", "", []string{})

	WAV.Init("audio/x-wav")
	WAV.SetFileExtension("wav")
	WAV.SetDetail(protocol.LanguageEnglish, "Waveform Audio Format", "", "", "", "", []string{})

	WEBA.Init("audio/webm")
	WEBA.SetFileExtension("webm")
	WEBA.SetDetail(protocol.LanguageEnglish, "WEBM audio", "", "", "", "", []string{})

	OGA.Init("audio/ogg")
	OGA.SetFileExtension("ogg")
	OGA.SetDetail(protocol.LanguageEnglish, "OGG audio", "", "", "", "", []string{})

	MID.Init("audio/mid")
	MID.SetFileExtension("mid")
	MID.SetDetail(protocol.LanguageEnglish, "Musical Instrument Digital Interface", "", "", "", "", []string{})

	MIDI.Init("audio/midi")
	MIDI.SetFileExtension("midi")
	MIDI.SetDetail(protocol.LanguageEnglish, "Musical Instrument Digital Interface", "", "", "", "", []string{})

	ThreeGPAudio.Init("audio/3gpp")
	ThreeGPAudio.SetFileExtension("3gp")
	ThreeGPAudio.SetDetail(protocol.LanguageEnglish, "3GPP audio container", "", "", "", "", []string{})

	ThreeG2Audio.Init("audio/3gpp2")
	ThreeG2Audio.SetFileExtension("3g2")
	ThreeG2Audio.SetDetail(protocol.LanguageEnglish, "3GPP2 audio container", "", "", "", "", []string{})
}
