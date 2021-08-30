/* For license and copyright information please see LEGAL file in repository */

package protocol

// GUIInformation store application, page, widget locale details
type GUIInformation interface {
	Language() string
	Name() string
	ShortName() string
	Tagline() string
	Slogan() string
	Description() string
	Tags() []string
}
