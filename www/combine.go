/* For license and copyright information please see LEGAL file in repository */

package www

import (
	"bytes"
	"hash/crc32"
	"strconv"
	"strings"
	"text/template"
	"unsafe"

	"../assets"
	"../log"
)

type combine struct {
	repo         *assets.Folder
	repoSDK      *assets.Folder
	repoGUI      *assets.Folder
	repoPages    *assets.Folder
	repoWidgets  *assets.Folder
	repoLandings *assets.Folder

	mainInline  []*assets.File
	initInline  []*assets.File
	inlined     map[string]*assets.File   // map key is file FullName
	localeFiles map[string][]*assets.File // map key is lang

	mainHTML *assets.File
	mainJS   *assets.File
	initsJS  []*assets.File
	landings []*assets.File
}

func (c *combine) init(repo *assets.Folder) {
	c.inlined = make(map[string]*assets.File)
	c.localeFiles = make(map[string][]*assets.File)

	c.repo = repo
	c.repoSDK = repo.GetDependency("sdk-js")
	c.repoGUI = repo.GetDependency("gui")
	c.repoPages = c.repoGUI.GetDependency("pages")
	c.repoWidgets = c.repoGUI.GetDependency("widgets")
	c.repoLandings = c.repoGUI.GetDependency("landings")
}

func (c *combine) do() {
	c.localeAndMixFolders()

	c.readyMainJSFile()
	c.readyInitsJSFile()
	c.readyLandingsFiles()
}

func (c *combine) readyMainJSFile() {
	var main = c.repoGUI.GetFile("main.js")
	c.mainJS = main.Copy()
	c.mainJS.Data = nil // due to main.js data must at the end of combined main.js

	// InlineFilesRecursively use to inline files to the given file Recursively!!
	var i int = len(c.mainInline) - 1
	for ; i >= 0; i-- {
		addJSToJSRecursively(c.repo, c.mainInline[i], c.mainJS, c.inlined)
	}

	// Add SDK-JS to main.js
	for _, sdk := range c.repoSDK.GetFiles() {
		c.inlined[sdk.FullName] = sdk
		c.mainJS.Data = append(c.mainJS.Data, sdk.Data...)
	}

	addJSToJSRecursively(c.repo, main, c.mainJS, c.inlined)
}

func (c *combine) readyInitsJSFile() {
	var lj = make(localize, 8)
	var jsonFile *assets.File = c.repoGUI.GetFile("init.json")
	if jsonFile == nil {
		log.Warn("Can't find json localize file for /init.js")
		return
	}
	lj.jsonDecoder(jsonFile.Data)

	var initsFiles = localizeJSFile(c.repoGUI.GetFile("init.js"), lj)

	for lang, files := range c.localeFiles {
		if initsFiles[lang] == nil {
			log.Warn("Some pages||widgets||... have locale files in '", lang, "' language that init.js doesn't support this language")
			continue
		}
		var initLocale = initsFiles[lang].Copy()
		initLocale.Data = nil
		for _, file := range files {
			addJSToJSRecursively(c.repo, file, initLocale, c.inlined)
			// c.inlined[file.FullName] = file
			// initLocale.Data = append(initsLocale[lang].Data, file.Data...)
		}
		addJSToJSRecursively(c.repo, initsFiles[lang], initLocale, c.inlined)
		c.initsJS = append(c.initsJS, initLocale)
	}
}

func (c *combine) localeAndMixFolders() {
	c.localeAndMixFiles(c.repoPages)
	c.localeAndMixFiles(c.repoWidgets)
	c.localeAndMixFiles(c.repoGUI.GetDependency("libjs").GetDependency("widget-localize"))
}

func (c *combine) localeAndMixFiles(repo *assets.Folder) {
	for _, file := range repo.GetFiles() {
		switch file.Extension {
		case "js":
			var cssFile = repo.GetFile(file.Name + ".css")
			if cssFile == nil {
				log.Warn("Can't find CSS style file for", file.FullName, ", Mix CSS to JS file skipped")
			} else {
				file = mixCSSToJS(file, cssFile)
			}

			var lj = make(localize, 8)
			var jsonFile *assets.File = repo.GetFile(file.Name + ".json")
			if jsonFile == nil {
				log.Warn("Can't find json localize file for", file.FullName, ", Localization skipped and this file wouldn't add to init.js")
			} else {
				lj.jsonDecoder(jsonFile.Data)
			}

			var jsFiles = localizeJSFile(file, lj)
			for lang, js := range jsFiles {
				c.localeFiles[lang] = append(c.localeFiles[lang], js)
			}

			var htmlFile = repo.GetFile(file.Name + ".html")
			if htmlFile == nil {
				log.Warn("Can't find HTML style file for", file.FullName, ", Mix HTML to JS file skipped")
			} else {
				var htmlFiles = localizeHTMLFile(htmlFile, lj)
				for lang, f := range htmlFiles {
					mixHTMLToJS(jsFiles[lang], f)
				}
			}

			for _, f := range repo.FindFiles(file.Name + "-template-") {
				var namesPart = strings.Split(f.Name, "-template-")
				if namesPart[0] == file.Name {
					var htmlFiles = localizeHTMLFile(f, lj)
					for lang, f := range htmlFiles {
						mixHTMLTemplateToJS(jsFiles[lang], f, namesPart[1])
					}
				}
			}
		}
	}
}

// readyLandingsFiles read needed files from given repo folder and do some logic and return files.
func (c *combine) readyLandingsFiles() {
	var landing *assets.File
	for _, landing = range c.repoLandings.GetFiles() {
		switch landing.Extension {
		case "html":
			var jsonFile *assets.File = c.repoLandings.GetFile(landing.Name + ".json")
			if jsonFile == nil {
				log.Warn("Can't find json localize file for ", landing.FullName)
				continue
			}

			var lj = make(localize, 8)
			lj.jsonDecoder(jsonFile.Data)

			var mixed = mixCSSToHTML(landing, c.repoLandings.GetFile(landing.Name+".css"))

			for _, f := range localizeHTMLFile(mixed, lj) {
				c.landings = append(c.landings, f)
			}
		}
	}
}

// readyMainHTMLFile read needed files from given repo folder and do some logic to make main.html.
func (c *combine) readyMainHTMLFile(mainJSName string) {
	var splash = mixCSSToHTML(c.repoLandings.GetFile("splash.html"), c.repoLandings.GetFile("splash.css"))

	var tempName = struct {
		MainJSName string
		Splash     string
	}{
		MainJSName: mainJSName,
		Splash:     *(*string)(unsafe.Pointer(&splash.Data)),
	}

	var err error
	var sf = new(bytes.Buffer)
	err = mainHTMLFile.Execute(sf, tempName)
	if err != nil {
		log.Warn(err)
	}

	c.mainHTML = &assets.File{}
	c.mainHTML.Data = sf.Bytes()
	c.mainHTML.Name = strconv.FormatUint(uint64(crc32.ChecksumIEEE(c.mainHTML.Data)), 10)
	c.mainHTML.FullName = c.mainHTML.Name + ".html"
	c.mainHTML.Extension = "html"
	c.mainHTML.MimeType = "text/html"
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
