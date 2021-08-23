/* For license and copyright information please see LEGAL file in repository */

package file

import (
	"io/fs"
	"os"
	"path"
	"sort"
	"strings"
)

// Folder use to store app needed data from repo like html, css, js, ...
type Folder struct {
	Name       string
	Path       string     // Folder location in FileSystems include folder name
	Event      chan uint8 // use in dev phase to update Folder if any change occur
	State      uint8
	readHidden bool
	Folder     *Folder            // Not nill if Folder belong to other folder!
	Folders    map[string]*Folder // Name
	Files      map[string]*File   // Name
}

// NewFolder make new Folder object
func NewFolder(name, path string, readHidden bool) *Folder {
	return &Folder{
		Name:       name,
		Path:       path,
		readHidden: readHidden,
		Files:      make(map[string]*File),
		Folders:    make(map[string]*Folder),
	}
}

// Init initialize the new Folder object
func (f *Folder) Init(name, path string, readHidden bool) {
	f.Name = name
	f.Path = path
	f.readHidden = readHidden
	f.Files = make(map[string]*File)
	f.Folders = make(map[string]*Folder)
}

// NewFolder make new Folder object in given Folder as dependency.
func (f *Folder) NewFolder(name, path string, readHidden bool) {
	f.Folders[name] = &Folder{
		Name:       name,
		Path:       path,
		readHidden: readHidden,
		Files:      make(map[string]*File),
		Folders:    make(map[string]*Folder),
	}
}

// CopyFrom copy from given folder to desire folder.
func (f *Folder) CopyFrom(folder *Folder) {
	f.Name = folder.Name
	f.Path = folder.Path
	f.State = folder.State
	f.readHidden = folder.readHidden
	f.Folder = folder.Folder
	f.Folders = folder.Folders
	f.Files = folder.Files
}

// GetFiles use to get all files in order by name.
func (f *Folder) GetFiles() (files []*File) {
	var keys = make([]string, 0, len(f.Files))
	for k := range f.Files {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	files = make([]*File, len(f.Files))
	for i, k := range keys {
		files[i] = f.Files[k]
	}
	return
}

// GetFile use to get a file by its full name with extension!
func (f *Folder) GetFile(fullName string) *File {
	return f.Files[fullName]
}

// GetFile use to get a file by its full name with extension!
// path is the file location in FileSystems include file name
func (f *Folder) GetFileByPath(path string) *File {
	var pathParts = strings.Split(path, "/")
	var pathPartsLen = len(pathParts)
	var fullFileName = pathParts[pathPartsLen-1]

	if pathPartsLen < 3 {
		return f.GetFile(fullFileName)
	}
	var desireFolder = f.GetFolderByPath(pathParts)
	if desireFolder == nil {
		return nil
	}
	return desireFolder.GetFile(fullFileName)
}

// GetFolderByPath use to get a folder by its full path location.
// path is the file location in FileSystems include file name
func (f *Folder) GetFolderByPath(pathParts []string) (desireFolder *Folder) {
	var pathPartsLen = len(pathParts)
	for i := 0; i < pathPartsLen; i++ {
		if pathParts[i] == "" {
			continue
		}
		var folder = f.GetDependency(pathParts[i])
		if folder == nil {
			return desireFolder
		}
		desireFolder = folder
	}
	return
}

// FindFiles use to get a file by its ful name with extension!
func (f *Folder) FindFiles(fullName string) (files []*File) {
	for _, file := range f.Files {
		if strings.Contains(file.FullName, fullName) {
			files = append(files, file)
		}
	}
	return
}

// FindFileRecursively use to get a file by its ful name with extension in recursively!
func (f *Folder) FindFileRecursively(fullName string) (file *File) {
	file = f.Files[fullName]
	if file != nil {
		return
	}
	if f.Folders != nil {
		for _, dep := range f.Folders {
			file = dep.FindFileRecursively(fullName)
			if file != nil {
				return
			}
		}
	}
	return nil
}

// SetFile use to set a file to given asset!
func (f *Folder) SetFile(file *File) {
	f.Files[file.FullName] = file
}

// SetFiles use to set files to given asset!
func (f *Folder) SetFiles(files ...*File) {
	for _, file := range files {
		f.Files[file.FullName] = file
	}
}

// MinifyCompressSet minify & compress file and set it to given asset. Mostly to serve file by servers.
func (f *Folder) MinifyCompressSet(file *File, compressType string) {
	file.Minify()
	file.Compress(compressType)
	f.Files[file.FullName] = file
}

// MinifyCompressSets minify & compress files and set them to given asset. Mostly to serve file by servers.
func (f *Folder) MinifyCompressSets(files []*File, compressType string) {
	for _, file := range files {
		file.Minify()
		file.Compress(compressType)
		f.Files[file.FullName] = file
	}
}

// DeleteFile use to delete the file from given folder!
func (f *Folder) DeleteFile(fullName string) {
	delete(f.Files, fullName)
}

// GetFilesNumbers use to get total files number in given Folder!
func (f *Folder) GetFilesNumbers() int {
	return len(f.Files)
}

// GetDependency use to get a dependency by its name!
func (f *Folder) GetDependency(name string) *Folder {
	return f.Folders[name]
}

// GetDependencyRecursively use to get a dependency by its name in recursively!
func (f *Folder) GetDependencyRecursively(name string) *Folder {
	var t *Folder
	var ok bool
	if f.Folders != nil {
		t, ok = f.Folders[name]
		if !ok {
			for _, dep := range f.Folders {
				t = dep.GetDependencyRecursively(name)
				if t != nil {
					break
				}
			}
		}
	}
	return t
}

// SetDependency use to set a dependency in given folder!
func (f *Folder) SetDependency(dep *Folder) {
	f.Folders[dep.Name] = dep
}

// SetDependencyFlat use to set a repo files and inner dependencies files to the Folder flatly!!
func (f *Folder) SetDependencyFlat(repo *Folder) {
	for _, dep := range repo.Folders {
		f.SetDependencyFlat(dep)
	}
	for _, file := range repo.Files {
		f.SetFile(file)
	}
}

// Update use to read||update Folder from disk or network.
func (f *Folder) Update() (err error) {
	var dirEntry []fs.DirEntry
	dirEntry, err = os.ReadDir(f.Path)
	if err != nil {
		return
	}

	for _, file := range dirEntry {
		var fileName = file.Name()
		var filePath = path.Join(f.Path, fileName)
		if file.IsDir() {
			if !f.readHidden && fileName[0] == '.' {
				continue
			}
			var innerRepo = &Folder{
				Name:       fileName,
				Path:       filePath,
				readHidden: f.readHidden,
				Folder:     f,
			}
			err = innerRepo.Update()
			if err != nil {
				return
			}
			f.SetDependency(innerRepo)
		} else {
			var fi = File{
				Path:   filePath,
				Folder: f,
			}
			fi.RenameFullName(fileName)

			f.SetFile(&fi)
		}
	}

	return
}

// Save use to write repository to file system!
// It print any error to screen and pass last error to caller!
func (f *Folder) Save() (err error) {
	// Just write folder if its not exist!
	if f.State == StateUnChanged {
		return
	}

	// Indicate state to not change to don't overwrite it again!
	f.State = StateUnChanged
	err = os.Mkdir(f.Path, 0700)

	for _, file := range f.Files {
		file.Save()
	}
	for _, dep := range f.Folders {
		err = dep.Save()
		if err != nil {
			return
		}
	}

	return
}
