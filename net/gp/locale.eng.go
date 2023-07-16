//go:build lang_eng

/* For license and copyright information please see the LEGAL file in the code repository */

package chapar

import (
	"libgo/protocol"
)

const domainEnglish = "Giti Network"

func init() {
	ErrFrameTooShort.SetDetail(protocol.LanguageEnglish, domainEnglish, "Frame Too Short",
		"Giti frame is empty or too short than standard header. It must include 44Byte header plus 16Byte min Payload",
		"",
		"",
		nil)
	ErrFrameArrivedAnterior.SetDetail(protocol.LanguageEnglish, domainEnglish, "Frame Arrived Anterior",
		"New frame arrive before some expected frame arrived. Usually cause of drop frame detection or high latency occur for some frame",
		"",
		"",
		nil)
	ErrFrameArrivedPosterior.SetDetail(protocol.LanguageEnglish, domainEnglish, "Frame Arrived Posterior",
		"New frame arrive after some expected frame arrived. Usually cause of drop frame detection or high latency occur for some frame",
		"",
		"",
		nil)
}
