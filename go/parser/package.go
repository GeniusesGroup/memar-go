/* For license and copyright information please see the LEGAL file in the code repository */

package parser

// NewRepository use to make new repository object otherwise initialize maps yourself!
func NewRepository() *Repository {
	var repo = Repository{
		Files:        make(map[string]*File),
		Imports:      make(map[string]*Import),
		Functions:    make(map[string]*Function),
		Types:        make(map[string]*Type),
		Dependencies: make(map[string]*Repository),
	}

	return &repo
}
