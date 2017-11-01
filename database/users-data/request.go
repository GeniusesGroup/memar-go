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
//limitations under the License..

package usersdata

// Request is standard struct of request to user area.
// MSName remove from first string of keys due duplicate data! we have it in microservice.MSDetails.OpenAPI.Info.Title!
type Request struct {
	Keys         []string    // Zones-Domain: "Zone, RecordID, SubRecordID, SubSubRecordID, ..." e.g. ["email", "123456789123456"]
	Data         interface{} // Add or Read data
	LastModified int64
	Options      Options
}

// Options is standard struct of database request Options
type Options struct {
	Transaction    bool
	CreateOnUpdate bool
	Offset         string
	Limit          string
}
