// Copyright 2017 SabzCity. All rights reserver.

package usersdata

import (
	"database/sql"
	"database/sql/driver"
	"strings"
	"time"

	"github.com/SabzCity/go-library/encoding/ejson"
	"github.com/SabzCity/go-library/errors"
)

// MySQL queries.
const (
	getKeysQuery1 = `SELECT UUID FROM {Table} WHERE UUID LIKE ? LIMIT ?, ?;`
	getKeysQuery2 = `SELECT JSON_KEYS(Data, ?) FROM {Table} WHERE UUID = ?;`

	getDataQuery1 = `SELECT Data, LastModified FROM {Table} WHERE UUID = ?;`
	getDataQuery2 = `SELECT JSON_EXTRACT(Data, ?), LastModified FROM {Table} WHERE UUID = ?;`

	addDataQuery1 = `INSERT INTO {Table} (UUID, Data, LastModified) VALUES (?, ?, ?);`
	addDataQuery2 = `UPDATE {Table} SET Data = JSON_INSERT(Data, ?, ?), LastModified = ? WHERE UUID = ?;`

	updateDataQuery1 = `UPDATE {Table} Set Data = ?, LastModified = ? WHERE UUID = ?;`
	updateDataQuery2 = `UPDATE {Table} SET Data = JSON_REPLACE(Data, ?, ?), LastModified = ? WHERE UUID = ?;`

	deleteDataQuery1 = `DELETE FROM {Table} WHERE UUID = ?;`
	deleteDataQuery2 = `UPDATE {Table} SET Data = JSON_REMOVE(Data, ?), LastModified = ? WHERE UUID = ?;`
)

// GetKeys get data keys from user area.
func GetKeys(r *Request) error {
	var (
		tx         *sql.Tx
		rows       *sql.Rows
		row        *sql.Row
		query      string
		jsonKey    string
		returnData string
		err        error
		keys       []string
	)

	if r.Options.Transaction {
		tx, err = MySQLPool.Begin()
		if err != nil {
			return errors.AddInformation(errors.CantPrepareStatement, map[string]interface{}{"ExtraInfo": err})
		}
	}

	switch len(r.Keys) {
	case 0:
		return errors.BadStatement
	case 1:
		return errors.BadStatement
	case 2:
		// Ready query.
		query = strings.Replace(getKeysQuery1, "{Table}", r.Keys[0], -1)

		// Get data from MySQL.
		if r.Options.Transaction {
			rows, err = tx.Query(query, r.Keys[1]+"%", r.Options.Offset, r.Options.Limit)
			if err != nil {
				return errors.AddInformation(errors.CantPrepareStatement, map[string]interface{}{"ExtraInfo": err})
			}
		} else {
			rows, err = MySQLPool.Query(query, r.Keys[1]+"%", r.Options.Offset, r.Options.Limit)
			if err != nil {
				return errors.AddInformation(errors.CantPrepareStatement, map[string]interface{}{"ExtraInfo": err})
			}
		}

		// Read database response.
		for rows.Next() {
			err = rows.Scan(&returnData)
			if err == sql.ErrNoRows {
				return errors.ContentNotExist
			}
			if err != nil {
				return errors.AddInformation(errors.CantPrepareStatement, map[string]interface{}{"ExtraInfo": err})
			}

			keys = append(keys, returnData)
		}
		*r.Data.(*[]string) = keys
	case 3:
		query = strings.Replace(getKeysQuery2, "{Table}", r.Keys[0], -1)

		if r.Options.Transaction {
			row = tx.QueryRow(query, "$", r.Keys[1], r.Options.Offset, r.Options.Limit)
		} else {
			row = MySQLPool.QueryRow(query, "$", r.Keys[1], r.Options.Offset, r.Options.Limit)
		}

		err = row.Scan(&returnData)
		if err == sql.ErrNoRows {
			return errors.ContentNotExist
		}
		if err != nil {
			return errors.AddInformation(errors.CantPrepareStatement, map[string]interface{}{"ExtraInfo": err})
		}

		// Convert from stored type to struct and write to r.Data.
		err = ejson.Unmarshal(returnData, r.Data)
		if err != nil {
			return errors.StoredDataConflict
		}
	default:
		query = strings.Replace(getKeysQuery2, "{Table}", r.Keys[0], -1)
		jsonKey = strings.Join(r.Keys[2:], ".") // 2 means third value of r.Keys becuase we declare first and second before!

		if r.Options.Transaction {
			row = tx.QueryRow(query, "$."+jsonKey, r.Keys[1], r.Options.Offset, r.Options.Limit)
		} else {
			row = MySQLPool.QueryRow(query, "$."+jsonKey, r.Keys[1], r.Options.Offset, r.Options.Limit)
		}

		err = row.Scan(&returnData)
		if err == sql.ErrNoRows {
			return errors.ContentNotExist
		}
		if err != nil {
			return errors.AddInformation(errors.CantPrepareStatement, map[string]interface{}{"ExtraInfo": err})
		}

		// Convert from stored type to struct and write to r.Data
		err = ejson.Unmarshal(returnData, r.Data)
		if err != nil {
			return errors.StoredDataConflict
		}
	}

	return nil
}

// ChangeOwner : Change owner of a record.
func ChangeOwner(oldOwner *Request, newOwner *Request) error {
	return nil
}

// GetData : Get data from user area.
func GetData(r *Request) error {
	var (
		row        *sql.Row
		query      string
		jsonKey    string
		returnData string
		err        error
	)

	switch len(r.Keys) {
	case 0:
		return errors.BadStatement
	case 1:
		return errors.BadStatement
	case 2:
		query = strings.Replace(getDataQuery1, "{Table}", r.Keys[0], -1)

		// Get data from MySQL.
		row = MySQLPool.QueryRow(query, r.Keys[1])

		err = row.Scan(&returnData, &r.LastModified)
		if err == sql.ErrNoRows {
			return errors.ContentNotExist
		}
		if err != nil {
			return errors.AddInformation(errors.CantPrepareStatement, map[string]interface{}{"ExtraInfo": err})
		}
	case 3:
		query = strings.Replace(getDataQuery1, "{Table}", r.Keys[0], -1)

		row = MySQLPool.QueryRow(query, r.Keys[1])

		err = row.Scan(&returnData, &r.LastModified)
		if err == sql.ErrNoRows {
			return errors.ContentNotExist
		}
		if err != nil {
			return errors.AddInformation(errors.CantPrepareStatement, map[string]interface{}{"ExtraInfo": err})
		}
	default:
		query = strings.Replace(getDataQuery2, "{Table}", r.Keys[0], -1)
		jsonKey = strings.Join(r.Keys[2:], ".") // 2 means third value of r.Keys becuase we declare first and second before!

		row = MySQLPool.QueryRow(query, "$."+jsonKey, r.Keys[1])

		err = row.Scan(&returnData, &r.LastModified)
		if err == sql.ErrNoRows {
			return errors.ContentNotExist
		}
		if err != nil {
			return errors.AddInformation(errors.CantPrepareStatement, map[string]interface{}{"ExtraInfo": err})
		}
	}

	// Convert from stored type to struct and write to r.Data.
	err = ejson.Unmarshal(returnData, r.Data)
	if err != nil {
		return errors.StoredDataConflict
	}

	return nil
}

// AddData : Add data to specify user area.
func AddData(r *Request) error {
	var (
		result       sql.Result
		query        string
		jsonKey      string
		rowsAffected int64
		err          error
		// lastInsertID int64
	)

	data, err := ejson.Marshal(r.Data)
	if err != nil {
		return errors.StoringDataConflict
	}

	switch len(r.Keys) {
	case 0:
		return errors.BadStatement
	case 1:
		return errors.BadStatement
	case 2:
		query = strings.Replace(addDataQuery1, "{Table}", r.Keys[0], -1)

		result, err = MySQLPool.Exec(query, r.Keys[1], data, time.Now().UTC().UnixNano())
		if err == driver.ErrBadConn {
			return errors.CantPrepareStatement
		} else if err != nil { // TODO : Check This case for other scenario!
			return errors.ContentAlreadyExist
		}

		rowsAffected, err = result.RowsAffected()
		if err != nil {
			// Check error
		}
		if rowsAffected < 1 {
			return errors.StoringDataNotComplete
		}
		// lastInsertID, err = result.LastInsertId()
		// if err != nil {
		//   // Check error
		// }
		// Check lastInsertID if need!!
	case 3:
		query = strings.Replace(addDataQuery1, "{Table}", r.Keys[0], -1)

		result, err = MySQLPool.Exec(query, r.Keys[1], data, time.Now().UTC().UnixNano())
		if err == driver.ErrBadConn {
			return errors.CantPrepareStatement
		} else if err != nil { // TODO : Check This case for other scenario!
			return errors.ContentAlreadyExist
		}

		rowsAffected, err = result.RowsAffected()
		if err != nil {
			// Check error
		}
		if rowsAffected < 1 {
			return errors.StoringDataNotComplete
		}
		// lastInsertID, err = result.LastInsertId()
		// if err != nil {
		//   // Check error
		// }
		// Check lastInsertID if need!!
	default:
		query = strings.Replace(addDataQuery2, "{Table}", r.Keys[0], -1)
		jsonKey = strings.Join(r.Keys[2:], ".") // 2 means third value of r.Keys becuase we declare first and second before!

		result, err = MySQLPool.Exec(query, "$."+jsonKey, data, time.Now().UTC().UnixNano(), r.Keys[1])
		if err == driver.ErrBadConn {
			return errors.CantPrepareStatement
		} else if err != nil { // TODO : Check This case for other scenario!
			return errors.ContentAlreadyExist
		}

		rowsAffected, err = result.RowsAffected()
		if err != nil {
			// Check error
		}
		if rowsAffected < 1 {
			return errors.StoringDataNotComplete
		}
		// lastInsertID, err = result.LastInsertId()
		// if err != nil {
		//   Check error
		// }
		// Check lastInsertID if need!!
	}

	return nil
}

// UpdateData : update data to user area
func UpdateData(r *Request) error {
	var (
		result       sql.Result
		query        string
		jsonKey      string
		rowsAffected int64
		err          error
	)

	data, err := ejson.Marshal(r.Data)
	if err != nil {
		return errors.StoringDataConflict
	}

	switch len(r.Keys) {
	case 0:
		return errors.BadStatement
	case 1:
		return errors.BadStatement
	case 2:
		query = strings.Replace(updateDataQuery1, "{Table}", r.Keys[0], -1)

		result, err = MySQLPool.Exec(query, data, time.Now().UTC().UnixNano(), r.Keys[1])
		if err != nil {
			return errors.CantPrepareStatement // TODO : Check This error for other scenario!
		}

		rowsAffected, err = result.RowsAffected()
		if err != nil {
			// Check error
		}
		if rowsAffected < 1 {
			if r.Options.CreateOnUpdate {
				query = strings.Replace(addDataQuery1, "{Table}", r.Keys[0], -1)

				result, err = MySQLPool.Exec(query, r.Keys[1], data, time.Now().UTC().UnixNano())
				if err == driver.ErrBadConn {
					return errors.CantPrepareStatement
				} else if err != nil { //TODO : Check This case for other scenario!
					return errors.ContentAlreadyExist
				}

				rowsAffected, err = result.RowsAffected()
				if err != nil {
					// Check error
				}
				if rowsAffected < 1 {
					return errors.StoringDataNotComplete
				}
			} else {
				return errors.ContentNotExist
			}
		}
	case 3:
		query = strings.Replace(updateDataQuery1, "{Table}", r.Keys[0], -1)

		result, err = MySQLPool.Exec(query, data, time.Now().UTC().UnixNano(), r.Keys[1])
		if err != nil {
			return errors.CantPrepareStatement //TODO : Check This error for other scenario!
		}

		rowsAffected, err = result.RowsAffected()
		if err != nil {
			// Check error
		}
		if rowsAffected < 1 {
			if r.Options.CreateOnUpdate {
				query = strings.Replace(addDataQuery1, "{Table}", r.Keys[0], -1)

				result, err = MySQLPool.Exec(query, r.Keys[1], data, time.Now().UTC().UnixNano())
				if err == driver.ErrBadConn {
					return errors.CantPrepareStatement
				} else if err != nil { // TODO : Check This case for other scenario!
					return errors.ContentAlreadyExist
				}

				rowsAffected, err = result.RowsAffected()
				if err != nil {
					// Check error
				}
				if rowsAffected < 1 {
					return errors.StoringDataNotComplete
				}
			} else {
				return errors.ContentNotExist
			}
		}
	default:
		query = strings.Replace(updateDataQuery2, "{Table}", r.Keys[0], -1)
		jsonKey = strings.Join(r.Keys[2:], ".") // 2 means third value of r.Keys becuase we declare first and second before!

		result, err = MySQLPool.Exec(query, "$."+jsonKey, data, time.Now().UTC().UnixNano(), r.Keys[1])
		if err != nil {
			return errors.CantPrepareStatement // TODO : Check This error for other scenario!
		}

		rowsAffected, err = result.RowsAffected()
		if err != nil {
			// Check error
		}
		if rowsAffected < 1 {
			if r.Options.CreateOnUpdate {
				query = strings.Replace(addDataQuery2, "{Table}", r.Keys[0], -1)

				result, err = MySQLPool.Exec(query, "$."+jsonKey, data, time.Now().UTC().UnixNano(), r.Keys[1])
				if err == driver.ErrBadConn {
					return errors.CantPrepareStatement
				} else if err != nil { // TODO : Check This case for other scenario!
					return errors.ContentAlreadyExist
				}

				rowsAffected, err = result.RowsAffected()
				if err != nil {
					//Check error
				}
				if rowsAffected < 1 {
					return errors.StoringDataNotComplete
				}
			} else {
				return errors.ContentNotExist
			}
		}
	}

	return nil
}

// DeleteData : delete data from user area
// If last string in Keys empty means delete all data with key Like one to end!
func DeleteData(r *Request) error {
	var (
		result       sql.Result
		query        string
		jsonKey      string
		rowsAffected int64
		err          error
	)

	switch len(r.Keys) {
	case 0:
		return errors.BadStatement
	case 1:
		return errors.BadStatement
	case 2:
		query = strings.Replace(deleteDataQuery1, "{Table}", r.Keys[0], -1)
		result, err = MySQLPool.Exec(query, r.Keys[1])
		if err != nil {
			return errors.CantPrepareStatement // TODO : Check This error for other scenario!
		}

		rowsAffected, err = result.RowsAffected()
		if err != nil {
			// Check error
		}
		if rowsAffected < 1 {
			return errors.ContentNotExist
		}
	case 3:
		query = strings.Replace(deleteDataQuery1, "{Table}", r.Keys[0], -1)
		result, err = MySQLPool.Exec(query, r.Keys[1])
		if err != nil {
			return errors.CantPrepareStatement // TODO : Check This error for other scenario!
		}

		rowsAffected, err = result.RowsAffected()
		if err != nil {
			// Check error
		}
		if rowsAffected < 1 {
			return errors.ContentNotExist
		}
	default:
		query = strings.Replace(deleteDataQuery2, "{Table}", r.Keys[0], -1)
		jsonKey = strings.Join(r.Keys[2:], ".") // 2 means third value of r.Keys becuase we declare first and second before!

		result, err = MySQLPool.Exec(query, "$."+jsonKey, time.Now().UTC().UnixNano(), r.Keys[1])
		if err != nil {
			return errors.CantPrepareStatement // TODO : Check This error for other scenario!
		}

		rowsAffected, err = result.RowsAffected()
		if err != nil {
			// Check error
		}
		if rowsAffected < 1 {
			return errors.ContentNotExist
		}
	}

	return nil
}
