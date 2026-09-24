/* For license and copyright information please see the LEGAL file in the code repository */

package record_id_p

import (
	memar_uuid "memar/identifier/uuid-memar"
)

type Field_RecordUUID interface {
	RecordUUID() UUID
}

type Field_RecordsUUID interface {
	RecordsUUID() []UUID
}

// In a relational database, a row or "record" or "tuple", represents a single, implicitly structured data item in a table.
// https://en.wikipedia.org/wiki/Row_(database)
type UUID memar_uuid.Time16
