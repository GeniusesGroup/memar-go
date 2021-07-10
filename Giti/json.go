/* For license and copyright information please see LEGAL file in repository */

package giti

// JSON is the interface that must implement by any struct to be a JSON object!
// Standards by https://www.json.org/json-en.html
type JSON interface {
	// ToJSON encode the struct pointer to JSON format
	ToJSON(buf Buffer)
	// FromJSON decode JSON to the struct pointer!
	FromJSON(buf Buffer) (err Error)

	// LenAsJSON return whole calculated length of JSON encoded of the struct
	LenAsJSON() int
}
