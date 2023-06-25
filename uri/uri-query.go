/* For license and copyright information please see the LEGAL file in the code repository */

package uri

import (
	"strings"

	"libgo/protocol"
	"libgo/utf8"
)

// Query maps a string key to a list of query key/value.
// It is typically used for query parameters and form values.
// Unlike in the http.Header map, the keys in a Query map
// are case-sensitive.
type Query struct {
	list map[string][]string
}

// Init initialize the Query.
// err describes the first decoding error encountered, if any.
//
// Query is expected to be a list of key=value settings separated by ampersands.
// A setting without an equals sign is interpreted as a key set to an empty
// value.
// Settings containing a non-URL-encoded semicolon are considered invalid.
func (q *Query) Init(query string) (err protocol.Error) {
	q.list = make(map[string][]string)
	err = q.FromString(query)
	return
}

// Get gets the first value associated with the given key.
// If there are no values associated with the key, Get returns
// the empty string. To access multiple values, use the map
// directly.
func (q *Query) Get(key string) string {
	if q.list == nil {
		return ""
	}
	var vs = q.list[key]
	if len(vs) == 0 {
		return ""
	}
	return vs[0]
}

// Set sets the key to value. It replaces any existing values.
func (q *Query) Set(key, value string) {
	q.list[key] = []string{value}
}

// Add adds the value to key.
// It appends to any existing values associated with key.
func (q *Query) Add(key, value string) {
	q.list[key] = append(q.list[key], value)
}

// Del deletes the values associated with key.
func (q *Query) Del(key string) {
	delete(q.list, key)
}

// Has checks whether a given key is set.
func (q *Query) Has(key string) bool {
	_, ok := q.list[key]
	return ok
}

// https://developer.mozilla.org/en-US/docs/Web/API/URLSearchParams/sort
// https://support.cloudflare.com/hc/en-us/articles/206776797-Understanding-Query-String-Sort
func (q *Query) Sort() {
}

// FromString decodes the queries from “URL encoded” form ("bar=baz&foo=quux").
//
//libgo:impl libgo/protocol.Stringer
func (q *Query) FromString(query string) (err protocol.Error) {
	for query != "" {
		var key, value string
		key, query, _ = utf8.CutByte(query, sign_And)

		if key == "" {
			continue
		}

		if strings.IndexByte(key, sign_Semicolon) > -1 {
			err = &ErrQueryBadKey
			continue
		}

		key, value, _ = utf8.CutByte(key, sign_Equal)

		var err1 protocol.Error
		key, err1 = Unescape(key, EscapeMode_QueryComponent)
		if err1 != nil {
			if err == nil {
				err = err1
			}
			continue
		}
		value, err1 = Unescape(value, EscapeMode_QueryComponent)
		if err1 != nil {
			if err == nil {
				err = err1
			}
			continue
		}
		q.list[key] = append(q.list[key], value)
	}
	return
}

// ToString encodes the queries into “URL encoded” form ("bar=baz&foo=quux").
//
//libgo:impl libgo/protocol.Stringer
func (q *Query) ToString() string {
	if q == nil {
		return ""
	}
	var buf strings.Builder
	for key, values := range q.list {
		var keyEscaped = Escape(key, EscapeMode_QueryComponent)
		for _, v := range values {
			if buf.Len() > 0 {
				buf.WriteByte('&')
			}
			buf.WriteString(keyEscaped)
			buf.WriteByte('=')
			buf.WriteString(Escape(v, EscapeMode_QueryComponent))
		}
	}
	return buf.String()
}
