/* For license and copyright information please see LEGAL file in repository */

package www

import (
	"bytes"

	"../minify"
	"../protocol"
)

// mixCSSToHTML add given CSS file to end of HTML file and returns new html file.
func mixCSSToHTML(html, css []byte) (mixedData []byte) {
	mixedData = make([]byte, len(html), len(html)+len(css)+len("\n<style></style>"))
	copy(mixedData, html)
	mixedData = append(mixedData, "\n<style>"...)
	mixedData = append(mixedData, css...)
	mixedData = append(mixedData, "</style>"...)
	return
}

// mixCSSToJS add given CSS file to specific part of JS file and returns new js file.
func mixCSSToJS(js, css []byte) (mixedData []byte, err protocol.Error) {
	css, err = minify.CSS.MinifyBytes(css)
	if err != nil {
		return
	}

	var funcLoc = bytes.Index(js, []byte("CSS: '"))
	if funcLoc < 0 {
		var minifiedCSS = append([]byte(`CSS = '`), css...)
		mixedData = bytes.Replace(js, []byte(`CSS = '`), minifiedCSS, 1)
	} else {
		var minifiedCSS = append([]byte(`CSS: '`), css...)
		mixedData = bytes.Replace(js, []byte(`CSS: '`), minifiedCSS, 1)
	}
	return
}
