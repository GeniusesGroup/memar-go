/* For license and copyright information please see LEGAL file in repository */

/*
dos abbreviations for "Default Operating System"
*/
package dos

import (
	goos "os"
	"sort"
	"strings"

	"../../file"
	"../../protocol"
)

// FileDirectory use to store app needed data from repo like html, css, js, ...
type FileDirectory struct {
	metadata        fileDirectoryMetaData
	parentDirectory *FileDirectory            // Not nill if the directory has parent!
	directories     map[string]*FileDirectory // map key is dir name
	files           map[string]*File          // map key is file name
}

func (dir *FileDirectory) MetaData() protocol.FileDirectoryMetaData { return &dir.metadata }
func (dir *FileDirectory) ParentDirectory() protocol.FileDirectory  { return dir.parentDirectory }

// Mkdir make new directory in relative path from given directory.
func (dir *FileDirectory) Directories(offset, limit uint64) (dirs []protocol.FileDirectory) {
	// TODO:::
	return
}

// Directory return the directory by its name or make new one if desire name not exist.
func (dir *FileDirectory) Directory(name string) (dr protocol.FileDirectory, err protocol.Error) {
	var goErr error
	var exist bool
	var dirPath = dir.metadata.uri.Path() + "/" + name + "/"
	dr, exist = dir.directories[name]
	if !exist {
		goErr = goos.Mkdir(dirPath, 0700)
		if goErr != nil {
			// err =
			return
		}
	}
	if dr == nil {
		var dir = FileDirectory{
			parentDirectory: dir,
		}
		err = dir.init(dirPath)
		if err != nil {
			return
		}
		dir.directories[name] = &dir
	}
	return
}

// Files use to get all files in order by name.
func (dir *FileDirectory) Files(offset, limit uint64) (files []protocol.File) {
	var keys = make([]string, 0, len(dir.files))
	for k := range dir.files {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	files = make([]protocol.File, len(dir.files))
	for i, k := range keys {
		files[i] = dir.files[k]
	}
	return
}

// File use to get a file by its full name with extension
// And make new one if desire name not exist.
func (dir *FileDirectory) File(name string) (file protocol.File, err protocol.Error) {
	file, _ = dir.files[name]
	if file == nil {
		// make new
		var fi = File{
			parentDirectory: dir,
		}
		err = fi.init(name)
		if err != nil {
			return
		}
		dir.files[name] = &fi
	}
	return
}

// FileByPath use to get a file by its path in the directory
func (dir *FileDirectory) FileByPath(uriPath string) (file protocol.File, err protocol.Error) {
	file, err = dir.fileByPath(uriPath)
	return
}

func (dir *FileDirectory) fileByPath(uriPath string) (file *File, err protocol.Error) {
	var pathParts = strings.Split(uriPath, "/")
	var pathPartsLen = len(pathParts)
	var fullFileName = pathParts[pathPartsLen-1]

	if pathPartsLen < 3 {
		file = dir.files[fullFileName]
		return
	}

	if pathParts[0] == "" {
		pathParts = pathParts[1:]
	}
	var directory *FileDirectory
	directory, err = dir.directoryByPathParts(pathParts)
	if directory == nil || err != nil {
		return
	}
	file = directory.files[fullFileName]
	return
}

// DirectoryByPath use to get a directory by its full path location.
// path is the file location in FileSystems include file name
func (dir *FileDirectory) DirectoryByPath(pathParts []string) (directory protocol.FileDirectory, err protocol.Error) {
	directory, err = dir.directoryByPathParts(pathParts)
	return
}

func (dir *FileDirectory) directoryByPath(uriPath string) (directory *FileDirectory, err protocol.Error) {
	var pathParts = strings.Split(uriPath, "/")
	if pathParts[0] == "" {
		pathParts = pathParts[1:]
	}
	directory, err = dir.directoryByPathParts(pathParts)
	return
}

func (dir *FileDirectory) directoryByPathParts(pathParts []string) (directory *FileDirectory, err protocol.Error) {
	var pathPartsLen = len(pathParts)
	directory = dir
	for i := 0; i < pathPartsLen; i++ {
		directory, _ = directory.directories[pathParts[i]]
		if directory == nil {
			return
		}
	}
	return
}

// FindFiles use to get a file by some part of its name!
func (dir *FileDirectory) FindFiles(partName string, num uint) (files []protocol.File) {
	for fileName, file := range dir.files {
		if strings.Contains(fileName, partName) {
			files = append(files, file)
			if len(files) == int(num) {
				return
			}
		}
	}
	return
}

// FindFiles use to get a file by some part of its name!
func (dir *FileDirectory) FindFile(partName string) (files protocol.File) {
	for fileName, file := range dir.files {
		if strings.Contains(fileName, partName) {
			return file
		}
	}
	return
}

// FindFileRecursively use to get a file by its ful name with extension in recursively!
func (dir *FileDirectory) FindFileRecursively(partName string) (file *File) {
	file = dir.files[partName]
	if file != nil {
		return
	}
	if dir.directories != nil {
		for _, dep := range dir.directories {
			file = dep.FindFileRecursively(partName)
			if file != nil {
				return
			}
		}
	}
	return nil
}

func (dir *FileDirectory) Rename(oldURIPath, newURIPath string) (err protocol.Error) {
	// TODO:::
	return
}

func (dir *FileDirectory) Copy(uriPath, newURIPath string) (err protocol.Error) {
	// var file = File{
	// 	Directory: f.Directory,
	// 	Path:      f.Path,
	// 	FullName:  f.FullName,
	// 	Name:      f.Name,
	// 	Extension: f.Extension,
	// 	mediaType: f.mediaType,
	// 	State:     f.State,
	// 	data:      f.data,
	// }
	return
}

func (dir *FileDirectory) Move(uriPath, newURIPath string) (err protocol.Error) {
	// var file = File{
	// 	Directory: f.Directory,
	// 	Path:      f.Path,
	// 	FullName:  f.FullName,
	// 	Name:      f.Name,
	// 	Extension: f.Extension,
	// 	mediaType: f.mediaType,
	// 	State:     f.State,
	// 	data:      make([]byte, len(f.data)),
	// }
	// copy(file.data, f.data)
	return
}

func (dir *FileDirectory) Delete(uriPath string) (err protocol.Error) {
	goos.Remove(uriPath)
	if file.IsPathDirectory(uriPath) {
		var dir, _ = dir.directoryByPath(uriPath)
		delete(dir.parentDirectory.directories, dir.metadata.URI().Name())
	} else {
		var file, _ = dir.fileByPath(uriPath)
		delete(file.parentDirectory.files, file.metadata.URI().Name())
	}
	return
}

func (dir *FileDirectory) Wipe(uriPath string) (err protocol.Error) {
	// TODO::: first write null data to uri path
	err = dir.Delete(uriPath)
	return
}

// GetDependencyRecursively use to get a dependency by its name in recursively!
func (dir *FileDirectory) GetDependencyRecursively(name string) *FileDirectory {
	var t *FileDirectory
	var ok bool
	if dir.directories != nil {
		t, ok = dir.directories[name]
		if !ok {
			for _, dep := range dir.directories {
				t = dep.GetDependencyRecursively(name)
				if t != nil {
					break
				}
			}
		}
	}
	return t
}

// init use to read||update Directory from disk.
func (dir *FileDirectory) init(path string) (err protocol.Error) {
	dir.metadata.URI().Init(path)
	dir.files = make(map[string]*File)
	dir.directories = make(map[string]*FileDirectory)

	var dirEntry, goErr = goos.ReadDir(dir.metadata.uri.Path())
	if goErr != nil {
		// err =
		return
	}
	for _, file := range dirEntry {
		var fileName = file.Name()
		if file.IsDir() {
			dir.directories[fileName] = nil
		} else {
			dir.files[fileName] = nil
		}
	}
	return
}
