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

package code

import (
	"strconv"
	"time"

	"strings"
)

const (
	// DefaultSinceYear : Default begin of time in TimeUUID().
	DefaultSinceYear int = 2016

	// SecondUUID : One of TimeUUID() options.
	SecondUUID byte = 1

	// NanoSecondUUID : One of TimeUUID() options.
	NanoSecondUUID byte = 2
)

// TimeUUID : Generate a UUID with time in special format.
func TimeUUID(length int, option byte, sinceyear int) string {

	// Return a second base UUID.
	switch option {
	case SecondUUID:
		// Get the number of seconds since the year.
		seconds := strconv.FormatInt((int64)(time.Since(time.Date(sinceyear, time.January, 1, 0, 0, 0, 0, time.UTC)).Seconds()), 10)

		if length-strings.Count(seconds, "")+1 <= 0 {
			// Add some zero at left.
			return strings.Repeat("0", length-strings.Count(seconds, "")+1) + seconds
		}

		return seconds

	case NanoSecondUUID:
		// Get the number of seconds since the year.
		nanoSeconds := strconv.FormatInt(time.Since(time.Date(sinceyear, time.January, 1, 0, 0, 0, 0, time.UTC)).Nanoseconds(), 10)

		if length-strings.Count(nanoSeconds, "")+1 <= 0 {
			// Add some zero at left.
			return strings.Repeat("0", length-strings.Count(nanoSeconds, "")+1) + nanoSeconds
		}

		return nanoSeconds
	}

	return ""
}
