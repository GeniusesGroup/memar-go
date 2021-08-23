/* For license and copyright information please see LEGAL file in repository */

package www

import (
	"bytes"
	"strconv"
	"strings"

	"../file"
	"../giti"
	"../log"
)

// JS store data for its method
type JS struct {
	imports []*file.File
	inlined map[string]*file.File
}

func (js *JS) checkInlined(srcJS *file.File) (inlined bool) {
	_, inlined = js.inlined[srcJS.FullName]
	return
}

// AddJSToJS use to add a JS file to other!
func (js *JS) addJSToJS(srcJS, desJS *file.File) {
	if js.checkInlined(srcJS) {
		return
	}
	js.inlined[srcJS.FullName] = srcJS
	desJS.AppendData(srcJS.Data())
}

// AddJSToJSRecursively use to add JSS and all import to JS file!
func (js *JS) addJSToJSRecursively(srcJS, desJS *file.File) {
	if js.checkInlined(srcJS) {
		return
	}
	// Tell other this file will add to desJS later!
	js.inlined[srcJS.FullName] = srcJS

	var imports = make([]*file.File, 0, 16)
	_ = js.extractJSImports(srcJS)
	for _, imp := range imports {
		js.addJSToJSRecursively(imp, desJS)
	}
	desJS.AppendData(srcJS.Data())
}

func (js *JS) extractJSImportsRecursively(jsFile *file.File) (err giti.Error) {
	var lastImportIndex = len(js.imports)
	_ = js.extractJSImports(jsFile)
	for _, imp := range js.imports[lastImportIndex:] {
		err = js.extractJSImportsRecursively(imp)
	}
	return
}

func (js *JS) extractJSImports(jsFile *file.File) (err giti.Error) {
	var jsData = jsFile.Data()
	var importKeywordIndex, st, en int
	var loc, fileName string
	var locPart []string
	var imDep *file.Folder
	var imFile *file.File
	for {
		importKeywordIndex = bytes.Index(jsData, []byte("import "))
		if importKeywordIndex == -1 {
			break
		}
		// Find start and end of import file location!
		st = importKeywordIndex + bytes.IndexByte(jsData[importKeywordIndex:], '\'') + 1
		en = st + bytes.IndexByte(jsData[st:], '\'')
		loc = string(jsData[st:en])

		locPart = strings.Split(loc, "/")
		if len(locPart) < 2 {
			// don't parse dynamically import in files
			break
		} else {
			imDep = jsFile.Folder
			for i := 0; i < len(locPart)-1; i++ { // -1 due have file name at end of locPart
				switch locPart[i] {
				case ".":
					// noting to do!
				case "..":
					imDep = imDep.Folder
				default:
					imDep = imDep.GetDependencyRecursively(locPart[i])
					if imDep == nil {
						// err =
						return
					}
				}
			}
		}

		fileName = locPart[len(locPart)-1]
		imFile = imDep.GetFile(fileName)
		if imFile == nil {
			if giti.AppDevMode {
				log.Warn("WWW - ", fileName, "indicate as import in", jsFile.FullName, "can't find in repo")
			}
			// err =
			return
		}
		js.imports = append(js.imports, imFile)

		copy(jsData[importKeywordIndex-1:], jsData[en+1:])
		jsData = jsData[:len(jsData)-(en-importKeywordIndex)-2]
	}
	jsFile.SetData(jsData)
	return
}

func localeAndMixJSFile(jsFile *file.File) (files map[string]*file.File) {
	var cssFile = jsFile.Folder.GetFile(jsFile.Name + ".css")
	if giti.AppDevMode && cssFile == nil {
		log.Warn("WWW - ", jsFile.FullName, "Can't find CSS style file, Mix CSS to JS file skipped")
	}
	if cssFile != nil {
		jsFile = mixCSSToJS(jsFile, cssFile)
	}

	var lj = make(localize, 8)
	var jsonFile *file.File = jsFile.Folder.GetFile(jsFile.Name + ".json")
	if giti.AppDevMode && jsonFile == nil {
		log.Warn("WWW - ", jsFile.FullName, "Can't find JSON localize file, Localization skipped")
	}
	if jsonFile != nil {
		lj.jsonDecoder(jsonFile.Data())
	}

	files = localizeJSFile(jsFile, lj)

	var htmlFile = jsFile.Folder.GetFile(jsFile.Name + ".html")
	if giti.AppDevMode && htmlFile == nil {
		log.Warn("WWW - ", jsFile.FullName, "Can't find HTML style file, Mix HTML to JS file skipped")
	}
	if htmlFile != nil {
		var htmlFiles = localizeHTMLFile(htmlFile, lj)
		for lang, f := range htmlFiles {
			mixHTMLToJS(files[lang], f)
		}
	}

	for _, f := range jsFile.Folder.FindFiles(jsFile.Name + "-template-") {
		var namesPart = strings.Split(f.Name, "-template-")
		if namesPart[0] == jsFile.Name {
			var htmlFiles = localizeHTMLFile(f, lj)
			for lang, f := range htmlFiles {
				mixHTMLTemplateToJS(files[lang], f, namesPart[1])
			}
		}
	}

	return
}

// localizeJSFile make and returns number of localize file by number of language indicate in JSON localize text
func localizeJSFile(fi *file.File, lj localize) (files map[string]*file.File) {
	files = make(map[string]*file.File, len(lj))
	if len(lj) == 0 {
		files[""] = fi
	} else {
		for lang, text := range lj {
			files[lang] = replaceLocalizeTextInJS(fi, text, lang)
		}
	}
	return
}

func replaceLocalizeTextInJS(js *file.File, text []string, lang string) (newFile *file.File) {
	newFile = js.Copy()
	var jsData = js.Data()
	var newFileData = make([]byte, 0, len(jsData))

	var replacerIndex int
	var bracketIndex int
	var textIndex uint64
	var err error
	for {
		replacerIndex = bytes.Index(jsData, []byte("LocaleText["))
		if replacerIndex < 0 {
			newFileData = append(newFileData, jsData...)
			newFile.SetData(newFileData)
			return
		}
		newFileData = append(newFileData, jsData[:replacerIndex]...)
		jsData = jsData[replacerIndex:]

		bracketIndex = bytes.IndexByte(jsData, ']')
		textIndex, err = strconv.ParseUint(string(jsData[11:bracketIndex]), 10, 8)
		if err != nil {
			log.Warn("WWW - ", js.FullName, "Index numbers is not valid, double check your file for bad structures")
		} else {
			newFileData = append(newFileData, text[textIndex]...)
		}

		jsData = jsData[bracketIndex+1:]
	}
}
