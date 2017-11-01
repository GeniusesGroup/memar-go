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

package usersdata

import (
	"database/sql"
	"time"

	// For layers architecture we have to add MySQL package in here.
	_ "github.com/go-sql-driver/mysql"

	"github.com/SabzCity/go-library/cryptography/aes"
	"github.com/SabzCity/go-library/errors"
	"github.com/SabzCity/go-library/log"
	"github.com/SabzCity/go-library/microservice"
)

// MySQLPool : Public MySQL database variable for this package.
var MySQLPool *sql.DB

// OpenConnection : Open connection to database with custom driver and ping it.
func OpenConnection() {
	var err error // !!!!!!!!!!!!!!!!! Very Important!!! Don't remove it!!!!!!!!!!!!!!!

	UserName := microservice.MSDetails.OpenAPI.Info.Title
	Password := microservice.MSDetails.OpenAPI.Info.Title
	// IP := ""
	// Port := ""
	DBName := microservice.MSDetails.OpenAPI.Info.Title

	// make password with UserName & microservice.Key in production
	if microservice.MSDetails.Production {
		buffer, err := aes.EncryptAES([]byte(UserName), microservice.Key(time.Now()))
		if err != nil {
			log.Fatal(errors.AddInformation(errors.DatabaseConnectionError, map[string]interface{}{"ExtraInfo": err}))
		}
		Password = string(buffer)
	}

	// Ready connection
	// Connection := UserName + ":" + Password + "@" + IP + ":" + Port + "/" + DBName + "?parseTime=true&multiStatements=true"
	connection := UserName + ":" + Password + "@" + "/" + DBName + "?parseTime=true&multiStatements=true"

	// Connect to the database.
	MySQLPool, err = sql.Open("mysql", connection)
	if err != nil {
		log.Fatal(errors.AddInformation(errors.DatabaseConnectionError, map[string]interface{}{"ExtraInfo": err}))
	}

	err = MySQLPool.Ping()
	if err != nil {
		log.Fatal(errors.AddInformation(errors.DatabasePingOut, map[string]interface{}{"ExtraInfo": err}))
	}
}
