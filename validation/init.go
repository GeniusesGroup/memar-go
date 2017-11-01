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

package validation

import (
	"sort"
)

var validPasswordTypes []string
var validOAuthPlatforms []string

func init() {
	// "validPasswordTypes" initialization.
	validPasswordTypes = []string{"Biology", "OTP", "Picture", "Text", "Voice"}
	sort.Strings(validPasswordTypes)

	// "validOAuthPlatforms" initialization.
	validOAuthPlatforms = []string{"Instagram", "Facebook", "Google", "Github", "Yahoo"}
	sort.Strings(validOAuthPlatforms)

	// Add validators.
	//Common validators
	validators["gender"] = IsGender
	validators["nationality"] = IsNationality
	validators["oauthPlatform"] = IsValidOAuth2Platform
	validators["passwordType"] = IsValidPasswordType
	validators["phone"] = IsValidPhone
	validators["username"] = IsUsername
	validators["limit"] = IsLimit
	validators["offset"] = IsOffset
	//SabzCity UUID
	validators["userID"] = IsUserID
	validators["realUserID"] = IsRealUserID
	validators["organizationID"] = IsOrganizationUserID
	validators["groupID"] = IsGroupID
	validators["coordinateID"] = IsCoordinateID
	validators["phoneID"] = IsPhoneID
	validators["warehouseID"] = IsWarehouseID
	validators["productID"] = IsProductID
	validators["wikiRecordID"] = IsWikiRecordID
}
