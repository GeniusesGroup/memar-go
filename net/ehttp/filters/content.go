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
	"sort"
	"strings"

	"github.com/SabzCity/go-library/errors"
	"github.com/SabzCity/go-library/net/ehttp"
	"github.com/SabzCity/go-library/net/ehttp/ehttputil"
)

// ContentChecker check ehttputil.RequestStandards.Headers.Content.Type in header support by related route.
// Filter level: Server
// Needed filter: before(StaticRestRouter/), after()
func ContentChecker(ctx *ehttp.Context) {

	if ctx.Error != nil {
		return
	}

	//// TODO : https://tools.ietf.org/html/rfc7231#section-5.3.2

	contentAccepts := strings.Split(ctx.Request.Header.Get(ehttputil.RequestStandards.Headers.Accept.Content), ",")
	// First succseful "contentType" respond. Do it in order of request "contentAccepts"
	if len(contentAccepts) != 0 {
		for _, contentAccept := range contentAccepts {
			contentAccept = strings.Trim(contentAccept, " ")
			contentAccept = strings.Split(contentAccept, ";")[0]
			if sort.SearchStrings(SabzCityPlatformAccptedProduces, contentAccept) < len(SabzCityPlatformAccptedProduces) {
				goto ContentTypeChecker
			}
		}
		ctx.Error = errors.AcceptTypeNotSupported
	}

ContentTypeChecker:
	// We just need first part.
	// e.g. Content-Type: text/html; charset=utf-8 || Content-Type: multipart/form-data; boundary=something
	contentType := strings.Split(ctx.Request.Header.Get(ehttputil.RequestStandards.Headers.Content.Type), ";")[0]

	if contentType != "" && sort.SearchStrings(SabzCityPlatformAccptedConsumes, contentType) == len(SabzCityPlatformAccptedConsumes) {
		ctx.Error = errors.ContentTypeNotSupported
	}
}
