/* For license and copyright information please see the LEGAL file in the code repository */

package detail

type Quiddity struct {
	name         string
	abbreviation string
	aliases      []string
}

//libgo:impl libgo/protocol.Quiddity
func (q *Quiddity) Name() string         { return q.name }
func (q *Quiddity) Abbreviation() string { return q.abbreviation }
func (q *Quiddity) Aliases() []string    { return q.aliases }

func (q Quiddity) SetName(v string) Quiddity         { q.name = v; return q }
func (q Quiddity) SetAbbreviation(v string) Quiddity { q.abbreviation = v; return q }
func (q Quiddity) SetAliases(v []string) Quiddity    { q.aliases = v; return q }
