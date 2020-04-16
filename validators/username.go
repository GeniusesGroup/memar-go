/* For license and copyright information please see LEGAL file in repository */

package validators

import "regexp"

const (
	// UsernamePattern : Regular expression of username.
	UsernamePattern = `^[a-zA-](([\._\-][a-z0-9])|[a-z0-9]){0,62}$`
)

// ValidateUsername use to validate a username string.
func ValidateUsername(username string) bool {
	match, _ := regexp.MatchString(UsernamePattern, username)
	return match
}
