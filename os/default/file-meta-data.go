/* For license and copyright information please see LEGAL file in repository */

package dos

import (
	etime "../../earth-time"
	"../../file"
	"../../protocol"
)

// fileMetaData implement protocol.FilemetaData interface
type fileMetaData struct {
	uri            file.URI
	size           uint64
	creationTime   etime.Time
	lastAccessTime etime.Time
	lastWriteTime  etime.Time
}

func (md *fileMetaData) URI() protocol.FileURI   { return &md.uri }
func (md *fileMetaData) Size() uint64            { return md.size }
func (md *fileMetaData) Created() protocol.Time  { return md.creationTime }
func (md *fileMetaData) Accessed() protocol.Time { return md.lastAccessTime }
func (md *fileMetaData) Modified() protocol.Time { return md.lastWriteTime }
func (md *fileMetaData) Hidden() bool            { return md.uri.Name()[0] == '.' }

// fileDirectoryMetaData implement protocol.FilemetaData interface
type fileDirectoryMetaData struct {
	dirNum  uint
	fileNum uint
	fileMetaData
}

func (md *fileDirectoryMetaData) DirNum() uint  { return md.dirNum }
func (md *fileDirectoryMetaData) FileNum() uint { return md.fileNum }
