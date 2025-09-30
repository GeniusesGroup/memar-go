/* For license and copyright information please see the LEGAL file in the code repository */

package harf

/*
This structure like UTF based encode system but with some improvement to reduce waste of bits!
Like UTF-8, This encode system is designed in a such way that all ASCII characters use the same byte representation.
But we have very strong argue that some ASCII codes like control one is waste of first byte and we reuse them for other character!
Also it is allowed it to be self-synchronizing on both compress or uncompressed style!
We believe 4 effective bytes is enough for encoding more than 266,338,304 characters but this encode system can use for
more than 4 bytes and can increase to n bytes!
["10000000" bit || "0x80" hex || "128" byte(uint8) || "-0" int8] can be omitted any where in compressed text!
Always first bit of first byte must be "1" and First bit of second and further byte must be "1".

Each script get a byte to encode their characters! English as ASCII always start with 0 but other scripts like Arabic always start with 1.
If we have 1 effective byte, it means ASCII characters exist!
If we have 2 effective bytes, it means first one is not ASCII character code and it use to detect script ID! and 2nd byte is script character!
If we have 3 effective bytes, it means 1th&2nd is not character code and it use to detect script ID! and 3rd byte is script character!
And this rule can go for ever, But we think we don't need more than 4 byte that can encode 2,097,152 script and 266,338,304 characters!

Effective	Byte 1		Byte 2		Byte 3		Byte 4		...
byte		(int8)		(int8)		(int8)		(int8)		...
1			0xxxxxxx	10000000	10000000	10000000	...
2			0sssssss	1xxxxxxx	10000000	10000000	...
3			0sssssss	1sssssss	1xxxxxxx	10000000	...
4			0sssssss	1sssssss	1sssssss	1xxxxxxx	...
...			...			...			...			...			...
*/

// CharacterDetail use to store detail for a character script
type CharacterDetail struct {
	Code                [4]byte
	Description         []string
	ScriptsUsesIDs      []uint32 // Automatically fulfill with scripts data
	RelatedCharacterIDs [][4]byte
	UnicodeID           rune
	Dir                 uint8
}

// UnicodeCharacter use to convert from unicode || to unicode
var UnicodeCharacter map[rune]*CharacterDetail

// Characters store all characters scripts
var Characters = map[[4]byte]CharacterDetail{
	[4]byte{0, 128, 128, 128}: {Description: []string{"Null", "تهی، نیم فاصله"}},
	[4]byte{1, 128, 128, 128}: {Description: []string{"New Line", "خط بعد"}},
	[4]byte{2, 128, 128, 128}: {Description: []string{"New Page", "صفحه بعد"}},

	[4]byte{0, 1, 128, 128}: {},
}
