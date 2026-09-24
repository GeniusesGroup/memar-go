/* For license and copyright information please see the LEGAL file in the code repository */

package http

import (
	"slices"

	error_p "memar/process/error/protocol"
)

// Header is represent HTTP header structure.
// Exported to let consumers use other methods that memar/libgo/net/http/protocol.Header
type Header struct {
	lines []header_KV
}

//memar:impl memar/computer/capsule/protocol.LifeCycle
func (h *Header) Init() (err error_p.Error) {
	h.lines = make([]header_KV, headerInitLen)
	return
}
func (h *Header) Reinit() (err error_p.Error) {
	clear(h.lines)
	// Below logic not work due to GC can't free strings.
	// h.lines = h.lines[:0]
	return
}
func (h *Header) Deinit() (err error_p.Error) {
	return
}

// Get returns the first value associated with the given key.
// Both given key and header key SHOULD already be in CanonicalHeaderKey form or in same shape.
//
//memar:impl memar/libgo/net/http/protocol.Header
func (h *Header) Header_Get(key string) (value string) {
	value, _ = h.Header_Find(0, key)
	return
}

// Add append the key, value pair to the end of the header.
// Key SHOULD already be in CanonicalHeaderKey form.
//
//memar:impl memar/libgo/net/http/protocol.Header
func (h *Header) Header_Add(key, value string) {
	var kv = header_KV{key, value}
	h.lines = append(h.lines, kv)
}

// Set replace given value in given key, or Add if given key not exist.
// Key SHOULD already be in CanonicalHeaderKey form.
//
//memar:impl memar/libgo/net/http/protocol.Header
func (h *Header) Header_Set(key, value string) {
	var ln = len(h.lines)
	var set bool
	for i := 0; i < ln; i++ {
		var hPair = h.lines[i]
		var hKey = hPair.Key()
		if hKey.Equivalence(key) {
			if !set {
				hPair.value.CopyFrom(value)
				set = true
			} else {
				h.lines = slices.Delete(h.lines, i, i)
				ln--
				i--
			}
		}
	}
	if !set {
		h.Header_Add(key, value)
	}
}

// Del deletes the values associated with key.
// Key SHOULD already be in CanonicalHeaderKey form.
//
//memar:impl memar/libgo/net/http/protocol.Header
func (h *Header) Header_Del(key string) {
	var ln = len(h.lines)
	for i := 0; i < ln; i++ {
		var hPair = h.lines[i]
		var hKey = hPair.Key()
		if hKey.Equivalence(key) {
			h.lines = slices.Delete(h.lines, i, i)
			ln--
			i--
		}
	}
}

func (h *Header) Header_All() []header_KV { return h.lines }

// Exclude eliminate headers by given keys.
func (h *Header) Exclude(exclude ...string) {
	for _, key := range exclude {
		h.Header_Del(key)
	}
}

// Header_Find returns the first value associated with the given key.
func (h *Header) Header_Find(startIndex int, key string) (value string, index int) {
	var ln = len(h.lines)
	for i := startIndex; i < ln; i++ {
		var hPair = h.lines[i]
		var hKey = hPair.Key()
		if hKey.Equivalence(key) {
			return hPair.Value(), i
		}
	}
	return
}

type header_KV struct {
	key   string
	value string
}

func (kv *header_KV) Key() string   { return kv.key }
func (kv *header_KV) Value() string { return kv.value }

type Key struct {
	key string
}

func (k *Key) Set(key string) {
	// if key.CharacterEncoding() != ascii.CharacterEncoding {
	// 	compiler.Log.Fatal("Linter MUST notify developers not call this method with string other than ASCII")
	// }
	k.key = key
	k.key.CopyFrom(key)
	k.key.CloneFrom(key)
}

type Value string
