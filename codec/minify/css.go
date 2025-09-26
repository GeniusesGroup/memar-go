/* For license and copyright information please see the LEGAL file in the code repository */

package minify

import (
	"memar/protocol"

	tcss "github.com/tdewolff/minify/css"
	"github.com/tdewolff/parse/buffer"
)

var CSS css

type css struct {
	tcss.Minifier
}

// Minify replace file data with minify of them.
func (css *css) Minify(data protocol.Codec) (err protocol.Error) {
	var rawData []byte
	rawData, err = data.Marshal()
	if err != nil {
		return
	}

	var minifiedData []byte
	minifiedData, err = css.MinifyBytes(rawData)
	if err != nil {
		return
	}

	_, err = data.Unmarshal(minifiedData)
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
