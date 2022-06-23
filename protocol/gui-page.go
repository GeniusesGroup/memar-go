/* For license and copyright information please see LEGAL file in repository */

package protocol

type GUIPages interface {
	RegisterPage(page GUIPage)
	GetPageByPath(path string) (page GUIPage)
	Pages() (pages []GUIPage)
}

// GUIPage indicate what is a GUI page.
type GUIPage interface {
	Robots() string
	Icon() Image
	Info() GUIInformation // It is locale info
	// think about a page that show a user medical records, doctor need to know user birthday, so /user page must ready to reach by doctor
	// or doctor need to know other doctors visits to know any advice from them for this user.
	RelatedPages() []GUIPage

	Path() string                                    // To route page by path of HTTP-URI
	AcceptedCondition(key string) (defaultValue any) // HTTP-URI queries

	ActiveState() GUIPageState
	ActiveStates() []GUIPageState

	// Below methods are custom methods that must implement in each page not gui library.

	// CreateState build the page in the requested state or reuse old states with SafeToSilentClose()
	CreateState(url string) GUIPageState

	// it is raw version of the page DOM. suggest to write HTML in a dedicated html file and have compile-time parser.
	// - Due to multi page state mechanism and separate concern(content vs logic), don't support inline event handlers in HTML files
	// dom() DOM
	// it is raw version of the page SOM
	// som() SOM
	// it is raw version of the page templates DOM e.g. products-template-card.html
	// template(name string) DOM

	MediaType
}
