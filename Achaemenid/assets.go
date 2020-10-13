/* For license and copyright information please see LEGAL file in repository */

package achaemenid

import (
	"fmt"

	as "../assets"
	"../log"
	"../www"
)

// Any data files to serve or use by server!
type assets struct {
	server  *Server
	GUI     *as.Folder
	Objects *as.Folder
	Secret  *as.Folder
	WWW     *as.Folder
	WWWMain *as.File
}

func (a *assets) init(s *Server) {
	a.server = s

	a.GUI = as.NewFolder("gui")
	a.Objects = as.NewFolder("objects")
	a.Secret = as.NewFolder("secret")
	a.WWW = as.NewFolder("www")

	a.LoadFromStorage()
}

func (a *assets) shutdown() {
	// write secret files to storage device if any change made
	a.Secret.WriteRepositoryToFileSystem(a.server.RepoLocation + "/secret")
}

// It block function and must call by seprate goroutine, otherwise it can block other app logic!
func (a *assets) LoadFromStorage() {
	a.Secret.ReadRepositoryFromFileSystem(a.server.RepoLocation+"/secret", true)
	a.WWWMain = www.LoadAssetsFromStorage(a.WWW, a.GUI, a.server.RepoLocation)
}

// It block function and must call by seprate goroutine, otherwise it can block other app logic!
func (a *assets) ReLoadFromStorage() {
	// defer Server.PanicHandler()
reload:
	log.Info("Write '''R''' & press '''Enter''' key to reload GUI changes")
	var non string
	fmt.Scanln(&non)
	if non == "R" || non == "r" {
		a.WWWMain = www.LoadAssetsFromStorage(a.WWW, a.GUI, a.server.RepoLocation)
	} else {
		log.Warn("Requested command not found")
	}
	goto reload
}
