//Copyright 2017 SabzCity. All rights reserved.
// 424 Failed Dependency - X-Error Related

package errors

import "net/http"

//Declare Errors SabzCity Code
const (
	storedDataConflict = 42400 + (iota + 1)
)

//Declare Errors Detials
var (
	StoredDataConflict = New("Last request may be corrupted stored data. Please retry again and if is failed again send data again", storedDataConflict, http.StatusBadRequest)
)
