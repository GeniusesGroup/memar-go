/* For license and copyright information please see LEGAL file in repository */

package lang

// Compress use to compress a string array!
func Compress(s []byte) ([]byte, error) { return nil, nil }

// UnCompress use to un-compress a string array!
func UnCompress(s []byte) ([]byte, error) { return nil, nil }

// Validate use to validate a string array for any error!
// We don't offer any fix proccess! Any suggestion can have security concern!
func Validate(s []byte) error { return nil }

// ValidateDeep use to validate a string array for any error deeply!
// We don't offer any fix proccess! Any suggestion can have security concern!
func ValidateDeep(s []byte) error { return nil }

// EncodeToUTF8 use to encode(convert) this package structure to UTF-8 structure!
func EncodeToUTF8(s []byte) ([]byte, error) { return nil, nil }

// DecodeFromUTF8 use to decode(convert) UTF-8 structure to this package structure!
func DecodeFromUTF8(s []byte) ([]byte, error) { return nil, nil }

// DetectScripts use to detect scripts IDs in a string
// It mostly use for compression and un-compression!
func DetectScripts(s []byte) []uint32 { return nil }

// DecodeCompressCharecter use to decode first charecter and its size in a valid compress string!
func DecodeCompressCharecter(s []byte) (ch [4]byte, size int) {
	var s0 byte = s[0]
	var s1 byte = s[1]
	if s1 < 128 {
		return [4]byte{s0, 128, 128, 128}, 1
	}
	var s2 byte = s[2]
	if s2 < 128 {
		return [4]byte{s0, s1, 128, 128}, 2
	}
	var s3 byte = s[3]
	if s3 < 128 {
		return [4]byte{s0, s1, s2, 128}, 3
	}
	return [4]byte{s0, s1, s2, s3}, 4
}
