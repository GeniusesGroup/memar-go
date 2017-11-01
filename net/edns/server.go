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

// Package edns is Extended DNS Package
package edns

import (
	"github.com/miekg/dns"
)

// Server is main object that save filters.
type Server struct {
	Filters []Filter //Server procces filters from first to end.
}

// ServeDNS : Implemetion of "dns.Server".
func (s *Server) ServeDNS(w dns.ResponseWriter, r *dns.Msg) {

	response := &dns.Msg{}
	response.SetReply(r)

	ctx := &Context{
		Request:       r,
		Response:      response,
		StringData:    map[string]string{},
		InterfaceData: map[string]interface{}{}}

	for _, filter := range s.Filters {
		filter(ctx)

		// TODO : Think about this !
		if ctx.Error != nil {
			return
		}
	}

	w.WriteMsg(response)
}
