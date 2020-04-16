/* For license and copyright information please see LEGAL file in repository */

package lang

const (
	// EnglishScript language code [4]byte{0,128,128,128}
	EnglishScript uint32 = 0
	// PersianScript language code [4]byte{0,129,128,128}
	PersianScript uint32 = 1
)

// Direction of scripts
const (
	LeftToRight uint8 = iota
	RightToLeft
	UpToDown
	DownToUp
)

// ScriptsDetail use to store other detail for a language
type ScriptsDetail struct {
	Code          uint32
	Description   []string
	Dir           uint8
	CharectersIDs [][4]byte
}

// ScriptsDetails store all available script in a way to find by given ScriptCode
var ScriptsDetails = map[uint32]ScriptsDetail{
	EnglishScript: ScriptsDetail{
		Code:        EnglishScript,
		Description: []string{"English", "English"},
		Dir:         LeftToRight,
	},
	PersianScript: ScriptsDetail{
		Code:        PersianScript,
		Description: []string{"Persian", "پارسی"},
		Dir:         RightToLeft,
	},
}
