/* For license and copyright information please see the LEGAL file in the code repository */

package gui

import (
	"libgo/protocol"
)

// Pages implements protocol.GUI_Pages interface
type Pages struct {
	poolByTimeAdded []protocol.GUI_Page
	poolByPath      map[string]protocol.GUI_Page
}

func (p *Pages) RegisterPage(page protocol.GUI_Page) {
	p.poolByTimeAdded = append(p.poolByTimeAdded, page)
	p.poolByPath[page.Path()] = page
}
func (p *Pages) GetPageByPath(path string) (page protocol.GUI_Page) {
	return p.poolByPath[path]
}
func (p *Pages) Pages() (pages []protocol.GUI_Page) {
	return p.poolByTimeAdded
}
