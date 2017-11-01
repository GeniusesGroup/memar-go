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
	"github.com/SabzCity/go-library/net/ehttp"
	"github.com/SabzCity/go-library/net/ehttp/ehttputil"
)

// Forwarded : Add and manage "Forwarded" header.
// Filter level: just Server
func Forwarded(ctx *ehttp.Context) {

	forwardedHeader := ehttputil.GetForwarded(ctx.Request)
	forwardedHeader.By = ProxyLocalIP
	ehttputil.SetForwarded(ctx.Request.Header, forwardedHeader)
}
