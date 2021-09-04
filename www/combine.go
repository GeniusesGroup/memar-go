/* For license and copyright information please see LEGAL file in repository */

package www

import (
	"bytes"
	"strings"
	"text/template"

	"../codec"
	"../convert"
	"../file"
	language "../language"
	"../log"
	"../protocol"
)

type combine struct {
	guiDir      protocol.FileDirectory
	dlDir       protocol.FileDirectory
	pagesDir    protocol.FileDirectory
	widgetsDir  protocol.FileDirectory
	landingsDir protocol.FileDirectory
	mainHTMLDir protocol.FileDirectory // files name is just language in iso format e.g. "en", "fa",

	inlined       map[string]protocol.File // map key is file path
	mainHTMLFiles map[string]protocol.File // map key is language in iso format e.g. "en", "fa", ...
	mainJSFiles   map[string]protocol.File // map key is language in iso format e.g. "en", "fa", ...
}

// Update use to add needed repo files that get from disk or network to the assets!!
func (c *combine) update() {
	c.inlined = make(map[string]protocol.File)
	c.mainHTMLFiles = make(map[string]protocol.File)
	c.mainJSFiles = make(map[string]protocol.File)

	c.guiDir, _ = protocol.App.FileDirectory().Directory(guiDirectoryName)
	c.dlDir, _ = c.guiDir.Directory("design-languages")
	c.pagesDir, _ = c.guiDir.Directory("pages")
	c.widgetsDir, _ = c.guiDir.Directory("widgets")
	c.landingsDir, _ = c.guiDir.Directory("landings")
	c.mainHTMLDir, _ = c.guiDir.Directory(mainHTMLDirectoryName)

	c.readyLandingsFiles()
	c.readyAppGoToJsFile()
	c.readyAppJSFiles()
	c.readyMainHTMLFile()
	c.readySWFile()
	c.readyDesignLanguagesFiles()
}

func (c *combine) readyAppGoToJsFile() {
	var mainGOFile, _ = c.guiDir.File("main.go")
	if mainGOFile == nil {
		log.Warn("WWW - no main.go file exist to ready UI apps to serve!")
		return
	}

	// TODO:: compile to JS by https://github.com/gopherjs/gopherjs
}

func (c *combine) readyAppJSFiles() {
	var mainJSFile, _ = c.guiDir.File("main.js")
	if mainJSFile == nil {
		log.Warn("WWW - no main.js file exist to ready UI apps to serve!")
		return
	}

	var lj = make(localize, 8)
	var jsonFile, _ = c.guiDir.File("main.json")
	if jsonFile == nil {
		log.Warn("WWW - Can't find json localize file for /main.js")
		return
	}
	lj.jsonDecoder(jsonFile.Data().Marshal())
	if len(lj) == 0 {
		log.Warn("WWW - Founded json not meet standards to localize for main.js")
		return
	}

	var mainJSCreator = JS{
		imports: make([]protocol.File, 1, 100),
		inlined: c.inlined,
	}
	mainJSCreator.imports[0] = mainJSFile
	_ = mainJSCreator.extractJSImportsRecursively(mainJSFile)
	if protocol.AppDevMode {
		var name strings.Builder
		for _, imp := range mainJSCreator.imports {
			name.WriteString(imp.MetaData().URI().Name())
			name.WriteByte(',')
		}
		log.Warn("WWW - main.js imports recursively:", name.String())
	}

	for lang, _ := range lj {
		var localeJSFile, _ = c.guiDir.File("main-" + lang + ".js")
		// write some data if file exist before
		localeJSFile.Data().UnMarshal([]byte("\n"))
		c.mainJSFiles[lang] = localeJSFile
	}

	for i := len(mainJSCreator.imports) - 1; i >= 0; i-- {
		c.addLocalized(mainJSCreator.imports[i])
	}

	for lang, mainFile := range c.mainJSFiles {
		// Add platfrom errors!
		mainFile.Data().Append(protocol.App.ErrorsSDK(language.GetLanguageByISO(lang), "js"))

		mainFile.Compress(codec.CompressTypeGZIP)
		file.AddHashToFileName(mainFile)
	}
}

func (c *combine) addLocalized(file protocol.File) {
	if c.isInlined(file.MetaData().URI().Path()) {
		return
	}
	c.inlined[file.MetaData().URI().Path()] = file

	var mixedData = localeAndMixJSFile(file)
	for lang, mainFile := range c.mainJSFiles {
		var localeData = mixedData[lang]
		if localeData != nil {
			mainFile.Data().Append(localeData)
		} else {
			if protocol.AppDevMode {
				log.Warn("WWW - ", file.MetaData().URI().Name(), "don't have locale files in '", lang, "' language that main.js support this language. WWW Use pure file in combine proccess.")
			}
		}
	}

	// reset file to storage state to undo extract import from file
	file.Data().UnMarshal(nil)
}

func (c *combine) isInlined(uriPath string) (ok bool) {
	_, ok = c.inlined[uriPath]
	return
}

// readyDesignLanguagesFiles read needed files from given repo folder and do some logic and return files.
func (c *combine) readyDesignLanguagesFiles() {
	var err protocol.Error
	for _, dlFile := range c.dlDir.Files(0, 0) {
		if dlFile.MetaData().URI().Extension() == "css" {
			var cssName = dlFile.MetaData().URI().Name()
			var cssFile, _ = c.dlDir.File("h-" + cssName)
			cssFile.Data().UnMarshal(dlFile.Data().Marshal())
			err = cssFile.Minify()
			if err != nil && protocol.AppDebugMode {
				log.Warn("Minify -", cssName, "occur this error:", err)
			}
			cssFile.Compress(codec.CompressTypeGZIP)
			file.AddHashToFileName(cssFile)

			for _, mainFile := range c.mainJSFiles {
				var fileNameWithoutHashAdded = convert.UnsafeStringToByteSlice(cssName)
				var fileNameWithHashAdded = convert.UnsafeStringToByteSlice(cssFile.MetaData().URI().Name())
				mainFile.Data().Replace(fileNameWithoutHashAdded, fileNameWithHashAdded, -1)
			}
		}
	}
	for _, dir := range c.dlDir.Directories(0, 0) {
		var combinedFileName = dir.MetaData().URI().Name() + ".css"
		var combinedFile, _ = c.dlDir.File(combinedFileName)
		var combinedFileData []byte
		for _, file := range dir.Files(0, 0) {
			if file.MetaData().URI().Extension() == "css" {
				combinedFileData = append(combinedFileData, file.Data().Marshal()...)
			}
		}
		combinedFile.Data().UnMarshal(combinedFileData)
		err = combinedFile.Minify()
		if err != nil && protocol.AppDebugMode {
			log.Warn("Minify -", combinedFileName, "occur this error:", err)
		}
		combinedFile.Compress(codec.CompressTypeGZIP)
		file.AddHashToFileName(combinedFile)

		for _, mainFile := range c.mainJSFiles {
			var fileNameWithoutHashAdded = convert.UnsafeStringToByteSlice(combinedFileName)
			var fileNameWithHashAdded = convert.UnsafeStringToByteSlice(combinedFile.MetaData().URI().Name())
			mainFile.Data().Replace(fileNameWithoutHashAdded, fileNameWithHashAdded, -1)
		}
	}
}

// readyLandingsFiles read needed files from given repo folder and do some logic and return files.
func (c *combine) readyLandingsFiles() {
	var err protocol.Error
	// TODO::: Need to change landings file name to hash of data??
	for _, landing := range c.landingsDir.Files(0, 0) {
		switch landing.MetaData().URI().Extension() {
		case "html":
			var landingName = landing.MetaData().URI().NameWithoutExtension()
			var landingJsonFile, _ = c.landingsDir.File(landingName + ".json")
			if landingJsonFile == nil {
				log.Warn("WWW - Can't find json localize json file for ", landing.MetaData().URI().Name())
				continue
			}
			var lj = make(localize)
			lj.jsonDecoder(landingJsonFile.Data().Marshal())
			if len(lj) == 0 {
				log.Warn("WWW - Founded json not meet standards to localize for ", landing.MetaData().URI().Name())
				continue
			}

			var landingCssFile, _ = c.landingsDir.File(landingName + ".css")
			var mixedHTMLFile = mixCSSToHTML(landing.Data().Marshal(), landingCssFile.Data().Marshal())

			var localized map[string][]byte
			localized, err = localizeHTMLFile(mixedHTMLFile, lj)
			if err != nil {
				log.Warn("WWW - Can't localize ", landing.MetaData().URI().Name(), " landing page due to ", err)
			}
			for lang, data := range localized {
				var localizedFileName = landingName + "-" + lang + ".html"
				var localizedLanding, _ = c.landingsDir.File(localizedFileName)
				localizedLanding.Data().UnMarshal(data)
				err = localizedLanding.Minify()
				if err != nil && protocol.AppDebugMode {
					log.Warn("WWW Minify -", localizedFileName, "occur this error:", err)
				}
				localizedLanding.Compress(codec.CompressTypeGZIP)
			}
		}
	}
}

// readyMainHTMLFile read needed files from given repo folder and do some logic to make main.html.
func (c *combine) readyMainHTMLFile() {
	for lang, jsFile := range c.mainJSFiles {
		var splashFile, _ = c.landingsDir.File("splash-" + lang + ".html")
		if splashFile == nil {
			log.Warn("WWW - no locale splash screen exist to make main.html files for " + lang + " language of main.js")
			continue
		}

		var tempName = struct {
			MainJSName string
			Splash     string
		}{
			MainJSName: jsFile.MetaData().URI().Name(),
			Splash:     convert.UnsafeByteSliceToString(splashFile.Data().Marshal()),
		}
		var err error
		var sf = new(bytes.Buffer)
		err = mainHTMLFileTemplate.Execute(sf, tempName)
		if err != nil {
			log.Warn("WWW - Making main.html failed with this error: ", err)
			continue
		}

		var localeMainFileName = "main-" + lang + ".html"
		var localeMainFile, _ = c.guiDir.File(localeMainFileName)
		localeMainFile.Data().UnMarshal(sf.Bytes())
		err = localeMainFile.Minify()
		if err != nil && protocol.AppDebugMode {
			log.Warn("Minify -", localeMainFileName, "occur this error:", err)
		}
		localeMainFile.Compress(codec.CompressTypeGZIP)
		file.AddHashToFileName(localeMainFile)
		c.mainHTMLFiles[lang] = localeMainFile

		var localeMainFileLangName, _ = c.mainHTMLDir.File(lang)
		localeMainFileLangName.Data().UnMarshal(localeMainFile.Data().Marshal())
	}
	// for _, fi := range c.mainHTMLFiles {
	// 	a.Folder.MinifyCompressSet(fi, codec.CompressTypeGZIP)
	// }
}

var mainHTMLFileTemplate = template.Must(template.New("mainHTML").Parse(`
<!-- For license and copyright information please see LEGAL file in repository -->
<!DOCTYPE html>
<html vocab="http://schema.org/" prefix="">

<head>
    <meta charset="utf-8" />
	<meta name="viewport" content="width=device-width, initial-scale=1" />
    <script src="/{{.MainJSName}}" async></script>
</head>

<body>
{{.Splash}}
</body>

</html>
`))

// readySWFile read needed files from given repo folder and do some logic to make main.html.
func (c *combine) readySWFile() {
	var libJSDir, _ = c.guiDir.Directory("libjs")
	var osDir, _ = libJSDir.Directory("os")
	var swFile, _ = osDir.File("sw.js")
	if swFile == nil {
		log.Warn("WWW - Service-Worker(sw.js) not exist in /gui/libjs/os/sw.js to make locales sw-{{lang}}.js files")
		return
	}

	for lang, mainHTMLFile := range c.mainHTMLFiles {
		var localeSWFileName = "sw-" + lang + ".js"
		var localeSWFile, _ = c.guiDir.File(localeSWFileName)
		localeSWFile.Data().UnMarshal(swFile.Data().Marshal())
		localeSWFile.Data().Replace([]byte("main.html"), []byte(mainHTMLFile.MetaData().URI().Name()), -1)
		var err = localeSWFile.Minify()
		if err != nil && protocol.AppDebugMode {
			log.Warn("Minify -", localeSWFileName, "occur this error:", err)
		}
		localeSWFile.Compress(codec.CompressTypeGZIP)
	}
}
