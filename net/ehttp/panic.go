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
	"fmt"
	"net/http"
	"runtime"

	"github.com/SabzCity/go-library/convert"
	"github.com/SabzCity/go-library/errors"
	"github.com/SabzCity/go-library/microservice"
)

func panicHandler(w http.ResponseWriter) {
	if panic := recover(); panic != nil {
		extendedError := errors.ServerPanic.(*errors.ExtendedError)

		w.Header().Set("X-Error", convert.IntToString(extendedError.Code))
		w.WriteHeader(extendedError.HTTPStatus)

		if !microservice.MSDetails.Production {
			w.Header().Set("X-Error-Addition", extendedError.Text)

			var buffer bytes.Buffer
			buffer.WriteString(fmt.Sprintf("[SabzCity-ehttp] recover from panic situation: - %v\r\n", panic))
			for i := 2; ; i++ {
				_, file, line, ok := runtime.Caller(i)
				if !ok {
					break
				}
				buffer.WriteString(fmt.Sprintf("    %s:%d\r\n", file, line))
			}

			w.Write(buffer.Bytes())
		}
	}
}
