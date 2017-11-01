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

package oauth2

import (
	"net/http"
	"time"

	"github.com/SabzCity/go-library/errors"
	restful "github.com/emicklei/go-restful"
	"github.com/markbates/goth/gothic"
)

// UserInformation : all information that returned by provider.
type UserInformation struct {
	RawData           map[string]interface{}
	Provider          *string
	Email             *string
	Name              *string
	FirstName         *string
	LastName          *string
	NickName          *string
	Description       *string
	UserID            *string
	AvatarURL         *string
	Location          *string
	AccessToken       *string
	AccessTokenSecret *string
	RefreshToken      *string
	ExpiresAt         *time.Time
}

// CompleteUserAuth : Complete oauth2 action and return the information.
func CompleteUserAuth(req *restful.Request, res *restful.Response, provider string) (*UserInformation, bool) {

	req.Request.URL.Query().Add("provider", provider)

	data, err := gothic.CompleteUserAuth(res, req.Request)
	if err != nil {
		errors.ResponseError(errors.SomeThingIsWrong, http.StatusInternalServerError, res)
		return nil, false
	}

	return &UserInformation{RawData: data.RawData,
		Provider:          &data.Provider,
		Email:             &data.Email,
		Name:              &data.Name,
		FirstName:         &data.FirstName,
		LastName:          &data.LastName,
		NickName:          &data.NickName,
		Description:       &data.Description,
		UserID:            &data.UserID,
		AvatarURL:         &data.AvatarURL,
		Location:          &data.Location,
		AccessToken:       &data.AccessToken,
		AccessTokenSecret: &data.AccessTokenSecret,
		RefreshToken:      &data.RefreshToken,
		ExpiresAt:         &data.ExpiresAt}, true
}
