/* For license and copyright information please see the LEGAL file in the code repository */

package directory_p

import (
	error_p "memar/process/error/protocol"
)

type Method_File_GUI interface {
	// make invisible by move to recycle bin
	MoveToRecycleBin() (err error_p.Error)
	
	ParentDirectory() (dir Directory, err error_p.Error)
}
