/* For license and copyright information please see the LEGAL file in the code repository */

package harf

// Script indicate
type Script uint64

// Scripts codes
const (
	ScriptControl Script = 128 + iota // 0b10000000

	ScriptMathematicsNumeral // https://en.wikipedia.org/wiki/Mathematical_operators_and_symbols_in_Unicode
	ScriptMathematics1
	ScriptMathematics2
	ScriptMathematics3
	ScriptMathematics4
	ScriptMathematics5
	ScriptMathematics6
	ScriptMathematics7
	ScriptMathematics8
	ScriptMathematics9
	ScriptMathematics10
	ScriptMathematics11
	ScriptMathematics12
	ScriptMathematics13
	ScriptMathematics14
	ScriptMathematics15
	ScriptMathematics16
	ScriptMathematics17
	ScriptMathematics18
	ScriptMathematics19
	ScriptMathematics20
	ScriptMathematics21
	ScriptMathematics22
	ScriptMathematics23
	ScriptMathematics24
	ScriptMathematicsArrow1
	ScriptMathematicsArrow2
	ScriptMathematicsArrow3
	ScriptMathematicsArrow4
	ScriptMathematicsArrow5
	ScriptMathematicsArrow6
	ScriptMathematicsArrow7
	ScriptMathematicsArrow8

	ScriptMusical1 // https://en.wikipedia.org/wiki/Musical_Symbols_(Unicode_block)
	ScriptMusical2
	ScriptMusical3
	ScriptMusical4
	ScriptMusical5
	ScriptMusical6
	ScriptMusical7
	ScriptMusical8

	ScriptPersian
	ScriptLatin
	ScriptCyrillic // https://en.wikipedia.org/wiki/Cyrillic_script

	ScriptEmoji1 Script = 32896 + iota // 0b1000000010000000
	// ...

	ScriptNotKnownYet1 Script = 8421504 + iota // 0b100000001000000010000000
	// ...

	ScriptNotKnownYet2 Script = 2155905152 + iota // 0b10000000100000001000000010000000
	// ...

	ScriptNotKnownYet3 Script = 551911719040 + iota // 0b1000000010000000100000001000000010000000
	// ...

	ScriptNotKnownYet4 Script = 141289400074368 + iota // 0b100000001000000010000000100000001000000010000000
	// ...

	ScriptNotKnownYet5 Script = 36170086419038336 + iota // 0b10000000100000001000000010000000100000001000000010000000
	// ...
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
	Code          Script
	Description   []string
	Dir           uint8
	CharactersIDs [][4]byte
}

// ScriptsDetails store all available script in a way to find by given ScriptCode
var ScriptsDetails = map[Script]ScriptsDetail{
	ScriptLatin: {
		Code:        ScriptLatin,
		Description: []string{"Latin", "Roman", "English"},
		Dir:         LeftToRight,
	},
	ScriptPersian: {
		Code:        ScriptPersian,
		Description: []string{"Persian", "پارسی"},
		Dir:         RightToLeft,
	},
}
