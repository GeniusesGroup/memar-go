/* For license and copyright information please see LEGAL file in repository */

package protocol

// Minify replace given data with minify of them if possible.
type Minifier interface {
	Minify(data Codec) (err Error)
	MinifyBytes(data []byte) (minified []byte, err Error)
}
