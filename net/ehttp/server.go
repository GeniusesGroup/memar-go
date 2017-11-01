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
	"bytes"
	"net/http"

	"github.com/SabzCity/go-library/net/ehttp/ehttputil"
)

// Server is the basic multilayer router. Implements "http.Handler".
type Server struct {
	Filters []Filter //Server procces filters from first to end.
}

// ServeHTTP is a standard "net/http" function. Shouldn't be used directly, "http.Serve()" will use it.
func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := Context{
		Request: r,
		Response: &Response{
			Header: http.Header{},
			Body:   bytes.NewBuffer([]byte{})},
		ResponseReady: false,
		StringData:    map[string]string{},
		InterfaceData: map[string]interface{}{}}

	// TODO : Research in limiting concurrency by hardware limits.

	// Recover from panics if exist.
	defer panicHandler(w)

	for _, filter := range s.Filters {
		filter(&ctx)
		if ctx.ResponseReady {
			break
		}
	}

	// Handle error and reject request if error exists.
	HandleError(&ctx)

	// Check status code is correct and set correctly, Otherwise set 200
	if ctx.Response.Status == 0 {
		if ctx.Response.Body.Len() == 0 {
			ctx.Response.Status = ehttputil.ResponseStandards.Status.NoContent.Code
		} else {
			ctx.Response.Status = ehttputil.ResponseStandards.Status.OK.Code
		}
	}

	// Add ehttp server signature
	ctx.Response.Header.Set("Server", "SCP") // SCP means "SabzCityPlatform".

	// Write response Headers
	ehttputil.CopyHeaders(w.Header(), ctx.Response.Header)
	// Wrtite response status
	w.WriteHeader(ctx.Response.Status)
	// Write response Body with||without compress!
	ctx.Response.Body.WriteTo(w)
}
