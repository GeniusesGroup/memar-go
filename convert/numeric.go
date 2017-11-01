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

package convert

import "strconv"

// StringToUInt64 : Convert a numeric "string" to "uint64".
func StringToUInt64(s string) (uint64, error) {

	return strconv.ParseUint(s, 10, 64)
}

// UInt64ToString : Convert a "uint64" to "string".
func UInt64ToString(i uint64) string {

	return strconv.FormatUint(i, 10)
}

// StringToInt : Convert a numeric "string" to "int".
func StringToInt(s string) (int, error) {

	return strconv.Atoi(s)
}

// IntToString : Convert a "int" to numeric "string".
func IntToString(i int) string {

	return strconv.Itoa(i)
}
