/* For license and copyright information please see LEGAL file in repository */

package service

import (
	etime "../earth-time"
	"../json"
	"../log"
	"../protocol"
	"../urn"
)

// Service store needed data for a service to implement protocol.Service when embed to other struct that implements other methods!
type Service struct {
	urn    urn.Giti
	uri    string // Fill just if any http like type handler needed! Simple URI not variabale included! API services can set like "/apis?{{.ServiceID}}" but it is not efficient, find services by ID.
	status protocol.ServiceStatus

	issueDate          int64
	expiryDate         int64
	expireInFavorOfURN string // Other service URN

	detail map[protocol.LanguageID]*ServiceDetail

	authorization Authorization

	JSON   []byte `json:"-" syllab:"-"`
	Syllab []byte `json:"-" syllab:"-"`
}

// New returns a new error!
func New(urn, uri string, status protocol.ServiceStatus, issueDate int64) (s Service) {
	if urn == "" {
		log.Fatal("Try to make a new service without give valid URN! It is rule to add more detail about service before register it!")
	}

	s = Service{
		uri:    uri,
		status: status,

		issueDate: issueDate,

		detail: make(map[protocol.LanguageID]*ServiceDetail),
	}
	s.urn.Init(urn)
	return
}

func (s *Service) SetDetail(lang protocol.LanguageID, name, description string, tags []string) {
	var ok bool
	_, ok = s.detail[lang]
	if ok {
		panic("Can't change service detail after first set! Ask service holder to change details!!")
	}
	s.detail[lang] = &ServiceDetail{
		name:        name,
		description: description,
		tags:        tags,
	}
}

func (s *Service) SetAuthorization(crud protocol.CRUD, userType protocol.UserType) {
	s.authorization.crud = crud
	s.authorization.userType = userType
}

func (s *Service) Expired(expiryDate int64, inFavorOf string) {
	s.expiryDate = expiryDate
	s.expireInFavorOfURN = inFavorOf
}

func (s *Service) Detail(lang protocol.LanguageID) protocol.ServiceDetail {
	var d, _ = s.detail[lang]
	return d
}
func (s *Service) URN() protocol.GitiURN          { return &s.urn }
func (s *Service) URI() string                    { return s.uri }
func (s *Service) Status() protocol.ServiceStatus { return s.status }
func (s *Service) IssueDate() protocol.Time       { return etime.Time(s.issueDate) }
func (s *Service) ExpiryDate() protocol.Time      { return etime.Time(s.issueDate) }
func (s *Service) CRUDType() protocol.CRUD        { return s.authorization.crud }
func (s *Service) UserType() protocol.UserType    { return s.authorization.userType }

/*
*********** Handlers ***********
not-implemented handlers of the service.
*/

func (s *Service) ServeSRPC(st protocol.Stream, srpcReq protocol.SRPCRequest) (res protocol.Syllab, err protocol.Error) {
	err = ErrServiceNotAcceptSRPC
	return
}
func (s *Service) ServeHTTP(st protocol.Stream, httpReq protocol.HTTPRequest, httpRes protocol.HTTPResponse) (err protocol.Error) {
	err = ErrServiceNotAcceptHTTP
	return
}
func (s *Service) ServeCLI(st protocol.Stream) (err protocol.Error) {
	err = ErrServiceNotAcceptCLI
	return
}

/*
*********** Service Encoders & Decoders ***********
 */

func (s *Service) FromSyllab() (err protocol.Error) {
	return
}

func (s *Service) ToSyllab() {
}

func (s *Service) LenOfSyllabStack() (ln uint32) {
	return
}

func (s *Service) LenOfSyllabHeap() (ln uint32) {
	return
}

func (s *Service) LenAsSyllab() uint64 {
	return uint64(s.LenOfSyllabStack() + s.LenOfSyllabHeap())
}

func (s *Service) jsonDecoder() (err protocol.Error) {
	err = json.Unmarshal(s.JSON, s)
	return
}

func (s *Service) ToJSON() {
	s.JSON, _ = json.Marshal(s)
}
