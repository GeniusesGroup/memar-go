/* For license and copyright information please see LEGAL file in repository */

package validators

import (
	"regexp"
)

const (
	// GenderPattern : Regular expression of gender.
	GenderPattern = `^(Male|Female|Gay|Lesbian|Bisexual|Asexual|Other)$`
)

// Gender : Validate a gender code.
func Gender(gender string) bool {
	match, _ := regexp.MatchString(GenderPattern, gender)
	return match
}
