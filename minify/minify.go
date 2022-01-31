/* For license and copyright information please see LEGAL file in repository */

package minify

import (
	"regexp"

	"../protocol"

	"github.com/tdewolff/minify"
	tcss "github.com/tdewolff/minify/css"
	thtml "github.com/tdewolff/minify/html"
	"github.com/tdewolff/minify/js"
	"github.com/tdewolff/minify/json"
	"github.com/tdewolff/minify/svg"
	"github.com/tdewolff/minify/xml"
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
func Minify(data protocol.Codec) (err protocol.Error) {
	var minifiedData, goErr = minifier.Bytes(data.MediaType().MediaType(), data.Marshal())
	if goErr != nil {
		return
	}
	err = data.Unmarshal(minifiedData)
	return
}
