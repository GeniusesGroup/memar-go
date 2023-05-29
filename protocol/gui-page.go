/* For license and copyright information please see the LEGAL file in the code repository */

package protocol

type GUIPages interface {
	RegisterPage(page GUIPage)
	GetPageByPath(path string) (page GUIPage)
	Pages() (pages []GUIPage)
}

// GUIPage indicate what is a GUI page.
type GUIPage interface {
	// "all", "noindex", "nofollow", "none", "noarchive", "nosnippet", "notranslate", "noimageindex", "unavailable_after: [RFC-850 date/time]"
	Robots() string
	Icon() Image
	Info() GUIInformation // It is locale info
	// think about a page that show a user medical records, doctor need to know user birthday, so /user page must ready to reach by doctor
	// or doctor need to know other doctors visits to know any advice from them for this user.
	RelatedPages() []GUIPage

	Path() string                                    // To route page by path of HTTP-URI
	// /product?id=1&title=book
	AcceptedCondition(key string) (defaultValue any) // HTTP-URI queries

	ActiveState() GUI_PageState
	ActiveStates() []GUI_PageState

	// Below methods are custom methods that must implement in each page not gui library.

	// CreateState build the page in the requested state or reuse old states with SafeToSilentClose()
	CreateState(uri URI) GUI_PageState

	// it is raw version of the page DOM. suggest to write HTML in a dedicated html file and have compile-time parser.
	// - Due to multi page state mechanism and separate concern(content vs logic), don't support inline event handlers in HTML files
	// dom() DOM
	// it is raw version of the page SOM
	// som() SOM
	// it is raw version of the page templates DOM e.g. products-template-card.html
	// template(name string) DOM

	MediaType
}
