/* For license and copyright information please see LEGAL file in repository */

package minify

import (
	"regexp"

	"../protocol"

	thtml "github.com/tdewolff/minify/html"
	"github.com/tdewolff/parse/buffer"
)

var HTML html

var htmlMinifier = regexp.MustCompile(`\r?\n|(<!--.*?-->)|(<!--[\w\W\n\s]+?-->)`)

type html struct {
	thtml.Minifier
}

// Minify replace file data with minify of them.
func (html *html) Minify(data protocol.Codec) (err protocol.Error) {
	var rawData = data.Marshal()
	var minifiedData []byte
	minifiedData, err = html.MinifyBytes(rawData)
	if err != nil {
		return
	}
	err = data.Unmarshal(minifiedData)
	return
}

// MinifyBytes replace file data with minify of them.
func (html *html) MinifyBytes(data []byte) (minifiedData []byte, err protocol.Error) {
	// TODO::: have problem minify HTML >> https://github.com/tdewolff/minify/issues/318
	if true {
		minifiedData = htmlMinifier.ReplaceAll(data, []byte{})
		return
	}
	var minifiedWriter = buffer.NewWriter(make([]byte, 0, len(data)))
	var goErr = html.Minifier.Minify(minifier, minifiedWriter, buffer.NewReader(data), nil)
	if goErr != nil {
		// err =
		return
	}
	minifiedData = minifiedWriter.Bytes()
	return
}
