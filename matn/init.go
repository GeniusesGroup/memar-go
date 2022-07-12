/* For license and copyright information please see LEGAL file in repository */

package matn

import (
	"../protocol"
)

func init() {
	protocol.OS.RegisterMediaType(&indexPhraseStructure)
	protocol.OS.RegisterMediaType(&indexWordStructure)
	// protocol.OS.RegisterMediaType(&)

	// protocol.App.RegisterService()
}
