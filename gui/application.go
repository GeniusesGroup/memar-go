/* For license and copyright information please see LEGAL file in repository */

package gui

func init() {
	
}

// Application store the data to run a GUI application
type Application struct {
	Domain                  string // full domain name use for gui app like gui.sabz.city
	Icon                    []byte
	Info                    []Information
	LocaleInfo              Information
	ContentPreferences      string
	PresentationPreferences string
	Pages                   Pages
	UserPreferences         UserPreferences
	DesignLanguageStyles    string
}

// Pages :
type Pages struct {
	List Page
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
