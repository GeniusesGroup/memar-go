/* For license and copyright information please see LEGAL file in repository */

package www

import (
	"strings"

	"../convert"
	"../file"
	language "../language"
	"../mediatype"
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

	contentEncodings []string
}

// Update use to add needed repo files that get from disk or network to the assets!!
func (c *combine) update() {
	c.inlined = make(map[string]protocol.File)
	c.mainHTMLFiles = make(map[string]protocol.File)
	c.mainJSFiles = make(map[string]protocol.File)

	c.guiDir, _ = protocol.App.Files().Directory(guiDirectoryName)
	c.dlDir, _ = c.guiDir.Directory("design-languages")
	c.pagesDir, _ = c.guiDir.Directory("pages")
	c.widgetsDir, _ = c.guiDir.Directory("widgets")
	c.landingsDir, _ = c.guiDir.Directory("landings")
	c.mainHTMLDir, _ = c.guiDir.Directory(mainHTMLDirectoryName)

	c.readyLandingsFiles()
	switch "js" {
	case "go":
		c.readyAppGoToJSFile()
	case "js":
		c.readyAppJSFiles()
	}
	c.readySWFile()
	c.readyDesignLanguagesFiles()
}

func (c *combine) readyAppGoToJSFile() {
	var mainGOFile, _ = c.guiDir.File("main.go")
	if mainGOFile == nil {
		protocol.App.Log(protocol.LogType_Warning, "WWW - no main.go file exist to ready UI apps to serve!")
		return
	}

	// TODO:: compile to JS by https://github.com/gopherjs/gopherjs
}

func (c *combine) readyAppJSFiles() {
	var mainHTMLFile, _ = c.landingsDir.File("main.html")
	if mainHTMLFile == nil {
		protocol.App.Log(protocol.LogType_Warning, "WWW - no main.html file exist to ready UI apps to serve!")
		return
	}

	var mainJSFile, _ = c.guiDir.File("main.js")
	if mainJSFile == nil {
		protocol.App.Log(protocol.LogType_Warning, "WWW - no main.js file exist to ready UI apps to serve!")
		return
	}

	var lj = make(localize, 8)
	var jsonFile, _ = c.guiDir.File("main.json")
	if jsonFile == nil {
		protocol.App.Log(protocol.LogType_Warning, "WWW - Can't find json localize file for /main.js")
		return
	}
	lj.jsonDecoder(jsonFile.Data().Marshal())
	if len(lj) == 0 {
		protocol.App.Log(protocol.LogType_Warning, "WWW - Founded json not meet standards to localize for main.js")
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
			name.WriteString(imp.Metadata().URI().Name())
			name.WriteByte(',')
		}
		protocol.App.Log(protocol.LogType_Warning, "WWW - main.js imports recursively:", name.String())
	}

	for lang := range lj {
		var localeJSFileName = "main-" + lang + ".js"
		var localeJSFile, _ = c.guiDir.File(localeJSFileName)
		// write some data if file exist before
		localeJSFile.Data().Unmarshal([]byte("\n"))
		c.mainJSFiles[lang] = localeJSFile
	}

	for i := len(mainJSCreator.imports) - 1; i >= 0; i-- {
		c.addLocalized(mainJSCreator.imports[i])
	}

	for lang, localeMainJSFile := range c.mainJSFiles {
		// Complete main JS file making proccess
		var sdk, _ = protocol.App.ErrorsSDK(language.GetLanguageByISO(lang), mediatype.JavaScript)
		// Add platfrom errors!
		localeMainJSFile.Data().Append(sdk)
		file.AddHashToFileName(localeMainJSFile)
		file.Compress(localeMainJSFile, c.contentEncodings, protocol.BestCompression)

		// make main.html
		var localeJSFileName = localeMainJSFile.Metadata().URI().Name()
		var localeMainHTMLFileName = "main-" + lang + ".html"
		var localeMainHTMLFile, _ = c.guiDir.File(localeMainHTMLFileName)
		var err = file.Minify(localeMainHTMLFile)
		if err != nil && protocol.AppDebugMode {
			protocol.App.Log(protocol.LogType_Warning, "Minify -", localeMainHTMLFileName, "occur this error:", err)
		}
		localeMainHTMLFile.Data().Replace([]byte("/main.js"), []byte(localeJSFileName), 1)
		file.AddHashToFileName(localeMainHTMLFile)
		file.Compress(localeMainHTMLFile, c.contentEncodings, protocol.BestCompression)
		c.mainHTMLFiles[lang] = localeMainHTMLFile

		// Add localized main html files without any compression to specific serve directory
		var localeMainHTMLFileLangName, _ = c.mainHTMLDir.File(lang)
		localeMainHTMLFileLangName.Data().Unmarshal(localeMainHTMLFile.Data().Marshal())
	}
}

func (c *combine) addLocalized(file protocol.File) {
	if c.isInlined(file.Metadata().URI().Path()) {
		return
	}
	c.inlined[file.Metadata().URI().Path()] = file

	var mixedData = localeAndMixJSFile(file)
	for lang, mainFile := range c.mainJSFiles {
		var localeData = mixedData[lang]
		if localeData != nil {
			mainFile.Data().Append(localeData)
		} else {
			if protocol.AppDevMode {
				protocol.App.Log(protocol.LogType_Warning, "WWW - ", file.Metadata().URI().Name(), "don't have locale files in '", lang, "' language that main.js support this language. WWW Use pure file in combine proccess.")
			}
		}
	}

	// reset file to storage state to undo extract import from file
	file.Data().Unmarshal(nil)
}

func (c *combine) isInlined(uriPath string) (ok bool) {
	_, ok = c.inlined[uriPath]
	return
}

// readyDesignLanguagesFiles read needed files from given repo folder and do some logic and return files.
func (c *combine) readyDesignLanguagesFiles() {
	var err protocol.Error
	for _, dlFile := range c.dlDir.Files(0, 0) {
		if dlFile.Metadata().URI().Extension() == "css" {
			var cssName = dlFile.Metadata().URI().Name()
			var cssFile, _ = c.dlDir.File("h-" + cssName)
			cssFile.Data().Unmarshal(dlFile.Data().Marshal())
			err = file.Minify(cssFile)
			if err != nil && protocol.AppDebugMode {
				protocol.App.Log(protocol.LogType_Warning, "Minify -", cssName, "occur this error:", err)
			}
			file.AddHashToFileName(cssFile)
			file.Compress(cssFile, c.contentEncodings, protocol.BestCompression)

			for _, mainFile := range c.mainJSFiles {
				var fileNameWithoutHashAdded = convert.UnsafeStringToByteSlice(cssName)
				var fileNameWithHashAdded = convert.UnsafeStringToByteSlice(cssFile.Metadata().URI().Name())
				mainFile.Data().Replace(fileNameWithoutHashAdded, fileNameWithHashAdded, -1)
			}
		}
	}
	for _, dir := range c.dlDir.Directories(0, 0) {
		var combinedFileName = dir.Metadata().URI().Name() + ".css"
		var combinedFile, _ = c.dlDir.File(combinedFileName)
		var combinedFileData []byte
		for _, file := range dir.Files(0, 0) {
			if file.Metadata().URI().Extension() == "css" {
				combinedFileData = append(combinedFileData, file.Data().Marshal()...)
			}
		}
		combinedFile.Data().Unmarshal(combinedFileData)
		err = file.Minify(combinedFile)
		if err != nil && protocol.AppDebugMode {
			protocol.App.Log(protocol.LogType_Warning, "Minify -", combinedFileName, "occur this error:", err)
		}
		file.AddHashToFileName(combinedFile)
		file.Compress(combinedFile, c.contentEncodings, protocol.BestCompression)

		for _, mainFile := range c.mainJSFiles {
			var fileNameWithoutHashAdded = convert.UnsafeStringToByteSlice(combinedFileName)
			var fileNameWithHashAdded = convert.UnsafeStringToByteSlice(combinedFile.Metadata().URI().Name())
			mainFile.Data().Replace(fileNameWithoutHashAdded, fileNameWithHashAdded, -1)
		}
	}
}

// readyLandingsFiles read needed files from given repo folder and do some logic and return files.
func (c *combine) readyLandingsFiles() {
	var err protocol.Error
	// TODO::: Need to change landings file name to hash of data??
	for _, landing := range c.landingsDir.Files(0, 0) {
		switch landing.Metadata().URI().Extension() {
		case "html":
			var landingName = landing.Metadata().URI().NameWithoutExtension()
			var landingJsonFile, _ = c.landingsDir.File(landingName + ".json")
			if landingJsonFile == nil {
				protocol.App.Log(protocol.LogType_Warning, "WWW - Can't find json localize json file for ", landing.Metadata().URI().Name())
				continue
			}
			var lj = make(localize)
			lj.jsonDecoder(landingJsonFile.Data().Marshal())
			if len(lj) == 0 {
				protocol.App.Log(protocol.LogType_Warning, "WWW - Founded json not meet standards to localize for ", landing.Metadata().URI().Name())
				continue
			}

			var landingCssFile, _ = c.landingsDir.File(landingName + ".css")
			var mixedHTMLFile = mixCSSToHTML(landing.Data().Marshal(), landingCssFile.Data().Marshal())

			var localized map[string][]byte
			localized, err = localizeHTMLFile(mixedHTMLFile, lj)
			if err != nil {
				protocol.App.Log(protocol.LogType_Warning, "WWW - Can't localize ", landing.Metadata().URI().Name(), " landing page due to ", err)
			}
			for lang, data := range localized {
				var localizedFileName = landingName + "-" + lang + ".html"
				var localizedLanding, _ = c.landingsDir.File(localizedFileName)
				localizedLanding.Data().Unmarshal(data)
				err = file.Minify(localizedLanding)
				if err != nil && protocol.AppDebugMode {
					protocol.App.Log(protocol.LogType_Warning, "WWW Minify -", localizedFileName, "occur this error:", err)
				}
				file.Compress(localizedLanding, c.contentEncodings, protocol.BestCompression)
			}
		}
	}
}

// readySWFile read needed files from given repo folder and do some logic to make main.html.
func (c *combine) readySWFile() {
	var libJSDir, _ = c.guiDir.Directory("libjs")
	var osDir, _ = libJSDir.Directory("os")
	var swFile, _ = osDir.File("sw.js")
	if swFile == nil {
		protocol.App.Log(protocol.LogType_Warning, "WWW - Service-Worker(sw.js) not exist in /gui/libjs/os/sw.js to make locales sw-{{lang}}.js files")
		return
	}

	for lang, mainHTMLFile := range c.mainHTMLFiles {
		var localeSWFileName = "sw-" + lang + ".js"
		var localeSWFile, _ = c.guiDir.File(localeSWFileName)
		localeSWFile.Data().Unmarshal(swFile.Data().Marshal())
		localeSWFile.Data().Replace([]byte("main.html"), []byte(mainHTMLFile.Metadata().URI().Name()), -1)
		var err = file.Minify(localeSWFile)
		if err != nil && protocol.AppDebugMode {
			protocol.App.Log(protocol.LogType_Warning, "Minify -", localeSWFileName, "occur this error:", err)
		}
		file.Compress(localeSWFile, c.contentEncodings, protocol.BestCompression)
	}
}
