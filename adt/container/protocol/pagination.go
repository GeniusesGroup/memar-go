/* For license and copyright information please see the LEGAL file in the code repository */

package container_p

type Offset = ElementIndex

type Field_Offset interface {
	Offset() Offset
}

type Limit = NumberOfElement

type Field_Limit interface {
	Limit() Limit
}

type Pagination_Offset_Request interface {
	Field_Offset
	Field_Limit
}

type Pagination_Cursor_Request interface {
	Location()
	Field_Limit
}

type Pagination_Response interface {
	TotalNumber() NumberOfElement
}
