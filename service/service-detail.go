/* For license and copyright information please see LEGAL file in repository */

package service

// ServiceDetail store detail about an service
type ServiceDetail struct {
	name        string
	description string
	tags        []string
}

func (sd *ServiceDetail) Name() string        { return sd.name }
func (sd *ServiceDetail) Description() string { return sd.description }
func (sd *ServiceDetail) TAGS() []string      { return sd.tags }

/*
*********** Service Codecs ***********
 */
