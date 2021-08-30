/* For license and copyright information please see LEGAL file in repository */

package protocol

type GUIPages interface {
	RegisterPage(page GUIPage)
	GetPageByURNName(urnName string) (page GUIPage)

	ActivePage() (page GUIPage)
	ActivatePage(page GUIPage, state GUIPageState)
}

// GUIPage :
type GUIPage interface {
	URN() GitiURN
	Icon() []byte
	Info() GUIInformation // It is locale info

	ActiveState() GUIPageState
	Equal(page GUIPage) bool
}
