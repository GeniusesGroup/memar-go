/* For license and copyright information please see the LEGAL file in the code repository */

package gui

// Application store the data to run a GUI application
// Such to implement phase 1 (compile go to js) to run in webview.
// https://github.com/webview/webview
type Application struct {
	Domain                  string // full domain name use for gui app like gui.geniuses.group
	Icon                    []byte
	Info                    []Information
	LocaleInfo              Information
	ContentPreferences      string
	PresentationPreferences string

	UserPreferences      UserPreferences
	DesignLanguageStyles string

	Pages
	Navigator
	History
}

// UserPreferences :
type UserPreferences struct {
	UsersState UsersState
}

// UsersState :
type UsersState struct {
	usersID      []string
	activeUserID string
}
