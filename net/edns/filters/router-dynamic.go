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

package filters

import (
	"time"

	"github.com/SabzCity/go-library/cache"
	"github.com/SabzCity/go-library/net/edns"
	"github.com/SabzCity/go-library/net/edns/ednsutil"
)

// staticZoneCache : Set DNS zones to this variable in static process.
var dynamicZoneCache = cache.NewObjectCache(time.Hour * 24 * 7)

// DynamicRouter :
// Needed filter: before(), after()
func DynamicRouter(ctx *edns.Context) {
	// Multiple queries in a question section have not been supported in DNS
	// due the applicability of some DNS Message Header flags (such as AA) and
	// of the RCODE field only to a single QNAME, QTYPE, and QCLASS.
	if len(ctx.Request.Question) > 1 {
		return
	}

	question := ctx.Request.Question[0]

	zone := getZone(question.Name)
	if zone == nil {
		return
	}

	responseHandler, ok := responseHandlers[question.Qtype]
	if !ok {
		return
	}

	responseHandler(ctx, zone, question.Name)

	// https://tools.ietf.org/html/rfc2308#section-2.2.1
	if len(ctx.Response.Answer) == 0 {
		SOAHandler(ctx, zone, question.Name)
	}
}

func getZone(domainName string) *ednsutil.MiniDNS {
	// Get data from cache
	zone, err := dynamicZoneCache.Get(domainName)
	if err != nil {
		// Get data from domains MS by AI
		return nil
	}

	return zone.(*ednsutil.MiniDNS)
}
