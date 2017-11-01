//Copyright 2017 SabzCity
//
//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.

// Package ehttp is Extended HTTP Package
package ehttp

import (
	"encoding/json"
	"encoding/xml"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"github.com/SabzCity/go-library/errors"
	"github.com/SabzCity/go-library/microservice"
	"github.com/SabzCity/go-library/net/ehttp/ehttputil"
)

// LogicHandler :
func (ctx *Context) LogicHandler(logicResponse *microservice.LogicResponse) {
	if logicResponse == nil {
		return
	}

	// Set logic error to ctx
	ctx.Error = logicResponse.Error

	// Set Last-Modified header
	if logicResponse.LastModified != 0 {
		ctx.Response.Header.Set(ehttputil.ResponseStandards.Headers.LastModified, time.Unix(0, logicResponse.LastModified).UTC().Format(http.TimeFormat))
	}

	if logicResponse.Cookie != nil {
		for i := range logicResponse.Cookie {
			ctx.Response.Header.Add(ehttputil.ResponseStandards.Headers.SetCookie, logicResponse.Cookie[i].String())
		}
	}

	if logicResponse.PayloadType == "" {
		// encode and set logic payload
		ctx.ContentEncoder(logicResponse.Payload)
	} else {
		ctx.Response.Header.Set(ehttputil.ResponseStandards.Headers.Content.Type, logicResponse.PayloadType)
		// Set payload to body of response
	}
}

// ContentDecoder decode ctx.Request.Body by ehttputil.RequestStandards.Headers.Content.Type in header ans save it in body
func (ctx *Context) ContentDecoder(body interface{}) {
	// We just need first part.
	// e.g. Content-Type: text/html; charset=utf-8 || Content-Type: multipart/form-data; boundary=something
	contentType := strings.Split(ctx.Request.Header.Get(ehttputil.RequestStandards.Headers.Content.Type), ";")[0]
	requestBody, err := ioutil.ReadAll(ctx.Request.Body)
	if err != nil {
		ctx.Error = errors.BodyNotValid
		return
	}

	// Decode request body
	switch contentType {
	case ehttputil.MimeTypes["JSON"]:
		err = json.Unmarshal(requestBody, body)
		if err != nil {
			ctx.Error = errors.BodyNotValid
			return
		}
	default:
		ctx.Error = errors.ContentTypeNotSupported
	}
}

// ContentEncoder encode body and save it in to ctx.Response.Body
func (ctx *Context) ContentEncoder(body interface{}) {
	var (
		encodedResponseBody []byte
		err                 error
	)

	//// TODO : https://tools.ietf.org/html/rfc7231#section-5.3.2

	contentAccepts := strings.Split(ctx.Request.Header.Get(ehttputil.RequestStandards.Headers.Accept.Content), ",")

	// first succseful contentType respond. Do it in order of request contentAccepts
Loop:
	for _, contentType := range contentAccepts {
		contentType = strings.Trim(contentType, " ")
		contentType = strings.Split(contentType, ";")[0]
		switch contentType {
		case ehttputil.MimeTypes["JSON"]:
			encodedResponseBody, err = json.Marshal(body)
			if err != nil {
				ctx.Error = errors.AddInformation(errors.SomeThingIsWrong, err)
				return
			}
			ctx.Response.Header.Add(ehttputil.ResponseStandards.Headers.Content.Type, ehttputil.MimeTypes["JSON"])
			break Loop
		case ehttputil.MimeTypes["XML"]:
			encodedResponseBody, err = xml.Marshal(body)
			if err != nil {
				ctx.Error = errors.AddInformation(errors.SomeThingIsWrong, err)
				return
			}
			ctx.Response.Header.Add(ehttputil.ResponseStandards.Headers.Content.Type, ehttputil.MimeTypes["XML"])
			break Loop
		case "*/*":
			encodedResponseBody, err = json.Marshal(body)
			if err != nil {
				ctx.Error = errors.AddInformation(errors.SomeThingIsWrong, err)
				return
			}
			ctx.Response.Header.Add(ehttputil.ResponseStandards.Headers.Content.Type, ehttputil.MimeTypes["JSON"])
			break Loop
		}
	}

	// Set vary for cache layers
	ctx.Response.Header.Add(ehttputil.ResponseStandards.Headers.Vary, ehttputil.RequestStandards.Headers.Accept.Content)
	ctx.Response.Body.Write(encodedResponseBody)
}

// SetFlag : Use flags as events.
func (ctx *Context) SetFlag(names ...string) {
	for _, name := range names {
		ctx.StringData[name] = "f"
	}
}

// FlagStatus : Use flags as events.
func (ctx *Context) FlagStatus(names ...string) bool {
	for _, name := range names {
		if ctx.StringData[name] == "f" {
			return true
		}
	}

	return false
}
