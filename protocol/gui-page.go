/* For license and copyright information please see LEGAL file in repository */

package protocol

type GUIPages interface {
	RegisterPage(page GUIPage)
	GetPageByPath(path string) (page GUIPage)

	ActivePage() (page GUIPage)
	ActivatePage(page GUIPage, state GUIPageState)
}

// GUIPage :
type GUIPage interface {
	URN() GitiURN
	Status() SoftwareStatus
	Robots() string
	Icon() []byte
	Info() GUIInformation // It is locale info

	Path() string                                            // To route page by path of HTTPURI
	AcceptedCondition(key string) (defaultValue interface{}) // HTTPURI queries

	DOM() DOM
	Template(name string) DOM
	SOM() SOM

	ActiveState() GUIPageState
	Equal(page GUIPage) bool
	// think about a page that show a user medical records, doctor need to know user birthday, so /user page must ready to reach by doctor
	//  or doctor need to know other doctors visits to know any advice from them for this user.
	RelatedPages() []GUIPage

	ActivatePage(state GUIPageState) // will activate the page to bring it front for user with requested state. Must just call by GUIPages.ActivatePage
	DeactivatePage()                 // will deactivate the page to bring it to back to allow other page being active.
}
