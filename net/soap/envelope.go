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

package soap

import (
	"github.com/clbanning/mxj"
)

// Envelope : Standard SOAP 1.1 envelope structure.
type Envelope struct {
	Header map[string]interface{}
	Body   struct {
		Content map[string]interface{}
		Fault   *Fault
	}
}

// Fault : SOAP fault code and etc.
type Fault struct {
	FaultCode   *string
	FaultString *string
	FaultActor  *string
	Detail      *string
}

// Marshal : Return SOAP encoding.
func Marshal(envelope *Envelope) ([]byte, error) {

	document := map[string]interface{}{
		"-xmlns": "http://schemas.xmlsoap.org/soap/envelope/",
		"Header": envelope.Header,
		"Body":   envelope.Body.Content,
	}
	if envelope.Body.Fault != nil {
		document["Body"].(map[string]interface{})["Fault"] = map[string]interface{}{
			"faultcode":   envelope.Body.Fault.FaultCode,
			"faultstring": envelope.Body.Fault.FaultString,
			"faultactor":  envelope.Body.Fault.FaultActor,
			"detail":      envelope.Body.Fault.Detail,
		}
	}

	return mxj.Map(document).Xml("Envelope")
}

// Unmarshal : Convert SOAP to envelope.
func Unmarshal(data []byte) (*Envelope, error) {

	document, err := mxj.NewMapXml(data, true)
	if err != nil {
		return nil, err
	}
	envelope := &Envelope{}

	header, _ := document.ValueForPath("Envelope.Header")
	envelope.Header = header.(map[string]interface{})

	fault, _ := document.ValueForPath("Envelope.Body.Fault")
	if fault != nil {
		faultMap := fault.(map[string]interface{})

		envelope.Body.Fault = &Fault{}
		if faultMap["faultcode"] != nil {
			envelope.Body.Fault.FaultCode = faultMap["faultcode"].(*string)
		}
		if faultMap["faultstring"] != nil {
			envelope.Body.Fault.FaultString = faultMap["faultstring"].(*string)
		}
		if faultMap["faultactor"] != nil {
			envelope.Body.Fault.FaultActor = faultMap["faultactor"].(*string)
		}
		if faultMap["detail"] != nil {
			envelope.Body.Fault.Detail = faultMap["detail"].(*string)
		}

		document.Remove("Envelope.Body.Fault")
	}

	body, _ := document.ValueForPath("Envelope.Body")
	envelope.Body.Content = body.(map[string]interface{})

	return envelope, nil
}
