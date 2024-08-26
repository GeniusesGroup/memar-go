/* For license and copyright information please see the LEGAL file in the code repository */

package gui_p

type Navigator interface {
	// Active a page like alt+tab in windows to show Pages() and their states
	SwitcherPage()

	HomePage() (page Page)
	ActivePage() (page Page)
	// It must reorder by recently active page be last item in the array
	ActivePages() (pages []Page)

	ActivatePage(url string) // Navigate(url string)
}
