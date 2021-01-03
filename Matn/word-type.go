/* For license and copyright information please see LEGAL file in repository */

package matn

// WordType indicate any word||phrase type
type WordType uint16

// Word||phrases types
const (
	WordTypeUnset    WordType = iota
	WordTypeTag               // Mentions something e.g. tag a user > @omid
	WordTypeHashTag           // Mentions some topic e.g. #Victory
	WordTypeURI               // Email addresses > Jane.Doe@example.com || URLs > https://github.com/jdkato/prose
	WordTypeEmoticon          // :-), >:(, o_0, etc.

	// ...
	/*
		CC	conjunction, coordinating
		CD	cardinal number
		DT	determiner
		EX	existential there
		FW	foreign word
		IN	conjunction, subordinating or preposition
		JJ	adjective
		JJR	adjective, comparative
		JJS	adjective, superlative
		LS	list item marker
		MD	verb, modal auxiliary
		NN	noun, singular or mass
		NNP	noun, proper singular
		NNPS	noun, proper plural
		NNS	noun, plural
		PDT	predeterminer
		POS	possessive ending
		PRP	pronoun, personal
		PRP$	pronoun, possessive
		RB	adverb
		RBR	adverb, comparative
		RBS	adverb, superlative
		RP	adverb, particle
		SYM	symbol
		TO	infinitival to
		UH	interjection
		VB	verb, base form
		VBD	verb, past tense
		VBG	verb, gerund or present participle
		VBN	verb, past participle
		VBP	verb, non-3rd person singular present
		VBZ	verb, 3rd person singular present
		WDT	wh-determiner
		WP	wh-pronoun, personal
		WP$	wh-pronoun, possessive
		WRB	wh-adverb
	*/
)
