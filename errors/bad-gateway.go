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

// 502 Bad Gateway - X-Error Related

package errors

import (
	"net/http"
)

//Declare Errors SabzCity Code.
const (
	canNotConnectToResource = 50200 + (iota + 1)

	proxyRoutingError

	hostNotExist

	urlIsWrong
)

//Declare Errors Detials.
var (
	CanNotConnectToResource = New("SabzCity Platform can't connect to the Container, Proxy or special resource", canNotConnectToResource, http.StatusBadGateway)

	ProxyRoutingError = New("Proxy server have error in routing the request", proxyRoutingError, http.StatusBadGateway)

	HostNotExist = New("The requested host not exist or accessible", hostNotExist, http.StatusBadGateway)

	URLIsWrong = New("The URL is not found", urlIsWrong, http.StatusBadGateway)
)
