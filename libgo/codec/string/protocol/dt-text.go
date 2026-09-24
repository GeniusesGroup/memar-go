/* For license and copyright information please see the LEGAL file in the code repository */

package go_string_p

type Field_Text interface {
	Text() Text
}

// Locale text(name, ...) in the Human world. It can be not unique in all quiddity contents.
// Text With Style. HTML & CSS (No JS??) is more expressive than markdown, suggest use them in article text to style text.
type Text string

// TODO::: dynamic data in title?? text substitution, text piping
