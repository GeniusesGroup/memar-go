/* For license and copyright information please see LEGAL file in repository */

package letsencrypt

// GenerateAccountReq is request structure of any generateAccount on any protocol
type GenerateAccountReq struct {
	Nonce                string
	TermsOfServiceAgreed bool
	Contact              string
}

func (ga *GenerateAccountReq) jsonDecoder() {}

func (ga *GenerateAccountReq) jsonEncoder() {}

// GenerateAccountRes is response structure of any generateAccount on any protocol
type GenerateAccountRes struct {
}

func (ga *GenerateAccountRes) jsonDecoder() {}

func (ga *GenerateAccountRes) jsonEncoder() {}
