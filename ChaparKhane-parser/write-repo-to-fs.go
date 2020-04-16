/* For license and copyright information please see LEGAL file in repository */

package parser

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
)

// WriteRepositoryToFileSystem use to write repository to file system!
// It print any error to screen and pass last error to caller!
func (repo *Repository) WriteRepositoryToFileSystem(dirname string) (err error) {
	for _, obj := range repo.Files {
		err = ioutil.WriteFile(path.Join(dirname, obj.Name), obj.Data, 0755)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%v\n", err)
		}
	}

	for _, dep := range repo.Dependencies {
		err = os.Mkdir(path.Join(dirname, dep.Name), 0755)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%v\n", err)
		}
		err = dep.WriteRepositoryToFileSystem(path.Join(dirname, dep.Name))
		if err != nil {
			fmt.Fprintf(os.Stderr, "%v\n", err)
		}
	}

	return err
}
