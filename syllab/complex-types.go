/* For license and copyright information please see LEGAL file in repository */

package syllab

/*
----encoder references in GoLang----
// Read more about this protocol : https://github.com/SabzCity/internet/blob/master/Syllab.md

fixed sized data types:
- [1]byte	: byte, int8, uint8, bool
- [2]byte	: int16, uint16
- [4]byte	: int32, uint32, float32, rune
- [8]byte	: int64, uint64, float64, int, uint, complex64
- [16]byte	: complex128
- [n]byte	: [n]byte

dynamically sized data types:
- Array		: Dynamically sized array of some any type e.g. []byte, string, []string, []uint8, ...!
*/

// Slice in go make by this struct internally! we just omit cap var due to transfer data has distinct size!
type Slice struct {
    len int
	cap int
	ptr uintptr
}

// DynamicallyArray use for dynamically sized array data type that peer don't know about length of array before get data like []string, []uint8, ... It is just to show encoder structure in better way, we never use this type in encoding or decoding process!
type DynamicallyArray struct {
	Address [4]byte
	Length  [4]byte
}

// HashTable can encode and decode HashTable by two way
// - By two Array one for keys and one for values [key0, key1, ...] & [value0, value1, ...]
// - By continuous key and value that need dedicated encoder and decoder for each need!
// By now we just support first way in this package! It is just to show encoder structure in better way, we never use this type!
// https://github.com/golang/go/blob/master/src/runtime/map.go
type HashTable struct {
	Keys   DynamicallyArray
	Values DynamicallyArray
}
