/* For license and copyright information please see LEGAL file in repository */

package www

import (
	"bytes"
	"regexp"

	"../assets"
)

var newLine = regexp.MustCompile(`\r?\n`)

// MinifyCSS use to minify css data.
func MinifyCSS(css []byte) []byte {
	var f = newLine.ReplaceAll(css, []byte{})
	return f
}

// mixCSSToHTML add given CSS file to end of HTML file and returns new html file.
func mixCSSToHTML(htmlFile, cssFile *assets.File) *assets.File {
	var minifiedCSS = MinifyCSS(cssFile.Data)

	var f = htmlFile.Copy()
	f.Data = make([]byte, len(htmlFile.Data), len(htmlFile.Data)+len(minifiedCSS)+len("<style></style>"))
	copy(f.Data, htmlFile.Data)
	f.Data = append(f.Data, "<style>"...)
	f.Data = append(f.Data, minifiedCSS...)
	f.Data = append(f.Data, "</style>"...)
	return f
}

// mixCSSToJS add given CSS file to specific part of JS file and returns new js file.
func mixCSSToJS(jsFile, cssFile *assets.File) *assets.File {
	var minifiedCSS = MinifyCSS(cssFile.Data)
	minifiedCSS = append([]byte(`CSS: '`), minifiedCSS...)

	var f = jsFile.Copy()
	// f.Data = make([]byte, 0, len(jsFile.Data)+len(minifiedCSS))
	f.Data = bytes.Replace(jsFile.Data, []byte(`CSS: '`), minifiedCSS, 1)
	return f
}
