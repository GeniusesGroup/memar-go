/* For license and copyright information please see LEGAL file in repository */

package chaparkhane

// Assets use to store app needed data from repo like html, css, js, ...
type Assets struct {
	Name         string
	Files        map[string]*AssetsFile // Name
	Dependencies map[string]*Assets     // Name
}

// AssetsFile :
type AssetsFile struct {
	Name string
	Data []byte
}

// NewAssets will make new assets object
func NewAssets() *Assets {
	return &Assets{
		Files:        make(map[string]*AssetsFile),
		Dependencies: make(map[string]*Assets),
	}
}
