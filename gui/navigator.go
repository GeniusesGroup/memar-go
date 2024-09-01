/* For license and copyright information please see the LEGAL file in the code repository */

package gui

import (
	gui_p "memar/gui/protocol"
)

type Navigator struct {
	switcherPage Page
	homePage     Page
	activePage   Page
	pages        []Page
}

// Active a page like alt+tab in windows to show Pages() and their states
func (n *Navigator) SwitcherPage() {}

func (n *Navigator) HomePage() (page gui_p.Page) {
	return
}
func (n *Navigator) ActivePage() (page gui_p.Page) {
	return
}

// It must reorder by recently active page be last item in the array
func (n *Navigator) Pages() (page []gui_p.Page) {
	return
}

func (n *Navigator) ActivatePage(url string) {} // Navigate(url string)
