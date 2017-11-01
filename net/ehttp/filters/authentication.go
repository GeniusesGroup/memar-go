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

package filters

import (
	"net/http"
	"strings"
	"time"

	"github.com/SabzCity/go-library/authentication"
	"github.com/SabzCity/go-library/code"
	"github.com/SabzCity/go-library/errors"
	"github.com/SabzCity/go-library/net/ehttp"
	"github.com/SabzCity/go-library/net/ehttp/ehttputil"
)

// CookieAuthentication check user identity.
// Filter level (Server, Route): both
// Needed filter: before(), after()
func CookieAuthentication(ctx *ehttp.Context) {

	var (
		authToken    authentication.Token
		activeUserID string
	)

	activeUserToken, err := ctx.Request.Cookie("AU") // AT means ActiveUser

	// If there is an error mean guest user make the request!
	if err != nil {
		activeUserID = "Guest"
		authToken = authentication.Token{
			Delegate:  false,
			TokenID:   code.TimeUUID(20, code.NanoSecondUUID, code.DefaultSinceYear),
			ExpiredAt: time.Now().Add(time.Hour * 24 * 365).Unix(),
			Audience:  "sabz.city",
			Plan:      false}
		ATValue, _ := authentication.EncryptToken(&authToken)

		// Make Guest cookies
		guestUserToken := &http.Cookie{
			Domain:   "sabz.city",
			Name:     "AU",
			Value:    activeUserID,
			Path:     "/",
			Expires:  time.Now().Add(time.Hour * 24 * 365),
			Secure:   true,
			HttpOnly: false}
		activeToken := &http.Cookie{
			Domain:   "sabz.city",
			Name:     activeUserID,
			Value:    ATValue,
			Path:     "/",
			Expires:  time.Now().Add(time.Hour * 24 * 365),
			Secure:   true,
			HttpOnly: true}

		// Set Guest cookies to response header
		ctx.Response.Header.Add(ehttputil.ResponseStandards.Headers.SetCookie, guestUserToken.String())
		ctx.Response.Header.Add(ehttputil.ResponseStandards.Headers.SetCookie, activeToken.String())
	} else {
		activeToken, err := ctx.Request.Cookie(activeUserToken.Value) // AT means ActiveToken
		if err != nil {
			// Add WWW-Authenticate header
			ctx.Error = errors.AddInformation(errors.CookiesHaveError, err)
			ctx.ResponseReady = true
			return
		}

		err = authentication.DecryptToken(activeToken.Value, &authToken)
		if err != nil {
			// Add WWW-Authenticate header
			ctx.Error = errors.AddInformation(errors.TokenIsNotValid, err)
			ctx.ResponseReady = true
			return
		}

		// Check token host.
		if !strings.HasSuffix(strings.Split(ctx.Request.Host, ":")[0], authToken.Audience) {
			// Add WWW-Authenticate header
			ctx.Error = errors.AccessForbidden
			ctx.ResponseReady = true
			return
		}
	}

	// Set some data for other filters
	ctx.StringData["ActiveUserID"] = activeUserID
	ctx.InterfaceData["AuthenticationToken"] = authToken
	//Set activeUser in header of request to proxy location
	ctx.Request.Header.Set("X-Active-User", activeUserID)
}

// AuthorizationHeader check user identity.
// Filter level (Server, Route): both
// Needed filter: before(), after()
func AuthorizationHeader(ctx *ehttp.Context) {}
