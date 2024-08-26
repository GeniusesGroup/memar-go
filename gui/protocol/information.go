/* For license and copyright information please see the LEGAL file in the code repository */

package gui_p

// Information store application, page, widget locale details
// https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Intl
type Information interface {
	Language() string
	Name() string
	ShortName() string
	Tagline() string
	Slogan() string
	Description() string
	Tags() []string
}
