/* For license and copyright information please see LEGAL file in repository */

package assets

import (
	"sort"
	"strings"
)

// Folder use to store app needed data from repo like html, css, js, ...
type Folder struct {
	Name         string
	FSPath       string     // Folder location in FileSystems
	Event        chan uint8 // use in dev phase to update Folder if any change occur!!
	State        uint8
	Files        map[string]*File   // Name
	Dependencies map[string]*Folder // Name
}

// NewFolder make new Folder object
func NewFolder(name string) *Folder {
	return &Folder{
		Name:         name,
		Files:        make(map[string]*File),
		Dependencies: make(map[string]*Folder),
	}
}

// NewFolder make new Folder object in given Folder as dependency.
func (f *Folder) NewFolder(name string) {
	f.Dependencies[name] = &Folder{
		Name:         name,
		Files:        make(map[string]*File),
		Dependencies: make(map[string]*Folder),
	}
}

// GetFiles use to get all files in order by name.
func (f *Folder) GetFiles() (files []*File) {
	var keys = make([]string, 0, len(f.Files))
	for k := range f.Files {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	for _, k := range keys {
		files = append(files, f.Files[k])
	}
	return
}

// GetFile use to get a file by its ful name with extension!
func (f *Folder) GetFile(fullName string) *File {
	return f.Files[fullName]
}

// GetFileRecursively use to get a file by its ful name with extension in recursively!
func (f *Folder) GetFileRecursively(fullName string) (file *File) {
	file = f.Files[fullName]
	if file != nil {
		return
	}
	if f.Dependencies != nil {
		for _, dep := range f.Dependencies {
			file = dep.GetFileRecursively(fullName)
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
func (f *Folder) SetFiles(files []*File) {
	for _, file := range files {
		f.Files[file.FullName] = file
	}
}

// SetAndCompressFile use to compress and set a file to given asset. Mostly to serve file by servers.
func (f *Folder) SetAndCompressFile(file *File, compressType string) {
	file.Compress(compressType)
	f.Files[file.FullName] = file
}

// SetAndCompressFiles use to compress and set files to given asset. Mostly to serve file by servers.
func (f *Folder) SetAndCompressFiles(files []*File, compressType string) {
	for _, file := range files {
		file.Compress(compressType)
		f.Files[file.FullName] = file
	}
}

// FindFiles use to get a file by its ful name with extension!
func (f *Folder) FindFiles(name string) (files []*File) {
	for _, file := range f.Files {
		if strings.Contains(file.FullName, name) {
			files = append(files, file)
		}
	}
	return
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
	return f.Dependencies[name]
}

// GetDependencyRecursively use to get a dependency by its name in recursively!
func (f *Folder) GetDependencyRecursively(name string) *Folder {
	var t *Folder
	var ok bool
	if f.Dependencies != nil {
		t, ok = f.Dependencies[name]
		if !ok {
			for _, dep := range f.Dependencies {
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
	f.Dependencies[dep.Name] = dep
}

// SetDependencyFlat use to set a repo files and inner dependencies files to the Folder flatly!!
func (f *Folder) SetDependencyFlat(repo *Folder) {
	for _, dep := range repo.Dependencies {
		f.SetDependencyFlat(dep)
	}
	for _, file := range repo.Files {
		f.SetFile(file)
	}
}

// UpdateRepo use to update Folder repo that watch from disk or network to any change to the Folder!
func (f *Folder) UpdateRepo(repo *Folder) {

}
