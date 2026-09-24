/* For license and copyright information please see the LEGAL file in the code repository */

package minify

import (
	"regexp"

	codec_p "memar/codec/protocol"
	error_p "memar/process/error/protocol"

	thtml "github.com/tdewolff/minify/v2/html"
	"github.com/tdewolff/parse/v2/buffer"
)

var HTML html

var htmlMinifier = regexp.MustCompile(`\r?\n|(<!--.*?-->)|(<!--[\w\W\n\s]+?-->)`)

type html struct {
	thtml.Minifier
}

// Minify replace file data with minify of them.
func (html *html) Minify(data codec_p.Codec) (err error_p.Error) {
	var rawData []byte
	rawData, err = data.Marshal()
	if err != nil {
		return
	}

	var minifiedData []byte
	minifiedData, err = html.MinifyBytes(rawData)
	if err != nil {
		return
	}

	_, err = data.Unmarshal(minifiedData)
	return
}

// MinifyBytes replace file data with minify of them.
func (html *html) MinifyBytes(data []byte) (minifiedData []byte, err error_p.Error) {
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
