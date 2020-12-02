/* For license and copyright information please see LEGAL file in repository */

package www

import (
	"bytes"

	"../assets"
)

// mixCSSToHTML add given CSS file to end of HTML file and returns new html file.
func mixCSSToHTML(htmlFile, cssFile *assets.File) *assets.File {
	cssFile.Minify()

	var f = htmlFile.Copy()
	f.Data = make([]byte, len(htmlFile.Data), len(htmlFile.Data)+len(cssFile.Data)+len("<style></style>"))
	copy(f.Data, htmlFile.Data)
	f.Data = append(f.Data, "<style>"...)
	f.Data = append(f.Data, cssFile.Data...)
	f.Data = append(f.Data, "</style>"...)
	return f
}

// mixCSSToJS add given CSS file to specific part of JS file and returns new js file.
func mixCSSToJS(jsFile, cssFile *assets.File) (f *assets.File) {
	cssFile.Minify()
	f = jsFile.Copy()

	var funcLoc = bytes.Index(jsFile.Data, []byte("CSS: '"))
	if funcLoc < 0 {
		var minifiedCSS = append([]byte(`CSS = '`), cssFile.Data...)
		f.Data = bytes.Replace(jsFile.Data, []byte(`CSS = '`), minifiedCSS, 1)
	} else {
		var minifiedCSS = append([]byte(`CSS: '`), cssFile.Data...)
		f.Data = bytes.Replace(jsFile.Data, []byte(`CSS: '`), minifiedCSS, 1)
	}

	// f.Data = make([]byte, 0, len(jsFile.Data)+len(minifiedCSS))
	return f
}
