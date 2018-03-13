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

package ehttp

import (
	"fmt"

	"github.com/SabzCity/go-library/convert"
	"github.com/SabzCity/go-library/errors"
	"github.com/SabzCity/go-library/microservice"
	"github.com/SabzCity/go-library/net/ehttp/ehttputil"
)

// HandleError is a middleware that do somthings with errors.
// Filter level (Server, Route): DON'T USE IT
func HandleError(ctx *Context) {
	if ctx.Error == nil {
		return
	}

	extendedError, ok := ctx.Error.(*errors.ExtendedError)
	if !ok {
		extendedError = errors.SomeThingIsWrong.(*errors.ExtendedError)
	}

	ctx.Response.Status = extendedError.HTTPStatus
	// Set related sabzcity error code to header.
	ctx.Response.Header.Set("X-Error", convert.IntToString(extendedError.Code))

	// Set extra info to response in development phase for easy debugging
	if microservice.MSDetails.ReleaseToken == "" {
		ctx.Response.Header.Set(ehttputil.ResponseStandards.Headers.Content.Type, ehttputil.MimeTypes["TEXT"])
		ctx.Response.Header.Set("X-Error-Addition", extendedError.Text)
		if extendedError.Information != nil {
			buffer := fmt.Sprintf("\n[SabzCity-ehttp] Error situation with extra information: %v\r\n", extendedError.Information)
			ctx.Response.Body.WriteString(buffer)
		}
	}

	// Set beautiful html response body!
}
