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

package microservice

import (
	"github.com/SabzCity/go-library/documentation/openapi"
)

// MSDetails : The global details of MS.
var MSDetails Details

// Details : Standard structure for microservice details.
type Details struct {
	OpenAPI      *openapi.OpenAPI
	HandlersPort HandlersPort
	Production   bool
	Other        map[string]interface{} // All other specify details for MS.
}

// HandlersPort : information to listen.
type HandlersPort struct {
	Restful string // The running port that used for serving restful service.
	MQTT    string
	RPC     string
}
