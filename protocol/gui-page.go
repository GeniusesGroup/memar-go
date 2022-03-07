/* For license and copyright information please see LEGAL file in repository */

package protocol

type GUIPages interface {
	RegisterPage(page GUIPage)
	GetPageByPath(path string) (page GUIPage)
}

// GUIPage :
type GUIPage interface {
	MediaType() MediaType
	Status() SoftwareStatus
	Robots() string
	Icon() []byte
	Info() GUIInformation // It is locale info
	// think about a page that show a user medical records, doctor need to know user birthday, so /user page must ready to reach by doctor
	// or doctor need to know other doctors visits to know any advice from them for this user.
	RelatedPages() []GUIPage

	Path() string                                            // To route page by path of HTTPURI
	AcceptedCondition(key string) (defaultValue interface{}) // HTTPURI queries

	ActiveState() GUIPageState
	ActiveStates() []GUIPageState

	// CreateState build the page in the requested state or reuse old states with SafeToSilentClose()
	CreateState(url string) GUIPageState

	// it is raw version of the page DOM. suggest to write HTML in a dedicated html file and have compile-time parser.
	// - Due to multi page state mechanism and seprate concern(content vs logic), don't support inline event handlers in HTML files
	// dom() DOM
	// som() SOM                 // it is raw version of the page SOM
	// template(name string) DOM // it is raw version of the page templates DOM
}

// GUIPageState :
type GUIPageState interface {
	Page() GUIPage
	URL() string
	Title() string
	Description() string
	Conditions() map[string]string
	Fragment() string

	CreateDate() Time
	LastActiveDate() Time
	EndDate() Time // When Close() called

	// DOM as Document-Object-Model will render to user screen
	// Same as Document interface of the Web declare in browsers in many ways but split not related concern to other interfaces
	// like Document.URI that not belong to this interface and relate to GUIPageState because we develope multi page app not web page
	DOM() GUIElement
	SOM() SOM
	Thumbnail() []byte // The picture of the page in current state, use in social linking and SwitcherPage()

	// Called before this state wants to add to the render tree (brings to front)
	Activate()
	// Called before this state wants to remove from the render tree (brings to back)
	// - approveLeave let the caller know user of the GUI app let page in this state bring to back.
	// - hadActiveDialog or hadActiveOverlay help navigator works in better way.
	// e.g. for some keyboard event like back buttom in android OS to close dialog not pop previous page state to front
	Deactivate() (approveLeave, hadActiveDialog bool)
	Refresh()

	SafeToSilentClose() bool
	//  Remove DOM and SOM and archive it.
	Close()

	// DynamicElements struct {}
}
