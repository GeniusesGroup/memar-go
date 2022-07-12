/* For license and copyright information please see LEGAL file in repository */

package matn

import (
	"../time"
	"../ganjine"
	lang "../language"
)

const wordSuggestionStructureID uint64 = 2096939569836997748

var wordSuggestionStructure = ds.DataStructure{
	ID:                2096939569836997748,
	IssueDate:         1599286551,
	ExpiryDate:        0,
	ExpireInFavorOfURN:   "", // Other structure nme
	ExpireInFavorOfID: 0,  // Other StructureID! Handy ID or Hash of ExpireInFavorOf!
	Status:            protocol.Software_PreAlpha,
	Structure:         WordSuggestion{},

	Name: map[protocol.LanguageID]string{
		protocol.LanguageEnglish: "WordSuggestion",
	},
	Description: map[protocol.LanguageID]string{
		protocol.LanguageEnglish: " store the hash index data by 32byte key and any 32byte value",
	},
	TAGS: []string{
		"",
	},
}

// WordSuggestion is standard structure to store any hash byte index!!
// It is simple secondary index e.g. hash("user@email.com")
type WordSuggestion struct {
	id       [32]byte
	start    int
	end      int
	surface  string
	class    uint8
	features []string
}
