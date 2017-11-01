package md5

import (
	"crypto/md5"
	"encoding/hex"
)

// SumString : Return md5 sum of "str".
func SumString(str string) string {

	sum := md5.Sum([]byte(str))

	return hex.EncodeToString(sum[:])
}
