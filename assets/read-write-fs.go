/* For license and copyright information please see LEGAL file in repository */

package assets

import (
	"io/ioutil"
	"mime"
	"os"
	"path"
)

// ReadRepositoryFromFileSystem use to get all repository by its name!
func (f *Folder) ReadRepositoryFromFileSystem(dirname string, readHidden bool) (err error) {
	var repoFiles []os.FileInfo
	repoFiles, err = ioutil.ReadDir(dirname)
	if err != nil {
		return err
	}

	for _, file := range repoFiles {
		var name = file.Name()
		if !readHidden && name[0] == '.' {
			continue
		}
		if file.IsDir() {
			var innerRepo = NewFolder(name)
			innerRepo.FSPath = path.Join(dirname, name)
			err = innerRepo.ReadRepositoryFromFileSystem(innerRepo.FSPath, readHidden)
			if err != nil {
				return err
			}
			f.Dependencies[name] = innerRepo
		} else {
			var data []byte
			data, err = ioutil.ReadFile(path.Join(dirname, name))
			if err != nil {
				return err
			}

			var fi = File{
				FullName:   name,
				Dep:        f,
				Data:       data,
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
			err = ioutil.WriteFile(path.Join(dirname, obj.FullName), obj.Data, 0755)
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
			err = os.Mkdir(path.Join(dirname, dep.Name), 0755)
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
