/* For license and copyright information please see LEGAL file in repository */

package lang

/*
This structure like UTF based encode system but with some improvement to reduce waste of bits!
Like UTF-8, This encode system is designed in a such way that all ASCII characters use the same byte representation.
But we have very strong argue that some ASCII codes like control one is waste of first byte and we reuse them for other charecter!
Also it is allowed it to be self-synchronizing on both compress or uncompressed style!
We believe 4 effective bytes is enough for encoding more than 266,338,304 charecters but this encode system can use for
more than 4 bytes and can increase to n bytes!
["10000000" bit || "0x80" hex || "128" byte(uint8) || "-0" int8] can be omitted any where in compressed text!
Always first bit of first byte must be "1" and First bit of second and further byte must be "1".

Each script get a byte to encode their charecters! English as ASCII always start with 0 but other scripts like Arabic always start with 1.
If we have 1 effective byte, it means ASCII charecters exist!
If we have 2 effective bytes, it means first one is not ASCII charecter code and it use to detect script ID! and 2nd byte is script charecter!
If we have 3 effective bytes, it means 1th&2nd is not charecter code and it use to detect script ID! and 3rd byte is script charecter!
And this rule can go for ever, But we think we don't need more than 4 byte that can encode 2,097,152 script and 266,338,304 charecters!

Effective	Byte 1		Byte 2		Byte 3		Byte 4		...
byte		(int8)		(int8)		(int8)		(int8)		...
1			0xxxxxxx	10000000	10000000	10000000	...
2			0sssssss	1xxxxxxx	10000000	10000000	...
3			0sssssss	1sssssss	1xxxxxxx	10000000	...
4			0sssssss	1sssssss	1sssssss	1xxxxxxx	...
...			...			...			...			...			...
*/

// CharecterDetail use to store detail for a charecter script
type CharecterDetail struct {
	Code                [4]byte
	Description         []string
	ScriptsUsesIDs      []uint32 // Automatically fulfill with scripts data
	RelatedCharecterIDs [][4]byte
	UnicodeID           rune
	Dir                 uint8
}

// UnicodeCharecter use to convert from unicode || to unicode
var UnicodeCharecter map[rune]*CharecterDetail

// Charecters store all characters scripts
var Charecters = map[[4]byte]CharecterDetail{
	[4]byte{0, 128, 128, 128}: CharecterDetail{Description: []string{"Null", "تهی، نیم فاصله"}},
	[4]byte{1, 128, 128, 128}: CharecterDetail{Description: []string{"New Line", "خط بعد"}},
	[4]byte{2, 128, 128, 128}: CharecterDetail{Description: []string{"New Page", "صفحه بعد"}},

	[4]byte{0, 1, 128, 128}: CharecterDetail{},
}
