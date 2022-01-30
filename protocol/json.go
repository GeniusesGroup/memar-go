/* For license and copyright information please see LEGAL file in repository */

package protocol

// JSON is the interface that must implement by any struct to be a JSON object.
// Standards by https://www.json.org/json-en.html
type JSON interface {
	// ToJSON encode the struct pointer to JSON format
	// actually payload is a byte slice buffer interface but due to prevent unnecessary memory allocation use simple []byte
	ToJSON(payload []byte) []byte
	// FromJSON decode JSON to the struct pointer
	// actually payload is a byte slice buffer interface but due to prevent unnecessary memory allocation use simple []byte
	FromJSON(payload []byte) (err Error)

	// LenAsJSON return whole calculated length of JSON encoded of the struct
	LenAsJSON() int
}
