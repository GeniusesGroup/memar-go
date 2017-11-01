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

package authorization

import (
	"sort"

	"github.com/SabzCity/go-library/convert"
	"github.com/SabzCity/go-library/errors"
)

// Request : The Request values for authorization
type Request struct {
	RequestedUserID       string // User that request to server!
	OwnerUserID           string
	MethodName            string // MSName.MethodName
	RecordID              string
	RequestedUserLocation string // GPS location
}

// Authorization :
func Authorization(r Request) error {

	// Check if method is public. Accept this type of requests.
	if r.OwnerUserID == "" {
		return nil
	}

	// Check owner user want to access.
	if r.RequestedUserID == r.OwnerUserID {
		return nil
	}

	// Check owner user type and Group||Relation access check
	if UUID, err := convert.StringToUInt64(r.OwnerUserID); err == nil {
		if UUID >= 100000000000000 && UUID <= 799999999999999 {
			// Check Relations Access

		} else if UUID >= 800000000000000 && UUID <= 999999999999999 {
			userJoinedGroups := UserJoinedGroups(r.RequestedUserID)
			orgGroups := OrgGroups(r.OwnerUserID)

			for i := range userJoinedGroups {
				if sort.SearchStrings(orgGroups, userJoinedGroups[i]) != len(orgGroups) {
					groupsDetails := GroupDetails(userJoinedGroups[i])

					// TODO : Check other group AccessControl
					if sort.SearchStrings(groupsDetails.AccessControl.Which, r.MethodName) != len(groupsDetails.AccessControl.Which) {
						return nil
					}
				}
			}

			// TODO : Check Registred + Guest one of org. becuase may be userID not in any org group!
		}
	}

	return errors.AccessForbidden
}
