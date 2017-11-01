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

// Package ehttputil is Extended HTTP package utility
package ehttputil

// CompressionTypes : All standard Compress types
var CompressionTypes = struct {
	Compress    string
	Deflate     string
	Exi         string
	Gzip        string
	Identity    string
	Pack200Gzip string
	Brotli      string
}{
	Compress:    "compress",
	Deflate:     "deflate",
	Exi:         "exi",
	Gzip:        "gzip",
	Identity:    "identity",
	Pack200Gzip: "pack200-gzip",
	Brotli:      "br"}
