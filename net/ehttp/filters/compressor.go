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
	"bytes"
	"compress/flate"
	"compress/gzip"
	"strings"

	"github.com/SabzCity/go-library/errors"
	"github.com/SabzCity/go-library/net/ehttp"
	"github.com/SabzCity/go-library/net/ehttp/ehttputil"
)

// Compressor Handle response compression.
// Filter level (Server, Route): both
func Compressor(ctx *ehttp.Context) {
	// First check body not compresed before.
	if ctx.Response.Header.Get(ehttputil.ResponseStandards.Headers.Content.Encoding) == "" {
		for _, encodingType := range strings.Split(ctx.Request.Header.Get(ehttputil.RequestStandards.Headers.Accept.Encoding), ",") {
			encodingType = strings.Trim(encodingType, " ")
			switch encodingType {
			case ehttputil.CompressionTypes.Gzip:
				ctx.Response.Header.Set(ehttputil.ResponseStandards.Headers.Content.Encoding, ehttputil.CompressionTypes.Gzip)
				ctx.Response.Header.Add(ehttputil.ResponseStandards.Headers.Vary, ehttputil.RequestStandards.Headers.Accept.Encoding)

				buffer := bytes.NewBuffer([]byte{})
				gzWriter := gzip.NewWriter(buffer)

				if _, err := ctx.Response.Body.WriteTo(gzWriter); err != nil {
					ctx.Error = errors.AddInformation(errors.SomeThingIsWrong, map[string]interface{}{"ExtraInfo": err})
					return
				}

				gzWriter.Close()
				ctx.Response.Body = buffer
				return
			case ehttputil.CompressionTypes.Deflate:
				ctx.Response.Header.Set(ehttputil.ResponseStandards.Headers.Content.Encoding, ehttputil.CompressionTypes.Deflate)
				ctx.Response.Header.Add(ehttputil.ResponseStandards.Headers.Vary, ehttputil.RequestStandards.Headers.Accept.Encoding)

				buffer := bytes.NewBuffer([]byte{})
				flateWriter, _ := flate.NewWriter(buffer, flate.DefaultCompression)

				if _, err := ctx.Response.Body.WriteTo(flateWriter); err != nil {
					ctx.Error = errors.AddInformation(errors.SomeThingIsWrong, map[string]interface{}{"ExtraInfo": err})
					return
				}

				flateWriter.Close()
				ctx.Response.Body = buffer
				return
			}
		}
	}
}
