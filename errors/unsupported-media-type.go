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

// 415 Unsupported Media Type - X-Error Related

package errors

import (
	"net/http"
)

//Declare Errors SabzCity Code
const (
	contentTypeNotSupported = 41500 + (iota + 1)

	contentEncodingNotSupported
)

//Declare Errors Detials
var (
	ContentTypeNotSupported = New("Payload of request (Content-Type) has not supported by platform", contentTypeNotSupported, http.StatusUnsupportedMediaType)

	ContentEncodingNotSupported = New("Payload encoding (Content-Encoding) of request has not supported by platform", contentEncodingNotSupported, http.StatusUnsupportedMediaType)
)
