/* For license and copyright information please see the LEGAL file in the code repository */

package gui

import (
	"memar/datatype"
	gui_p "memar/gui/protocol"
	picture_p "memar/picture/protocol"
)

// Page implements gui_p.Page interface
type Page struct {
	robots string
	icon   picture_p.Image
	info   Information

	relatedPages []gui_p.Page

	path               string         // To route page by path of HTTP-URI
	acceptedConditions map[string]any // HTTP-URI queries

	activeState  gui_p.PageState
	activeStates []gui_p.PageState

	datatype.DataType
}

func (p *Page) Robots() string             { return p.robots }
func (p *Page) Icon() picture_p.Image      { return p.icon }
func (p *Page) Info() gui_p.Information    { return &p.info }
func (p *Page) RelatedPages() []gui_p.Page { return p.relatedPages }

func (p *Page) Path() string                                    { return p.path }
func (p *Page) AcceptedCondition(key string) (defaultValue any) { return p.acceptedConditions[key] }

func (p *Page) ActiveState() gui_p.PageState    { return p.activeState }
func (p *Page) ActiveStates() []gui_p.PageState { return p.activeStates }
