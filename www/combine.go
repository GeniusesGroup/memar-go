/* For license and copyright information please see LEGAL file in repository */

package www

import (
	"bytes"
	"strings"
	"text/template"

	"../assets"
	"../convert"
	"../log"
)

type combine struct {
	repo         *assets.Folder
	repoSDK      *assets.Folder
	repoGUI      *assets.Folder
	repoPages    *assets.Folder
	repoWidgets  *assets.Folder
	repoLandings *assets.Folder

	inlined map[string]*assets.File // map key is file FullName

	mainHTML *assets.File
	mainJS   *assets.File
	initsJS  []*assets.File
	landings []*assets.File
}

func (c *combine) init(repo *assets.Folder) {
	c.inlined = make(map[string]*assets.File)

	c.repo = repo
	c.repoSDK = repo.GetDependency("sdk-js")
	c.repoGUI = repo.GetDependency("gui")
	c.repoPages = c.repoGUI.GetDependency("pages")
	c.repoWidgets = c.repoGUI.GetDependency("widgets")
	c.repoLandings = c.repoGUI.GetDependency("landings")
}

func (c *combine) readyAppJSFiles() {
	c.mainJS = c.repoGUI.GetFile("main.js")
	var mainImports, _ = extractJSImportsRecursively(c.mainJS)
	if log.DevMode {
		var name strings.Builder
		for _, imp := range mainImports {
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
	var initImports, _ = extractJSImportsRecursively(initRepoFile)
	if log.DevMode {
		var name strings.Builder
		for _, imp := range initImports {
			name.WriteString(imp.FullName)
			name.WriteByte(',')
		}
		log.Warn("WWW - init.js imports recursively:", name.String())
	}
	var initsFiles = localizeJSFile(initRepoFile, lj)

	for i := 0; i < len(mainImports); i++ {
		if c.isInlined(mainImports[i].FullName) {
			continue
		}
		c.inlined[mainImports[i].FullName] = mainImports[i]
		var files = localeAndMixJSFile(mainImports[i])
		if len(files) < 2 {
			c.mainJS.Data = append(files[""].Data, c.mainJS.Data...)
		} else {
			for lang, initFile := range initsFiles {
				var localeFile = files[lang]
				if localeFile == nil {
					if log.DevMode {
						log.Warn("WWW - ", mainImports[i].FullName, "don't have locale files in '", lang, "' language that init.js support this language. Use pure file in combine proccess.")
					}
					localeFile = files[""]
				}
				initFile.Data = append(localeFile.Data, initFile.Data...)
			}
		}
	}
	for i := 0; i < len(initImports); i++ {
		if c.isInlined(initImports[i].FullName) {
			continue
		}
		c.inlined[initImports[i].FullName] = initImports[i]
		var files = localeAndMixJSFile(initImports[i])
		if len(files) < 2 {
			c.mainJS.Data = append(files[""].Data, c.mainJS.Data...)
		} else {
			for lang, initFile := range initsFiles {
				var localeFile = files[lang]
				if localeFile == nil {
					if log.DevMode {
						log.Warn("WWW - ", initImports[i].FullName, "don't have locale files in '", lang, "' language that init.js support this language. Use pure file in combine proccess.")
					}
					localeFile = files[""]
				}
				initFile.Data = append(localeFile.Data, initFile.Data...)
			}
		}
	}

	var hashedName = "init-" + initsFiles["en"].GetHashOfData() + "-"
	for lang, initFile := range initsFiles {
		initFile.Rename(hashedName + lang)
		c.initsJS = append(c.initsJS, initFile)
	}
}

func (c *combine) isInlined(fullName string) (ok bool) {
	_, ok = c.inlined[fullName]
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
