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
	"time"

	"github.com/SabzCity/go-library/cache"
)

var userJoinedGroupsCache = cache.NewObjectCache(time.Hour * 24 * 7)
var orgGroupsCache = cache.NewObjectCache(time.Hour * 24 * 7)
var groupsDetailsCache = cache.NewObjectCache(time.Hour * 24 * 7)

// UserJoinedGroups :
func UserJoinedGroups(userID string) []string {
	userJoinedGroupsCached, err := userJoinedGroupsCache.Get(userID)
	if err != nil {
		// Get and cache needed data
		// TODO : Complete with "Groups" MS.
		userJoinedGroups := []string{}
		userJoinedGroupsCache.Set(userID, userJoinedGroups, -1)

		return userJoinedGroups
	}

	return userJoinedGroupsCached.([]string)
}

// OrgGroups :
func OrgGroups(orgID string) []string {
	orgGroupsCached, err := orgGroupsCache.Get(orgID)
	if err != nil {
		// Get and cache needed data
		// TODO : Complete with "Groups" MS.
		orgGroups := []string{}
		userJoinedGroupsCache.Set(orgID, orgGroups, -1)

		return orgGroups
	}

	return orgGroupsCached.([]string)
}

// GroupDetails :
func GroupDetails(groupID string) Group {
	groupsDetailsCached, err := groupsDetailsCache.Get(groupID)
	if err != nil {
		// Get and cache needed data
		// TODO : Complete with "Groups" MS.
		groupsDetails := Group{}
		userJoinedGroupsCache.Set(groupID, groupsDetails, -1)

		return groupsDetails
	}

	return groupsDetailsCached.(Group)
}
