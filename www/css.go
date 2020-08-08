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
func mixCSSToJS(jsFile, cssFile *assets.File) *assets.File {
	cssFile.Minify()

	var minifiedCSS = append([]byte(`CSS: '`), cssFile.Data...)

	var f = jsFile.Copy()
	// f.Data = make([]byte, 0, len(jsFile.Data)+len(minifiedCSS))
	f.Data = bytes.Replace(jsFile.Data, []byte(`CSS: '`), minifiedCSS, 1)
	return f
}
