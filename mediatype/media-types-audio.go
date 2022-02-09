/* For license and copyright information please see LEGAL file in repository */

package mediatype

import "../protocol"

var (
	AAC = New("audio/aac").SetFileExtension("aac").
		SetDetail(protocol.LanguageEnglish, "AAC audio file", "", "", "", "", []string{})

	WAV = New("audio/x-wav").SetFileExtension("wav").
		SetDetail(protocol.LanguageEnglish, "Waveform Audio Format", "", "", "", "", []string{})
	WEBA = New("audio/webm").SetFileExtension("webm").
		SetDetail(protocol.LanguageEnglish, "WEBM audio", "", "", "", "", []string{})

	OGA = New("audio/ogg").SetFileExtension("ogg").
		SetDetail(protocol.LanguageEnglish, "OGG audio", "", "", "", "", []string{})

	MID = New("audio/mid").SetFileExtension("mid").
		SetDetail(protocol.LanguageEnglish, "Musical Instrument Digital Interface", "", "", "", "", []string{})
	MIDI = New("audio/midi").SetFileExtension("midi").
		SetDetail(protocol.LanguageEnglish, "Musical Instrument Digital Interface", "", "", "", "", []string{})

	ThreeGPAudio = New("audio/3gpp").SetFileExtension("3gp").
			SetDetail(protocol.LanguageEnglish, "3GPP audio container", "", "", "", "", []string{})
	ThreeG2Audio = New("audio/3gpp2").SetFileExtension("3g2").
			SetDetail(protocol.LanguageEnglish, "3GPP2 audio container", "", "", "", "", []string{})
)
