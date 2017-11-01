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

package objectsdata

import (
	"io/ioutil"
	"os"

	"github.com/SabzCity/go-library/errors"
)

// FilesPath is default store location
var FilesPath = "/root/storage/"

func init() {
	//Get ZoneID from OS default File
	// /settings/
	// /objects/
}

// TODO : Add lock state for Delete operate.
// TODO : Add Metadata and ... at the begin of object.
// TODO : find better algorithm for saving data.

// SaveObject save the data of special object in storage.
func SaveObject(UUID string, data []byte) error {
	if err := ioutil.WriteFile(FilesPath+UUID, data, os.ModePerm); err != nil {
		// TODO : Choose a better error.
		return errors.AddInformation(errors.SomeThingIsWrong, map[string]interface{}{"ExtraInfo": err})
	}

	return nil
}

// LoadObject read the data of special object in storage.
func LoadObject(UUID string) ([]byte, error) {
	data, err := ioutil.ReadFile(FilesPath + UUID)
	if err != nil {
		return nil, errors.AddInformation(errors.ContentNotExist, map[string]interface{}{"ExtraInfo": err})
	}

	return data, nil
}

// DeleteObject delete the object from storage.
func DeleteObject(UUID string) error {
	err := os.Remove(FilesPath + UUID)
	if err != nil {
		return errors.AddInformation(errors.ContentNotExist, map[string]interface{}{"ExtraInfo": err})
	}

	return nil
}
