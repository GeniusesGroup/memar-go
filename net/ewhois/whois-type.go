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

// Extended Whois Package
package ewhois

import "time"

// Whois : The standard struct of WHOIS in usersdata.
type Whois struct {
	Details  Details
	Date     Date
	Contacts Contacts
	Transfer struct {
		Code   string
		Status bool //true for lock, false for unlock
	}
}

// Details : The standard struct of WHOIS details in usersdata.
type Details struct {
	Status             bool
	RegistryDomainID   string
	RegistrarContactID string
	Registrar          string
}

// Date : The standard struct of WHOIS date in usersdata.
type Date struct {
	Register   time.Time
	Creation   time.Time
	Updated    time.Time
	Expiration time.Time
}

// Contacts : The standard struct of Contacts
type Contacts struct {
	Owner   Contact
	Admin   Contact
	Tech    Contact
	Finance Contact
}

// Contact : The standard struct of Contact infromation
type Contact struct {
	Organization string
	Name         string
	Email        string
	Phone        string
	Fax          string
	Address      string
}
