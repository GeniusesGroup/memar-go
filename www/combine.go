/* For license and copyright information please see LEGAL file in repository */

package www

import (
	"bytes"
	"strings"
	"text/template"

	"../convert"
	er "../error"
	"../file"
	"../giti"
	language "../language"
	"../log"
)

type combine struct {
	repo         *file.Folder
	repoGUI      *file.Folder
	repoDL       *file.Folder
	repoPages    *file.Folder
	repoWidgets  *file.Folder
	repoLandings *file.Folder

	inlined map[string]*file.File // map key is file FullName

	swFiles         map[string]*file.File            // map key is language in iso format e.g. "en", "fa", ...
	mainHTMLFiles   map[string]*file.File            // map key is language in iso format e.g. "en", "fa", ...
	mainJSFiles     map[string]*file.File            // map key is language in iso format e.g. "en", "fa", ...
	landingsFiles   map[string]map[string]*file.File // first map key is file name, second map key is language in iso format e.g. "en", "fa",
	designLanguages []*file.File
	otherFiles      []*file.File
}

func (c *combine) init(repo *file.Folder) {
	c.inlined = make(map[string]*file.File)

	c.repo = repo
	c.repoGUI = repo.GetDependency("gui")
	c.repoDL = c.repoGUI.GetDependency("design-languages")
	c.repoPages = c.repoGUI.GetDependency("pages")
	c.repoWidgets = c.repoGUI.GetDependency("widgets")
	c.repoLandings = c.repoGUI.GetDependency("landings")

	c.landingsFiles = make(map[string]map[string]*file.File)

	c.readyLandingsFiles()
	c.readyAppJSFiles()
	c.readyMainHTMLFile()
	c.readySWFile()
	c.readyDesignLanguagesFiles()
}

func (c *combine) readyAppJSFiles() {
	var mainJSFile = c.repoGUI.GetFile("main.js")
	if mainJSFile == nil {
		log.Warn("WWW - no main.js file exist to ready UI apps to serve!")
		return
	}

	var lj = make(localize, 8)
	var jsonFile *file.File = c.repoGUI.GetFile("main.json")
	if jsonFile == nil {
		log.Warn("WWW - Can't find json localize file for /main.js")
		return
	}
	lj.jsonDecoder(jsonFile.Data())

	var mainJSCreator = JS{
		imports: make([]*file.File, 1, 100),
		inlined: c.inlined,
	}
	mainJSCreator.imports[0] = mainJSFile
	_ = mainJSCreator.extractJSImportsRecursively(mainJSFile)
	if giti.AppDevMode {
		var name strings.Builder
		for _, imp := range mainJSCreator.imports {
			name.WriteString(imp.FullName)
			name.WriteByte(',')
		}
		log.Warn("WWW - main.js imports recursively:", name.String())
	}

	c.mainJSFiles = localizeJSFile(mainJSFile, lj)
	for _, mainFile := range c.mainJSFiles {
		mainFile.SetData(nil)
	}

	for i := len(mainJSCreator.imports) - 1; i >= 0; i-- {
		c.addLocalized(mainJSCreator.imports[i])
	}

	// Add platfrom errors!
	for lang, mainFile := range c.mainJSFiles {
		mainFile.AppendData(er.Errors.GetErrorsInJsFormat(language.GetLanguageByISO(lang)))
		mainFile.Rename(mainFile.Name + "-" + lang)
		mainFile.AddHashToName()
	}
}

func (c *combine) addLocalized(file *file.File) {
	if c.isInlined(file.FullName) {
		return
	}
	c.inlined[file.FullName] = file

	if file.Extension != "js" {
		c.otherFiles = append(c.otherFiles, file)
		return
	}

	var files = localeAndMixJSFile(file)
	for lang, mainFile := range c.mainJSFiles {
		var localeFile = files[lang]
		if localeFile == nil {
			if giti.AppDevMode {
				log.Warn("WWW - ", file.FullName, "don't have locale files in '", lang, "' language that main.js support this language. WWW Use pure file in combine proccess.")
			}
			localeFile = files[""]
		}
		mainFile.AppendData(localeFile.Data())
	}
}

func (c *combine) isInlined(fullName string) (ok bool) {
	_, ok = c.inlined[fullName]
	return
}

// readyDesignLanguagesFiles read needed files from given repo folder and do some logic and return files.
func (c *combine) readyDesignLanguagesFiles() {
	for _, folder := range c.repoDL.Folders {
		var combinedFile = file.File{
			Name:      "design-language--" + folder.Name,
			Extension: "css",
		}
		combinedFile.CheckAndFix()
		var fullName = combinedFile.FullName
		var combinedFileData []byte
		for _, file := range folder.Files {
			if file.Extension == "css" {
				combinedFileData = append(combinedFileData, file.Data()...)
			}
		}
		combinedFile.SetData(combinedFileData)
		combinedFile.AddHashToName()
		c.designLanguages = append(c.designLanguages, &combinedFile)

		for _, mainFile := range c.mainJSFiles {
			mainFile.ReplaceAll(convert.UnsafeStringToByteSlice(fullName), convert.UnsafeStringToByteSlice(combinedFile.FullName))
		}
	}
	for _, file := range c.repoDL.Files {
		if file.Extension == "css" {
			var fullName = file.FullName
			file.AddHashToName()
			c.designLanguages = append(c.designLanguages, file)
			for _, mainFile := range c.mainJSFiles {
				mainFile.ReplaceAll(convert.UnsafeStringToByteSlice(fullName), convert.UnsafeStringToByteSlice(file.FullName))
			}
		}
	}
}

// readyLandingsFiles read needed files from given repo folder and do some logic and return files.
func (c *combine) readyLandingsFiles() {
	// TODO::: Need to change landings file name to hash of data??
	var landing *file.File
	for _, landing = range c.repoLandings.GetFiles() {
		switch landing.Extension {
		case "html":
			var jsonFile *file.File = c.repoLandings.GetFile(landing.Name + ".json")
			if jsonFile == nil {
				if giti.AppDevMode {
					log.Warn("WWW - Can't find json localize file for ", landing.FullName)
				}
				continue
			}

			var lj = make(localize, 8)
			lj.jsonDecoder(jsonFile.Data())

			var mixedHTMLFile = mixCSSToHTML(landing, c.repoLandings.GetFile(landing.Name+".css"))
			c.landingsFiles[mixedHTMLFile.Name] = localizeHTMLFile(mixedHTMLFile, lj)
		}
	}
}

// readyMainHTMLFile read needed files from given repo folder and do some logic to make main.html.
func (c *combine) readyMainHTMLFile() {
	var splashFiles = c.landingsFiles["splash"]
	if splashFiles == nil {
		log.Warn("WWW - no splash screen exist to make main.html files")
		return
	}

	for lang, splashFile := range splashFiles {
		var mainJSFile = c.mainJSFiles[lang]
		if mainJSFile == nil {
			log.Warn("WWW - no main js in " + lang + " language exist to make main.html file")
			continue
		}

		var tempName = struct {
			MainJSName string
			Splash     string
		}{
			MainJSName: mainJSFile.Name,
			Splash:     convert.UnsafeByteSliceToString(splashFile.Data()),
		}

		var err error
		var sf = new(bytes.Buffer)
		err = mainHTMLFileTemplate.Execute(sf, tempName)
		if err != nil {
			log.Warn(err)
			continue
		}

		var mainHTMLFile = &file.File{
			Name:      "main+" + lang,
			Extension: "html",
		}
		mainHTMLFile.CheckAndFix()
		mainHTMLFile.SetData(sf.Bytes())
		mainHTMLFile.AddHashToName()
		c.mainHTMLFiles[lang] = mainHTMLFile
	}
}

var mainHTMLFileTemplate = template.Must(template.New("mainHTML").Parse(`
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

// readySWFile read needed files from given repo folder and do some logic to make main.html.
func (c *combine) readySWFile() {
	// set /os/sw.js
	var swFile = c.repoGUI.GetDependency("libjs").GetDependency("os").GetFile("sw.js")
	if swFile == nil {
		log.Warn("WWW - no Service-Worker(sw.js) exist to make locales sw-{{lang}}.js files")
		return
	}

	for lang, mainHTMLFile := range c.mainHTMLFiles {
		var localeSWFile = swFile.DeepCopy()
		localeSWFile.Rename(localeSWFile.Name + "-" + lang)
		localeSWFile.ReplaceAll([]byte("main.html"), []byte(mainHTMLFile.FullName))
		c.swFiles[lang] = localeSWFile
	}
}
