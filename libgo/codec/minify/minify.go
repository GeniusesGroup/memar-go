/* For license and copyright information please see the LEGAL file in the code repository */

package minify

import (
	"regexp"

	codec_p "memar/codec/protocol"
	error_p "memar/process/error/protocol"

	"github.com/tdewolff/minify/v2"
	tcss "github.com/tdewolff/minify/v2/css"
	thtml "github.com/tdewolff/minify/v2/html"
	"github.com/tdewolff/minify/v2/js"
	"github.com/tdewolff/minify/v2/json"
	"github.com/tdewolff/minify/v2/svg"
	"github.com/tdewolff/minify/v2/xml"
)

var minifier = minify.New()

func init() {
	minifier.AddFunc("text/css", tcss.Minify)
	minifier.AddFunc("text/html", thtml.Minify)
	minifier.AddFunc("image/svg+xml", svg.Minify)
	minifier.AddFuncRegexp(regexp.MustCompile("^(application|text)/(x-)?(java|ecma)script$"), js.Minify)
	minifier.AddFuncRegexp(regexp.MustCompile("[/+]json$"), json.Minify)
	minifier.AddFuncRegexp(regexp.MustCompile("[/+]xml$"), xml.Minify)
}

// Minify replace file data with minify of them.
func Minify(data codec_p.Codec) (err error_p.Error) {
	var rawData []byte
	rawData, err = data.Marshal()
	if err != nil {
		return
	}

	var minifiedData, goErr = minifier.Bytes(data.MediaType(), rawData)
	if goErr != nil {
		return
	}
	_, err = data.Unmarshal(minifiedData)
	return
}
