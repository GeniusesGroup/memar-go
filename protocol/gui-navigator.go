/* For license and copyright information please see LEGAL file in repository */

package protocol

type GUINavigator interface {
	// Active a page like alt+tab in windows to show Pages() and their states
	SwitcherPage()

	HomePage() (page GUIPage)
	ActivePage() (page GUIPage)
	// It must reorder by recently active page be last item in the array
	ActivePages() (pages []GUIPage)

	ActivatePage(url string) // Navigate(url string)
}
