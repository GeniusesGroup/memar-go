/* For license and copyright information please see LEGAL file in repository */

package parser

import (
	"io/ioutil"
	"path"
	"strconv"
	"strings"
)

// ReadRepositoryFromFileSystem use to read or update and write to repo variable!
func (repo *Repository) ReadRepositoryFromFileSystem(dirname string) (err error) {
	repoFiles, err := ioutil.ReadDir(dirname)
	if err != nil {
		return err
	}

	for _, file := range repoFiles {
		if file.IsDir() {
			// Don't read or parse ChaparKhane || git folder!
			if file.Name() == "libgo" || file.Name() == ".git" {
				continue
			}

			var innerRepo = NewRepository()
			err = innerRepo.ReadRepositoryFromFileSystem(path.Join(dirname, file.Name()))
			if err != nil {
				return err
			}
			innerRepo.Name = file.Name()
			innerRepo.FSPath = file.Name()
			repo.Dependencies[file.Name()] = innerRepo
		} else {
			data, err := ioutil.ReadFile(path.Join(dirname, file.Name()))
			if err != nil {
				return err
			}
			folderNameSplit := strings.Split(dirname, "-")
			FolderID, err := strconv.ParseUint(folderNameSplit[len(folderNameSplit)-1], 10, 16)
			repo.Parse(file.Name(), data, uint16(FolderID))
		}
	}

	return nil
}
