/* For license and copyright information please see LEGAL file in repository */

package matn

import (
	etime "../earth-time"
	"../ganjine"
	lang "../language"
)

const wordSuggestionStructureID uint64 = 2096939569836997748

var wordSuggestionStructure = ganjine.DataStructure{
	ID:                2096939569836997748,
	IssueDate:         1599286551,
	ExpiryDate:        0,
	ExpireInFavorOf:   "", // Other structure name
	ExpireInFavorOfID: 0,  // Other StructureID! Handy ID or Hash of ExpireInFavorOf!
	Status:            ganjine.DataStructureStatePreAlpha,
	Structure:         WordSuggestion{},

	Name: map[lang.Language]string{
		lang.LanguageEnglish: "WordSuggestion",
	},
	Description: map[lang.Language]string{
		lang.LanguageEnglish: " store the hash index data by 32byte key and any 32byte value",
	},
	TAGS: []string{
		"",
	},
}

// WordSuggestion is standard structure to store any hash byte index!!
// It is simple secondary index e.g. hash("user@email.com")
type WordSuggestion struct {
	/* Common header data */
	RecordID          [32]byte
	RecordStructureID uint64
	RecordSize        uint64
	WriteTime         etime.Time
	OwnerAppID        [32]byte

	/* Unique data */
	id       [32]byte
	start    int
	end      int
	surface  string
	class    uint8
	features []string
}
