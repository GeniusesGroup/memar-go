/* For license and copyright information please see the LEGAL file in the code repository */

package pehrest_p

// Key MUST be in similar to urn style to have global unique key.
// Maybe a dev use `car` in both `key_value` and `text` domain,
// so made key in this manner: `kv:????`, `txt:????`, ...
type Key []byte
