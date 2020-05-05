/* For license and copyright information please see LEGAL file in repository */

package www

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"path"
	"runtime"
	"strings"

	"../assets"
)

// Server store needed data to start a www server!
// This server is to serve GUI architecture
type Server struct {
	Domain string
	Assets *assets.Folder
}

// ServeHTTP use to serve s!
func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var path = strings.Split(r.URL.Path, "/")
	var file = s.Assets.GetFile(path[len(path)-1])
	if file == nil {
		file = s.Assets.GetFile("main.html")
	}

	// serve-to-robots
	// serve-index
	// serve-assets

	w.Header().Set("Content-Type", file.MimeType)
	w.Header().Set("Server", "ChaparKhane")

	w.WriteHeader(200)

	w.Write(file.Data)
}

// CheckDev use to serve the server by local files!
func (s *Server) CheckDev() {
	var dev *bool
	dev = flag.Bool("dev", false, "Add this flag if you want to start server in development phase!")
	flag.Parse()
	if *dev == true {
		// Indicate repoLocation
		_, filename, _, ok := runtime.Caller(0)
		if !ok {
			panic("No caller information, So we can't specify service root location")
		}
		var repoLocation = path.Dir(path.Dir(path.Dir(filename)))

		s.Assets = assets.NewFolder(s.Domain)
		var repo = assets.NewFolder(s.Domain)
	reload:
		readRepositoryFromFileSystem(repoLocation, repo)
		addRepo(s.Assets, repo)
		addGUIToMain(s.Assets, repo)

		fmt.Fprintf(os.Stderr, "%v\n", "Press '''Enter''' key to reload GUI changes")
		var s string
		fmt.Scanln(&s)
		goto reload
	}
}
