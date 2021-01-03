/* For license and copyright information please see LEGAL file in repository */

package matn

import (
	er "../error"
)

// TextIndexSetReq is request structure of TextIndex()
type TextIndexSetReq struct {
	RecordID           [32]byte
	RecordStructure    uint64
	RecordPrimaryKey   [32]byte // Store any primary ID or any data up to 32 byte length e.g. ID
	RecordSecondaryKey [32]byte // Store any secondary ID or any data up to 32 byte length e.g. GroupID
	RecordOwnerID      [32]byte
	RecordFieldID      uint8
	Text               string
}

// TextIndexSet index given text
func TextIndexSet(req *TextIndexSetReq) (err *er.Error) {
	var indexes = WordTokenization(req)

	for _, index := range indexes {
		err = index.SaveOrUpdate()
	}

	// TODO::: need to index or do anything here about phrases???

	return
}

// WordTokenization uses the delimiters categorized under [Dash, Hyphen, Pattern_Syntax, Quotation_Mark, Terminal_Punctuation, White_Space]
// https://github.com/jdkato/prose
func WordTokenization(req *TextIndexSetReq) (indexes map[string]*IndexWord) {
	// var fields = strings.Fields(req.Text)
	// fmt.Println(fields)

	indexes = map[string]*IndexWord{}

	var (
		ok                bool
		index             *IndexWord
		word              string
		wordStart         bool
		sentenceEnd       bool
		lastSentenceIndex int

		WordOffsetInSentence uint64 //  Position of the word in the sentence
		WordOffsetInText     uint64 //  Position of the word in the text
		OffsetInSentence     uint64 //  First word charecter possition in the sentence
		OffsetInText         uint64 //  First word charecter possition in the text
	)

	// TODO::: drop language-specific stop words??? (e.g. in English, the, an, a, and, etc.)
	for i, char := range req.Text + " " { // TODO::: hack situation!! need to fix this and remove + " "
		switch char {
		case ' ', '\t', '\n', '\v', '\f', '\r', 0x85, 0xA0, '\'', '"', '`', ':', ',', '-': // unicode.IsSpace(char)
			if wordStart {
				word = req.Text[OffsetInText:i]
			} else {
				OffsetInSentence = uint64(i-lastSentenceIndex) + 1 // indicate next charecter as start of word
				OffsetInText = uint64(i) + 1                       // indicate next charecter as start of word
				continue
			}
		case '_':
			// TODO:::
			continue
		case '#':
			// TODO:::
			continue
		case '@':
			// TODO:::
			continue
		case '$':
			// TODO:::
			continue
		case '.', ';', '?', '!':
			sentenceEnd = true
			lastSentenceIndex = i
			if wordStart {
				word = req.Text[OffsetInText:i]
			} else {
				continue
			}
		case '[', '(', '{':
		case ']', ')', '}':
		default:
			wordStart = true
			continue
		}

		index, ok = indexes[word]
		if ok {
			index.Tokens = append(index.Tokens, WordToken{
				RecordID:             req.RecordID,
				RecordFieldID:        req.RecordFieldID,
				WordOffsetInSentence: WordOffsetInSentence,
				WordOffsetInText:     WordOffsetInText,
				OffsetInSentence:     OffsetInSentence,
				OffsetInText:         OffsetInText,
			})
		} else {
			index = &IndexWord{
				Word:               word,
				RecordStructure:    req.RecordStructure,
				RecordPrimaryKey:   req.RecordPrimaryKey,
				RecordSecondaryKey: req.RecordSecondaryKey,
				RecordOwnerID:      req.RecordOwnerID,
				Tokens: []WordToken{
					{
						RecordID:             req.RecordID,
						RecordFieldID:        req.RecordFieldID,
						WordOffsetInSentence: WordOffsetInSentence,
						WordOffsetInText:     WordOffsetInText,
						OffsetInSentence:     OffsetInSentence,
						OffsetInText:         OffsetInText,
					},
				},
			}
			indexes[word] = index
		}

		OffsetInSentence = uint64(i-lastSentenceIndex) + 1 // indicate next charecter as start of word
		OffsetInText = uint64(i) + 1                       // indicate next charecter as start of word
		WordOffsetInSentence++
		WordOffsetInText++
		wordStart = false

		if sentenceEnd {
			sentenceEnd = false
			WordOffsetInSentence = 0
			OffsetInSentence = 0
		}
	}
	return
}
