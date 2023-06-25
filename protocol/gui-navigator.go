/* For license and copyright information please see the LEGAL file in the code repository */

package protocol

type GUI_Navigator interface {
	// Active a page like alt+tab in windows to show Pages() and their states
	SwitcherPage()

	HomePage() (page GUI_Page)
	ActivePage() (page GUI_Page)
	// It must reorder by recently active page be last item in the array
	ActivePages() (pages []GUI_Page)

	ActivatePage(url string) // Navigate(url string)
}
