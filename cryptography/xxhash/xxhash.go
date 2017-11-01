package xxhash

import "github.com/OneOfOne/xxhash"

// SumString64 : Return xxhash64 sum of "str".
func SumString64(str string) uint64 {

	return xxhash.ChecksumString64(str)
}
