/* For license and copyright information please see LEGAL file in repository */

package protocol

type CompressTypes interface {
	RegisterCompressType(ct CompressType)
	GetCompressTypeByID(urnID uint64) CompressType
	GetCompressTypeByURN(urnURI string) CompressType
	GetCompressTypeByFileExtension(ex string) CompressType
	GetCompressTypeByContentEncoding(ce string) CompressType
}

// CompressType is standard shape of any compress coding type
// https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Content-Encoding
// https://en.wikipedia.org/wiki/HTTP_compression
type CompressType interface {
	URN() GitiURN
	ContentEncoding() string
	Extension() string
	Compression(raw Codec, compressLevel CompressLevel) (compress Codec)
	Decompression(compress Codec) (raw Codec)
}

type CompressLevel int

// Compress Levels
const (
	NoCompression        CompressLevel = 0
	BestSpeedCompression CompressLevel = 1
	BestCompression      CompressLevel = 9
	DefaultCompression   CompressLevel = -1

	// HuffmanOnly disables Lempel-Ziv match searching and only performs Huffman
	// entropy encoding. This mode is useful in compressing data that has
	// already been compressed with an LZ style algorithm (e.g. Snappy or LZ4)
	// that lacks an entropy encoder. Compression gains are achieved when
	// certain bytes in the input stream occur more frequently than others.
	//
	// Note that HuffmanOnly produces a compressed output that is
	// RFC 1951 compliant. That is, any valid DEFLATE decompressor will
	// continue to be able to decompress this output.
	HuffmanOnlyCompression CompressLevel = -2
)
