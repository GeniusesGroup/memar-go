/* For license and copyright information please see the LEGAL file in the code repository */

package gui_p

type Pages interface {
	RegisterPage(page Page)
	GetPageByPath(path string) (page Page)
	Pages() (pages []Page)
}

// Page indicate what is a GUI page.
type Page interface {
	// "all", "noindex", "nofollow", "none", "noarchive", "nosnippet", "notranslate", "noimageindex", "unavailable_after: [RFC-850 date/time]"
	Robots() string
	Icon() Image
	Info() Information // It is locale info
	// think about a page that show a user medical records, doctor need to know user birthday, so /user page must ready to reach by doctor
	// or doctor need to know other doctors visits to know any advice from them for this user.
	RelatedPages() []Page

	Path() string // To route page by path of HTTP-URI
	// /product?id=1&title=book
	AcceptedCondition(key string) (defaultValue any) // HTTP-URI queries

	ActiveState() PageState
	ActiveStates() []PageState

	// Below methods are custom methods that must implement in each page not gui library.

	// CreateState build the page in the requested state or reuse old states with SafeToSilentClose()
	// CreateState(uri URI) PageState

	// it is raw version of the page DOM. suggest to write HTML in a dedicated html file and have compile-time parser.
	// - Due to multi page state mechanism and separate concern(content vs logic), don't support inline event handlers in HTML files
	// dom() DOM
	// it is raw version of the page SOM
	// som() SOM
	// it is raw version of the page templates DOM e.g. products-template-card.html
	// template(name string) DOM

	MediaType
}

type Page_DefaultConditions interface {
	Editable() bool
	// MarketingUTM
}
