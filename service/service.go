/* For license and copyright information please see LEGAL file in repository */

package service

import (
	"../protocol"
	"../urn"
)

// Service store needed data for a service to implement protocol.Service when embed to other struct that implements other methods!
type Service struct {
	urn    urn.Giti
	uri    string // Fill just if any http like type handler needed! Simple URI not variabale included! API services can set like "/m?{{.ServiceID}}" but it is not efficient, find services by ID.
	status protocol.SoftwareStatus
	weight protocol.Weight // Use to queue requests by services weights

	issueDate       protocol.TimeUnixSec
	expiryDate      protocol.TimeUnixSec
	expireInFavorOf urn.Giti

	detail map[protocol.LanguageID]*ServiceDetail

	authorization Authorization
}

// New returns a new error!
func New(urn, uri string, status protocol.SoftwareStatus, issueDate protocol.TimeUnixSec) (s *Service) {
	if urn == "" {
		panic("Try to make a new service without give valid URN! It is rule to add more detail about service before register it!")
	}

	s = &Service{
		uri:    uri,
		status: status,

		issueDate: issueDate,

		detail: make(map[protocol.LanguageID]*ServiceDetail),
	}
	s.urn.Init(urn)
	return
}

func (s *Service) SetWeight(weight protocol.Weight) *Service {
	s.weight = weight
	return s
}

func (s *Service) SetAuthorization(crud protocol.CRUD, userType protocol.UserType) *Service {
	s.authorization.crud = crud
	s.authorization.userType = userType
	return s
}

func (s *Service) SetDetail(lang protocol.LanguageID, domain, summary, overview, description string, tags []string) *Service {
	var _, ok = s.detail[lang]
	if ok {
		panic("Can't change service detail after first set! Ask service holder to change details!!")
	}
	s.detail[lang] = &ServiceDetail{
		languageID:  lang,
		domain:      domain,
		summary:     summary,
		overview:    overview,
		description: description,
		tags:        tags,
	}
	return s
}

func (s *Service) Expired(expiryDate protocol.TimeUnixSec, inFavorOf string) Service {
	s.expiryDate = expiryDate
	s.expireInFavorOf.Init(inFavorOf)
	return *s
}

func (s *Service) Detail(lang protocol.LanguageID) protocol.ServiceDetail { return s.detail[lang] }
func (s *Service) URN() protocol.GitiURN                                  { return &s.urn }
func (s *Service) URI() string                                            { return s.uri }
func (s *Service) Status() protocol.SoftwareStatus                        { return s.status }
func (s *Service) IssueDate() protocol.TimeUnixSec                        { return s.issueDate }
func (s *Service) ExpiryDate() protocol.TimeUnixSec                       { return s.issueDate }
func (s *Service) ExpireInFavorOf() protocol.GitiURN                      { return &s.expireInFavorOf }
func (s *Service) Weight() protocol.Weight                                { return s.weight }
func (s *Service) CRUDType() protocol.CRUD                                { return s.authorization.crud }
func (s *Service) UserType() protocol.UserType                            { return s.authorization.userType }

/*
*********** Handlers ***********
not-implemented handlers of the service.
*/

func (s *Service) ServeSRPC(st protocol.Stream) (err protocol.Error) {
	err = ErrServiceNotAcceptSRPC
	return
}
func (s *Service) ServeSRPCDirect(conn protocol.Connection, request []byte) (response []byte, err protocol.Error) {
	err = ErrServiceNotAcceptSRPCDirect
	return
}
func (s *Service) ServeHTTP(st protocol.Stream, httpReq protocol.HTTPRequest, httpRes protocol.HTTPResponse) (err protocol.Error) {
	err = ErrServiceNotAcceptHTTP
	return
}
