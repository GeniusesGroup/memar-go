/* For license and copyright information please see LEGAL file in repository */

package file

import (
	"os"
	"sort"
)

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
