/* For license and copyright information please see the LEGAL file in the code repository */

package gui

import "libgo/protocol"

type Navigator struct {
	switcherPage Page
	homePage     Page
	activePage   Page
	pages        []Page
}

// Active a page like alt+tab in windows to show Pages() and their states
func (n *Navigator) SwitcherPage() {}

func (n *Navigator) HomePage() (page protocol.GUI_Page) {
	return
}
func (n *Navigator) ActivePage() (page protocol.GUI_Page) {
	return
}

// It must reorder by recently active page be last item in the array
func (n *Navigator) Pages() (page []protocol.GUI_Page) {
	return
}

func (n *Navigator) ActivatePage(url string) {} // Navigate(url string)
