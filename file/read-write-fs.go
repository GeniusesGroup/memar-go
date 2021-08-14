/* For license and copyright information please see LEGAL file in repository */

package file

import (
	"io/ioutil"
	"mime"
	"os"
	"path"
	"sort"
)

// ReadRepositoryFromFileSystem use to get all repository by its name!
func (f *Folder) ReadRepositoryFromFileSystem(dirname string, readHidden bool) (err error) {
	var repoFiles []os.FileInfo
	repoFiles, err = ioutil.ReadDir(dirname)
	if err != nil {
		return
	}

	for _, file := range repoFiles {
		var name = file.Name()
		if file.IsDir() {
			if !readHidden && name[0] == '.' {
				continue
			}
			var innerRepo = NewFolder(name)
			innerRepo.Dep = f
			innerRepo.FSPath = path.Join(dirname, name)
			err = innerRepo.ReadRepositoryFromFileSystem(innerRepo.FSPath, readHidden)
			if err != nil {
				return
			}
			f.Dependencies[name] = innerRepo
		} else {
			var data []byte
			data, err = os.ReadFile(path.Join(dirname, name))
			if err != nil {
				return
			}

			var fi = File{
				FullName: name,
				Dep:      f,
				Data:     data,
			}
			for i := len(fi.FullName) - 1; i >= 0; i-- {
				if fi.FullName[i] == '.' {
					fi.Name = fi.FullName[:i]
					fi.Extension = fi.FullName[i+1:]
					fi.MimeType = mime.TypeByExtension(fi.FullName[i:])
				}
			}

			f.SetFile(&fi)
		}
	}

	return
}

// WriteRepositoryToFileSystem use to write repository to file system!
// It print any error to screen and pass last error to caller!
func (f *Folder) WriteRepositoryToFileSystem(dirname string) (err error) {
	for _, obj := range f.Files {
		// Just write changed file
		if obj.State > 0 {
			// Indicate state to not change to don't overwrite it again!
			obj.State = StateUnChanged
			err = ioutil.WriteFile(path.Join(dirname, obj.FullName), obj.Data, 0700)
			if err != nil {
				return
			}
		}
	}

	for _, dep := range f.Dependencies {
		// Just write folder if its not exist!
		if dep.State > 0 {
			// Indicate state to not change to don't overwrite it again!
			dep.State = StateUnChanged
			err = os.Mkdir(path.Join(dirname, dep.Name), 0700)
			if err != nil {
				return
			}
		}
		err = dep.WriteRepositoryToFileSystem(path.Join(dirname, dep.Name))
		if err != nil {
			return
		}
	}

	return
}

// ByModTime use to
type ByModTime []os.FileInfo

func (fis ByModTime) Len() int {
	return len(fis)
}

func (fis ByModTime) Swap(i, j int) {
	fis[i], fis[j] = fis[j], fis[i]
}

func (fis ByModTime) Less(i, j int) bool {
	return fis[i].ModTime().Before(fis[j].ModTime())
}

// SortFilesDec sort given slice in dec
func SortFilesDec(repoFiles []os.FileInfo) {
	sort.Sort(ByModTime(repoFiles))
}
