/* For license and copyright information please see LEGAL file in repository */

package captcha

// State use to indicate captcha state
type state uint8

// Captcha State
const (
	StateCreated  state = iota
	StateLastAnswerNotValid
	StateSolved
)

// Language indicate 
type Language uint8

// supported languages
const (
	LanguageEnglish Language = iota
)

// ImageFormat indicate 
type ImageFormat uint8

// Supported image format
const (
	ImageFormatPNG ImageFormat = iota
	ImageFormatJPEG
)

// AudioFormat indicate 
type AudioFormat uint8

// Supported audio format
const ()