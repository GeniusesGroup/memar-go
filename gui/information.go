/* For license and copyright information please see the LEGAL file in the code repository */

package gui

// Information :
type Information struct {
	language    string
	name        string
	shortName   string
	tagline     string
	slogan      string
	description string
	tags        []string
}

func (i *Information) Language() string    { return i.language }
func (i *Information) Name() string        { return i.name }
func (i *Information) ShortName() string   { return i.shortName }
func (i *Information) Tagline() string     { return i.tagline }
func (i *Information) Slogan() string      { return i.slogan }
func (i *Information) Description() string { return i.description }
func (i *Information) Tags() []string      { return i.tags }
