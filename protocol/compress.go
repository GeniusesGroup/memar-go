/* For license and copyright information please see the LEGAL file in the code repository */

package protocol

// CompressType is standard shape of any compress coding type
// https://en.wikipedia.org/wiki/HTTP_compression
type CompressType /* [T Codec] */ interface {
	MediaType

	// https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Content-Encoding
	// https://www.iana.org/assignments/http-parameters/http-parameters.xml#content-coding
	ContentEncoding() string

	Compress(raw Codec, options CompressOptions) (compressed Codec, err Error)

	Decompress(compressed Codec) (raw Codec, err Error)
}

type CompressOptions struct {
	CompressLevel CompressLevel
}

type CompressLevel int

// Compress Levels
const (
	CompressLevel_NoCompression   CompressLevel = 0
	CompressLevel_BestSpeed       CompressLevel = 1
	CompressLevel_BestCompression CompressLevel = 9
	CompressLevel_Default         CompressLevel = -1

	// HuffmanOnly disables Lempel-Ziv match searching and only performs Huffman
	// entropy encoding. This mode is useful in compressing data that has
	// already been compressed with an LZ style algorithm (e.g. Snappy or LZ4)
	// that lacks an entropy encoder. Compression gains are achieved when
	// certain bytes in the input stream occur more frequently than others.
	//
	// Note that HuffmanOnly produces a compressed output that is
	// RFC 1951 compliant. That is, any valid DEFLATE decompressor will
	// continue to be able to decompress this output.
	CompressLevel_HuffmanOnly CompressLevel = -2
)
