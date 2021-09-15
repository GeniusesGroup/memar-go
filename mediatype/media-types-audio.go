/* For license and copyright information please see LEGAL file in repository */

package mediatype

var (
	AAC          = newMediaType("", "audio/aac", "aac", "AAC audio file")
	WAV          = newMediaType("", "audio/x-wav", "wav", "Waveform Audio Format")
	WEBA         = newMediaType("", "audio/webm", "webm", "WEBM audio")
	OGA          = newMediaType("", "audio/ogg", "ogg", "OGG audio")
	MID          = newMediaType("", "audio/mid", "mid", "Musical Instrument Digital Interface")
	MIDI         = newMediaType("", "audio/midi", "midi", "Musical Instrument Digital Interface")
	ThreeGPAudio = newMediaType("", "audio/3gpp", "3gp", "3GPP audio container")
	ThreeG2Audio = newMediaType("", "audio/3gpp2", "3g2", "3GPP2 audio container")
)
