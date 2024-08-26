/* For license and copyright information please see the LEGAL file in the code repository */

package gui_p

// Element
type Element interface {
	Name() string // It must be unique e.g. product
	Type() Element_Type
	Status() Element_Status

	HTML(outer bool) string
	Text() string

	// Create new element as copy of existing element, new element is a full (deep) copy of the element and
	// is disconnected initially from the DOM.
	Clone() Element
	CloneTo(dom Element)

	GetElementById(id string) Element
	GetElementsByName(name string) Element
	GetElementsByClassName(class string) []Element
	GetElementsByTagName(tag string) []Element
	GetElementsByTagNameNS(tag string) []Element

	HasFocus() bool

	Append(Element)
	Prepend(Element)

	Parent() Element
	NthChild(n int) Element
	// NextSibling()
	// PreviousSibling()
	// RemoveChild()
	// ReplaceChild()

	// CaptureEvents()
	// CreateEvent()
	// ReleaseEvents()

	// EventTarget[Element]
	// https://developer.mozilla.org/en-US/docs/Web/API/Element/clientHeight
	Scroll
}

type Element_Type uint8
