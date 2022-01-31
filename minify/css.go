/* For license and copyright information please see LEGAL file in repository */

package minify

import (
	"../protocol"

	tcss "github.com/tdewolff/minify/css"
	"github.com/tdewolff/parse/buffer"
)

var CSS css

type css struct {
	tcss.Minifier
}

// Minify replace file data with minify of them.
func (css *css) Minify(data protocol.Codec) (err protocol.Error) {
	var rawData = data.Marshal()
	var minifiedData []byte
	minifiedData, err = css.MinifyBytes(rawData)
	if err != nil {
		return
	}
	err = data.Unmarshal(minifiedData)
	return
}

// MinifyBytes replace file data with minify of them.
func (css *css) MinifyBytes(data []byte) (minifiedData []byte, err protocol.Error) {
	var minifiedWriter = buffer.NewWriter(make([]byte, 0, len(data)))
	var goErr = css.Minifier.Minify(minifier, minifiedWriter, buffer.NewReader(data), nil)
	if goErr != nil {
		// err = 
		return
	}
	minifiedData = minifiedWriter.Bytes()
	return
}
