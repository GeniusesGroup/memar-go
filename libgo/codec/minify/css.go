/* For license and copyright information please see the LEGAL file in the code repository */

package minify

import (
	codec_p "memar/codec/protocol"
	error_p "memar/process/error/protocol"

	tcss "github.com/tdewolff/minify/v2/css"
	"github.com/tdewolff/parse/v2/buffer"
)

var CSS css

type css struct {
	tcss.Minifier
}

// Minify replace file data with minify of them.
func (css *css) Minify(data codec_p.Codec) (err error_p.Error) {
	var rawData []byte
	rawData, err = data.Encode()
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
func (css *css) MinifyBytes(data []byte) (minifiedData []byte, err error_p.Error) {
	var minifiedWriter = buffer.NewWriter(make([]byte, 0, len(data)))
	var goErr = css.Minifier.Minify(minifier, minifiedWriter, buffer.NewReader(data), nil)
	if goErr != nil {
		// err =
		return
	}
	minifiedData = minifiedWriter.Bytes()
	return
}
