/* For license and copyright information please see LEGAL file in repository */

package service

import (
	"../protocol"
)

// ServiceDetail store detail about an service
type ServiceDetail struct {
	languageID  protocol.LanguageID
	domain      string
	summary     string
	overview    string
	description string
	tags        []string
}

func (sd *ServiceDetail) Language() protocol.LanguageID { return sd.languageID }
func (sd *ServiceDetail) Domain() string                { return sd.domain }
func (sd *ServiceDetail) Summary() string               { return sd.summary }
func (sd *ServiceDetail) Overview() string              { return sd.overview }
func (sd *ServiceDetail) Description() string           { return sd.description }
func (sd *ServiceDetail) TAGS() []string                { return sd.tags }

/*
*********** Service Codecs ***********
 */
