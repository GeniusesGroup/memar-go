/* For license and copyright information please see the LEGAL file in the code repository */

package gui

type ScrollEvent uint8

const (
	ScrollEvent_Unset ScrollEvent = iota
	ScrollEvent_HOME
	ScrollEvent_END
	ScrollEvent_STEP_PLUS
	ScrollEvent_STEP_MINUS
	ScrollEvent_PAGE_PLUS
	ScrollEvent_PAGE_MINUS
	ScrollEvent_POS
	ScrollEvent_SLIDER_RELEASED
	ScrollEvent_CORNER_PRESSED
	ScrollEvent_CORNER_RELEASED
)
