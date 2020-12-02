/* For license and copyright information please see LEGAL file in repository */

package www

import (
	"bytes"
	"strconv"
	"strings"

	"../assets"
	er "../error"
	"../log"
)

// JS store data for its method
type JS struct {
	imports []*assets.File
	inlined map[string]*assets.File
}

func (js *JS) checkInlined(srcJS *assets.File) (inlined bool) {
	_, inlined = js.inlined[srcJS.FullName]
	return
}

// AddJSToJS use to add a JS file to other!
func (js *JS) addJSToJS(srcJS, desJS *assets.File) {
	if js.checkInlined(srcJS) {
		return
	}
	js.inlined[srcJS.FullName] = srcJS
	desJS.Data = append(desJS.Data, srcJS.Data...)
}

// AddJSToJSRecursively use to add JSS and all import to JS file!
func (js *JS) addJSToJSRecursively(srcJS, desJS *assets.File) {
	if js.checkInlined(srcJS) {
		return
	}
	// Tell other this file will add to desJS later!
	js.inlined[srcJS.FullName] = srcJS

	var imports = make([]*assets.File, 0, 16)
	_ = js.extractJSImports(srcJS)
	for _, imp := range imports {
		js.addJSToJSRecursively(imp, desJS)
	}

	desJS.Data = append(desJS.Data, srcJS.Data...)
}

func (js *JS) extractJSImportsRecursively(jsFile *assets.File) (err *er.Error) {
	var lastImportIndex = len(js.imports)
	_ = js.extractJSImports(jsFile)
	for _, imp := range js.imports[lastImportIndex:] {
		err = js.extractJSImportsRecursively(imp)
	}
	return
}

func (js *JS) extractJSImports(jsFile *assets.File) (err *er.Error) {
	var importKeywordIndex, st, en int
	var loc, fileName string
	var locPart []string
	var imDep *assets.Folder
	var imFile *assets.File
	for {
		importKeywordIndex = bytes.Index(jsFile.Data, []byte("import "))
		if importKeywordIndex == -1 {
			break
		}
		// Find start and end of import file location!
		st = importKeywordIndex + bytes.IndexByte(jsFile.Data[importKeywordIndex:], '\'') + 1
		en = st + bytes.IndexByte(jsFile.Data[st:], '\'')
		loc = string(jsFile.Data[st:en])

		locPart = strings.Split(loc, "/")
		if len(locPart) < 2 {
			// don't parse dynamically import in files
			break
		} else {
			imDep = jsFile.Dep
			for i := 0; i < len(locPart)-1; i++ { // -1 due have file name at end of locPart
				switch locPart[i] {
				case ".":
					// noting to do!
				case "..":
					imDep = imDep.Dep
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
			if log.DevMode {
				log.Warn("WWW - ", fileName, "indicate as import in", jsFile.FullName, "can't find in repo")
			}
			// err =
			return
		}
		js.imports = append(js.imports, imFile)

		copy(jsFile.Data[importKeywordIndex-1:], jsFile.Data[en+1:])
		jsFile.Data = jsFile.Data[:len(jsFile.Data)-(en-importKeywordIndex)-2]
	}

	return
}

func localeAndMixJSFile(jsFile *assets.File) (files map[string]*assets.File) {
	var cssFile = jsFile.Dep.GetFile(jsFile.Name + ".css")
	if log.DevMode && cssFile == nil {
		log.Warn("WWW - ", jsFile.FullName, "Can't find CSS style file, Mix CSS to JS file skipped")
	}
	if cssFile != nil {
		jsFile = mixCSSToJS(jsFile, cssFile)
	}

	var lj = make(localize, 8)
	var jsonFile *assets.File = jsFile.Dep.GetFile(jsFile.Name + ".json")
	if log.DevMode && jsonFile == nil {
		log.Warn("WWW - ", jsFile.FullName, "Can't find JSON localize file, Localization skipped")
	}
	if jsonFile != nil {
		lj.jsonDecoder(jsonFile.Data)
	}

	files = localizeJSFile(jsFile, lj)

	var htmlFile = jsFile.Dep.GetFile(jsFile.Name + ".html")
	if log.DevMode && htmlFile == nil {
		log.Warn("WWW - ", jsFile.FullName, "Can't find HTML style file, Mix HTML to JS file skipped")
	}
	if htmlFile != nil {
		var htmlFiles = localizeHTMLFile(htmlFile, lj)
		for lang, f := range htmlFiles {
			mixHTMLToJS(files[lang], f)
		}
	}

	for _, f := range jsFile.Dep.FindFiles(jsFile.Name + "-template-") {
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
func localizeJSFile(file *assets.File, lj localize) (files map[string]*assets.File) {
	files = make(map[string]*assets.File, len(lj))
	if len(lj) == 0 {
		files[""] = file
	} else {
		for lang, text := range lj {
			files[lang] = replaceLocalizeTextInJS(file, text, lang)
		}
	}
	return
}

func replaceLocalizeTextInJS(js *assets.File, text []string, lang string) (newFile *assets.File) {
	newFile = js.Copy()
	newFile.Data = nil

	var jsData = js.Data

	var replacerIndex int
	var bracketIndex int
	var textIndex uint64
	var err error
	for {
		replacerIndex = bytes.Index(jsData, []byte("LocaleText["))
		if replacerIndex < 0 {
			newFile.Data = append(newFile.Data, jsData...)
			return
		}
		newFile.Data = append(newFile.Data, jsData[:replacerIndex]...)
		jsData = jsData[replacerIndex:]

		bracketIndex = bytes.IndexByte(jsData, ']')
		textIndex, err = strconv.ParseUint(string(jsData[11:bracketIndex]), 10, 8)
		if err != nil {
			log.Warn("WWW - ", js.FullName, "Index numbers is not valid, double check your file for bad structures")
		} else {
			newFile.Data = append(newFile.Data, text[textIndex]...)
		}

		jsData = jsData[bracketIndex+1:]
	}
}
