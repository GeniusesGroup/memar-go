/* For license and copyright information please see LEGAL file in repository */

package www

import (
	"bytes"

	"../file"
)

// mixCSSToHTML add given CSS file to end of HTML file and returns new html file.
func mixCSSToHTML(htmlFile, cssFile *file.File) (mixFile *file.File) {
	cssFile.Minify()
	mixFile = htmlFile.Copy()
	var htmlData = htmlFile.Data()
	var cssData = cssFile.Data()

	var mixFileData = make([]byte, len(htmlData), len(htmlData)+len(cssData)+len("<style></style>"))
	copy(mixFileData, htmlData)
	mixFileData = append(mixFileData, "<style>"...)
	mixFileData = append(mixFileData, cssData...)
	mixFileData = append(mixFileData, "</style>"...)

	mixFile.SetData(mixFileData)
	return
}

// mixCSSToJS add given CSS file to specific part of JS file and returns new js file.
func mixCSSToJS(jsFile, cssFile *file.File) (mixFile *file.File) {
	cssFile.Minify()
	mixFile = jsFile.Copy()
	var jsData = jsFile.Data()
	var cssData = cssFile.Data()
	var mixFileData []byte

	var funcLoc = bytes.Index(jsData, []byte("CSS: '"))
	// mixFileData = make([]byte, 0, len(jsData)+len(minifiedCSS))
	if funcLoc < 0 {
		var minifiedCSS = append([]byte(`CSS = '`), cssData...)
		mixFileData = bytes.Replace(jsData, []byte(`CSS = '`), minifiedCSS, 1)
	} else {
		var minifiedCSS = append([]byte(`CSS: '`), cssData...)
		mixFileData = bytes.Replace(jsData, []byte(`CSS: '`), minifiedCSS, 1)
	}

	mixFile.SetData(mixFileData)
	return
}
