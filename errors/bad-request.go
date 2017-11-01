//Copyright 2017 SabzCity. All rights reserved.
// 400 Bad Request - X-Error Related
// https://tools.ietf.org/html/rfc7231#section-6.5.1

package errors

import (
	"net/http"
)

//Declare Errors SabzCity Code
const (
	bodyNotValid = 40000 + (iota + 1)

	queryNotValid

	contentAlreadyConfirmed

	requestDataIsNotValid

	passwordIsNotCorrect

	cookieIsTooLarge

	productStockIsNotEnough

	productStockMismatch

	itemNotAddable

	invoiceNotCheckoutYet

	invoiceNotChanged
)

//Declare Errors Detials
var (
	BodyNotValid = New("Body of request isn't valid", bodyNotValid, http.StatusBadRequest)

	QueryNotValid = New("Query of URL isn't valid or some of inner values isn't valid value for this request", queryNotValid, http.StatusBadRequest)

	ContentAlreadyConfirmed = New("This content was already confirmed", contentAlreadyConfirmed, http.StatusBadRequest)

	RequestDataIsNotValid = New("Requested data validation failed", requestDataIsNotValid, http.StatusBadRequest)

	PasswordIsNotCorrect = New("Password in request is not correct", passwordIsNotCorrect, http.StatusBadRequest)

	CookieIsTooLarge = New("The cookie with this value is not valid", cookieIsTooLarge, http.StatusBadRequest)

	ProductStockIsNotEnough = New("Product stock isn't enough to register invoice! maybe some one register invoice before you process your invoice", productStockIsNotEnough, http.StatusBadRequest)

	ProductStockMismatch = New("Mismatch in OwnerStock & InputStock & RealStock", productStockMismatch, http.StatusBadRequest)

	ItemNotAddable = New("For each organization products or services must register new invoice. User can add just same organization items to an invoice", itemNotAddable, http.StatusForbidden)

	InvoiceNotCheckoutYet = New("Invoice not checkout yet. for this operation you must checkout it first", invoiceNotCheckoutYet, http.StatusForbidden)

	InvoiceNotChanged = New("Invoice change request is not complete. Maybe Item in Invoice not exist. Please Check your request and send it again", invoiceNotChanged, http.StatusForbidden)
)
