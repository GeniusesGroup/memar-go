/* For license and copyright information please see the LEGAL file in the code repository */

package http

import (
	"slices"

	error_p "memar/process/error/protocol"
	string_p "memar/codec/string/protocol"
)

// Header is represent HTTP header structure.
// Exported to let consumers use other methods that memar/net/http/protocol.Header
type Header[STR string_p.String] struct {
	lines []header_KV[STR]
}

//memar:impl memar/computer/capsule/protocol.LifeCycle
func (h *Header[STR]) Init() (err error_p.Error) {
	h.lines = make([]header_KV[STR], headerInitLen)
	return
}
func (h *Header[STR]) Reinit() (err error_p.Error) {
	clear(h.lines)
	// Below logic not work due to GC can't free strings.
	// h.lines = h.lines[:0]
	return
}
func (h *Header[STR]) Deinit() (err error_p.Error) {
	return
}

// Get returns the first value associated with the given key.
// Both given key and header key SHOULD already be in CanonicalHeaderKey form or in same shape.
//
//memar:impl memar/net/http/protocol.Header
func (h *Header[STR]) Header_Get(key string_p.String) (value string_p.String) {
	value, _ = h.Header_Find(0, key.(STR))
	return
}

// Add append the key, value pair to the end of the header.
// Key SHOULD already be in CanonicalHeaderKey form.
//
//memar:impl memar/net/http/protocol.Header
func (h *Header[STR]) Header_Add(key, value string_p.String) {
	var kv = header_KV[STR]{key.(STR), value.(STR)}
	h.lines = append(h.lines, kv)
}

// Set replace given value in given key, or Add if given key not exist.
// Key SHOULD already be in CanonicalHeaderKey form.
//
//memar:impl memar/net/http/protocol.Header
func (h *Header[STR]) Header_Set(key, value string_p.String) {
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
//memar:impl memar/net/http/protocol.Header
func (h *Header[STR]) Header_Del(key string_p.String) {
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

func (h *Header[STR]) Header_All() []header_KV[STR] { return h.lines }

// Exclude eliminate headers by given keys.
func (h *Header[STR]) Exclude(exclude ...STR) {
	for _, key := range exclude {
		h.Header_Del(key)
	}
}

// Header_Find returns the first value associated with the given key.
func (h *Header[STR]) Header_Find(startIndex int, key STR) (value STR, index int) {
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

type header_KV[STR string_p.String] struct {
	key   STR
	value STR
}

func (kv *header_KV[STR]) Key() STR   { return kv.key }
func (kv *header_KV[STR]) Value() STR { return kv.value }

type Key[STR string_p.String] struct {
	key STR
}

func (k *Key[STR]) Set(key STR) {
	// if key.CharacterEncoding() != ascii.CharacterEncoding {
	// 	compiler.Log.Fatal("Linter MUST notify developers not call this method with string other than ASCII")
	// }
	k.key = key
	k.key.CopyFrom(key)
	k.key.CloneFrom(key)
}

type Value[STR string_p.String] string_p.String
