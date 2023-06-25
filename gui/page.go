/* For license and copyright information please see the LEGAL file in the code repository */

package gui

import (
	"libgo/mediatype"
	"libgo/protocol"
)

// Page implements protocol.GUI_Page interface
type Page struct {
	robots string
	icon   protocol.Image
	info   Information

	relatedPages []protocol.GUI_Page

	path               string         // To route page by path of HTTP-URI
	acceptedConditions map[string]any // HTTP-URI queries

	activeState  protocol.GUI_PageState
	activeStates []protocol.GUI_PageState

	mediatype.MT
}

func (p *Page) Robots() string                    { return p.robots }
func (p *Page) Icon() protocol.Image              { return p.icon }
func (p *Page) Info() protocol.GUI_Information    { return &p.info }
func (p *Page) RelatedPages() []protocol.GUI_Page { return p.relatedPages }

func (p *Page) Path() string                                    { return p.path }
func (p *Page) AcceptedCondition(key string) (defaultValue any) { return p.acceptedConditions[key] }

func (p *Page) ActiveState() protocol.GUI_PageState    { return p.activeState }
func (p *Page) ActiveStates() []protocol.GUI_PageState { return p.activeStates }
