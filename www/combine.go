/* For license and copyright information please see LEGAL file in repository */

package www

import (
	"bytes"
	"strings"
	"text/template"

	"../assets"
	"../convert"
	er "../error"
	"../log"
)

type combine struct {
	repo         *assets.Folder
	repoGUI      *assets.Folder
	repoPages    *assets.Folder
	repoWidgets  *assets.Folder
	repoLandings *assets.Folder

	inlined    map[string]*assets.File // map key is file FullName
	initsFiles map[string]*assets.File

	mainHTML   *assets.File
	mainJS     *assets.File
	initsJS    []*assets.File
	landings   []*assets.File
	otherFiles []*assets.File
}

func (c *combine) init(repo *assets.Folder) {
	c.inlined = make(map[string]*assets.File)

	c.repo = repo
	c.repoGUI = repo.GetDependency("gui")
	c.repoPages = c.repoGUI.GetDependency("pages")
	c.repoWidgets = c.repoGUI.GetDependency("widgets")
	c.repoLandings = c.repoGUI.GetDependency("landings")
}

func (c *combine) readyAppJSFiles() {
	var mainJSFile = c.repoGUI.GetFile("main.js")
	c.mainJS = mainJSFile.Copy()
	c.mainJS.Data = nil

	var mainJSCreator = JS{
		imports: make([]*assets.File, 0, 100),
		inlined: c.inlined,
	}
	mainJSCreator.imports = append(mainJSCreator.imports, mainJSFile)
	_ = mainJSCreator.extractJSImportsRecursively(mainJSFile)
	if log.DevMode {
		var name strings.Builder
		for _, imp := range mainJSCreator.imports {
			name.WriteString(imp.FullName)
			name.WriteByte(',')
		}
		log.Warn("WWW - main.js imports recursively:", name.String())
	}

	var lj = make(localize, 8)
	var jsonFile *assets.File = c.repoGUI.GetFile("init.json")
	if jsonFile == nil {
		log.Warn("WWW - Can't find json localize file for /init.js")
		return
	}
	lj.jsonDecoder(jsonFile.Data)

	var initRepoFile = c.repoGUI.GetFile("init.js")
	var initJSCreator = JS{
		imports: make([]*assets.File, 0, 100),
		inlined: c.inlined,
	}
	initJSCreator.imports = append(initJSCreator.imports, initRepoFile)
	_ = initJSCreator.extractJSImportsRecursively(initRepoFile)
	if log.DevMode {
		var name strings.Builder
		for _, imp := range initJSCreator.imports {
			name.WriteString(imp.FullName)
			name.WriteByte(',')
		}
		log.Warn("WWW - init.js imports recursively:", name.String())
	}
	c.initsFiles = localizeJSFile(initRepoFile, lj)
	for _, initFile := range c.initsFiles {
		initFile.Data = nil
	}

	for i := len(mainJSCreator.imports) - 1; i >= 0; i-- {
		if c.isInlined(mainJSCreator.imports[i].FullName) {
			continue
		}
		if c.isLocalize(mainJSCreator.imports[i]) {
			c.addLocalized(mainJSCreator.imports[i])
			continue
		}
		c.inlined[mainJSCreator.imports[i].FullName] = mainJSCreator.imports[i]
		var files = localeAndMixJSFile(mainJSCreator.imports[i])
		c.mainJS.Data = append(c.mainJS.Data, files[""].Data...)
	}

	for i := len(initJSCreator.imports) - 1; i >= 0; i-- {
		c.addLocalized(initJSCreator.imports[i])
	}

	for _, err := range er.ERRPoolSlice {
		var domain, short, long, idAsString string
		for lang, initFile := range c.initsFiles {
			if lang == "en" {
				domain, short, long, idAsString = err.GetDetail(0)
			} else {
				domain, short, long, idAsString = err.GetDetail(1)
			}
			if short == "" {
				domain, short, long, idAsString = err.GetDetail(0)
			}
			var textOfError = "errors.New(" + idAsString + ",\"" + domain + "\",\"" + short + "\",\"" + long + "\")\n"
			initFile.Data = append(initFile.Data, textOfError...)
		}
	}

	var hashedName = "init-" + c.initsFiles["en"].GetHashOfData() + "-"
	for lang, initFile := range c.initsFiles {
		initFile.Rename(hashedName + lang)
		c.initsJS = append(c.initsJS, initFile)
	}
}

func (c *combine) addLocalized(file *assets.File) {
	if c.isInlined(file.FullName) {
		return
	}
	c.inlined[file.FullName] = file

	if file.Extension != "js" {
		c.otherFiles = append(c.otherFiles, file)
		return
	}

	var files = localeAndMixJSFile(file)
	if len(files) < 2 {
		c.mainJS.Data = append(c.mainJS.Data, files[""].Data...)
	} else {
		for lang, initFile := range c.initsFiles {
			var localeFile = files[lang]
			if localeFile == nil {
				if log.DevMode {
					log.Warn("WWW - ", file.FullName, "don't have locale files in '", lang, "' language that init.js support this language. Use pure file in combine proccess.")
				}
				localeFile = files[""]
			}
			initFile.Data = append(initFile.Data, localeFile.Data...)
		}
	}
}

func (c *combine) isInlined(fullName string) (ok bool) {
	_, ok = c.inlined[fullName]
	return
}

func (c *combine) isLocalize(file *assets.File) (ok bool) {
	var jsonFile *assets.File = file.Dep.GetFile(file.Name + ".json")
	if jsonFile != nil {
		return true
	}
	return
}

// readyLandingsFiles read needed files from given repo folder and do some logic and return files.
func (c *combine) readyLandingsFiles() {
	var landing *assets.File
	for _, landing = range c.repoLandings.GetFiles() {
		switch landing.Extension {
		case "html":
			var jsonFile *assets.File = c.repoLandings.GetFile(landing.Name + ".json")
			if jsonFile == nil {
				if log.DevMode {
					log.Warn("WWW - Can't find json localize file for ", landing.FullName)
				}
				continue
			}

			var lj = make(localize, 8)
			lj.jsonDecoder(jsonFile.Data)

			var mixed = mixCSSToHTML(landing, c.repoLandings.GetFile(landing.Name+".css"))

			for lang, f := range localizeHTMLFile(mixed, lj) {
				f.Rename(f.Name + "-" + lang)
				c.landings = append(c.landings, f)
			}
		}
	}
}

// readyMainHTMLFile read needed files from given repo folder and do some logic to make main.html.
func (c *combine) readyMainHTMLFile(mainJSName string) {
	var splash *assets.File
	for _, lan := range c.landings {
		if lan.Name == "splash-en" {
			splash = lan
		}
	}
	if splash == nil {
		log.Warn("WWW - no english splash screen exist to make main.html file")
		return
	}

	var tempName = struct {
		MainJSName string
		Splash     string
	}{
		MainJSName: mainJSName,
		Splash:     convert.UnsafeByteSliceToString(splash.Data),
	}

	var err error
	var sf = new(bytes.Buffer)
	err = mainHTMLFile.Execute(sf, tempName)
	if err != nil {
		log.Warn(err)
	}

	c.mainHTML = &assets.File{
		FullName:  "main.html",
		Name:      "main",
		Extension: "html",
		MimeType:  "text/html",
		Data:      sf.Bytes(),
	}
	c.mainHTML.AddHashToName()
}

var mainHTMLFile = template.Must(template.New("mainHTML").Parse(`
<!-- For license and copyright information please see LEGAL file in repository -->
<!DOCTYPE html>
<html vocab="http://schema.org/" prefix="">

<head>
    <meta charset="utf-8" />
	<meta name="viewport" content="width=device-width, initial-scale=1" />
    <script src="/{{.MainJSName}}.js" async></script>
</head>

<body>
{{.Splash}}
</body>

</html>
`))
