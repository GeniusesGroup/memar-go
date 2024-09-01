/* For license and copyright information please see the LEGAL file in the code repository */

package gui

import gui_p "memar/gui/protocol"

// Pages implements gui_p.Pages interface
type Pages struct {
	poolByTimeAdded []gui_p.Page
	poolByPath      map[string]gui_p.Page
}

func (p *Pages) RegisterPage(page gui_p.Page) {
	p.poolByTimeAdded = append(p.poolByTimeAdded, page)
	p.poolByPath[page.Path()] = page
}
func (p *Pages) GetPageByPath(path string) (page gui_p.Page) {
	return p.poolByPath[path]
}
func (p *Pages) Pages() (pages []gui_p.Page) {
	return p.poolByTimeAdded
}
