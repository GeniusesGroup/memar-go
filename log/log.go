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

package log

import (
	"log"

	"os"

	"github.com/SabzCity/go-library/encoding/ejson"
	"github.com/SabzCity/go-library/errors"
	"github.com/SabzCity/go-library/microservice"
)

// Fatal : Send fatal log to log service.
func Fatal(err error) {

	if extendedError, ok := err.(*errors.ExtendedError); ok {
		jDoc, _ := ejson.Marshal(map[string]interface{}{
			"Code":        extendedError.Code,
			"Text":        extendedError.Text,
			"Information": extendedError.Information})

		if !microservice.MSDetails.Production {
			log.Println("FATAL : ", jDoc)
		}
		if extendedError.HTTPStatus == 0 {
			os.Exit(1)
		}
	} else {
		log.Println(err)
	}

	// TODO : We have to send this to log MS too.
}
