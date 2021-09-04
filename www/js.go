/* For license and copyright information please see LEGAL file in repository */

package www

import (
	"bytes"
	"strconv"
	"strings"

	"../file"
	"../log"
	"../protocol"
)

// JS store data for its method
type JS struct {
	imports []protocol.File
	inlined map[string]protocol.File
}

func (js *JS) checkInlined(srcJS protocol.File) (inlined bool) {
	_, inlined = js.inlined[srcJS.MetaData().URI().Path()]
	return
}

// AddJSToJS use to add a JS file to other!
func (js *JS) addJSToJS(srcJS, desJS protocol.File) {
	if js.checkInlined(srcJS) {
		return
	}
	js.inlined[srcJS.MetaData().URI().Path()] = srcJS
	desJS.Data().Append(srcJS.Data().Marshal())
}

// AddJSToJSRecursively use to add JSS and all import to JS file!
func (js *JS) addJSToJSRecursively(srcJS, desJS protocol.File) {
	if js.checkInlined(srcJS) {
		return
	}
	// Tell other this file will add to desJS later!
	js.inlined[srcJS.MetaData().URI().Path()] = srcJS

	var imports = make([]protocol.File, 0, 16)
	_ = js.extractJSImports(srcJS)
	for _, imp := range imports {
		js.addJSToJSRecursively(imp, desJS)
	}
	desJS.Data().Append(srcJS.Data().Marshal())
}

func (js *JS) extractJSImportsRecursively(jsFile protocol.File) (err protocol.Error) {
	var lastImportIndex = len(js.imports)
	_ = js.extractJSImports(jsFile)
	for _, imp := range js.imports[lastImportIndex:] {
		err = js.extractJSImportsRecursively(imp)
	}
	return
}

func (js *JS) extractJSImports(jsFile protocol.File) (err protocol.Error) {
	var jsData = jsFile.Data().Marshal()
	var importKeywordIndex, st, en int
	for {
		importKeywordIndex = bytes.Index(jsData, []byte("import "))
		if importKeywordIndex == -1 {
			break
		}

		// TODO::: check and don't parse dynamically import

		// Find start and end of import file location!
		st = importKeywordIndex + bytes.IndexByte(jsData[importKeywordIndex:], '\'') + 1
		en = st + bytes.IndexByte(jsData[st:], '\'')
		var importPath = string(jsData[st:en])
		var importFile = file.FindByRelativeFrom(jsFile, importPath)
		if importFile == nil {
			if protocol.AppDevMode {
				log.Warn("WWW - '", importPath, "' indicate as import in", jsFile.MetaData().URI().Name(), "not found")
			}
			// err =
		} else {
			js.imports = append(js.imports, importFile)
		}

		copy(jsData[importKeywordIndex-1:], jsData[en+1:])
		jsData = jsData[:len(jsData)-(en-importKeywordIndex)-2]
	}
	jsFile.Data().UnMarshal(jsData)
	return
}

func localeAndMixJSFile(jsFile protocol.File) (mixedData map[string][]byte) {
	var js = jsFile.Data().Marshal()
	var jsFileNameWithoutExtension = jsFile.MetaData().URI().NameWithoutExtension()

	var lj = make(localize, 8)
	var jsonFile, _ = jsFile.ParentDirectory().File(jsFileNameWithoutExtension + ".json")
	if protocol.AppDevMode && jsonFile == nil {
		log.Warn("WWW - ", jsFile.MetaData().URI().Name(), "Can't find JSON localize file, Localization skipped")
		return
	}

	var cssFile, _ = jsFile.ParentDirectory().File(jsFileNameWithoutExtension + ".css")
	if protocol.AppDevMode && cssFile == nil {
		log.Warn("WWW - ", jsFile.MetaData().URI().Name(), "Can't find CSS style file, Mix CSS to JS file skipped")
	}
	if cssFile != nil {
		js, _ = mixCSSToJS(js, cssFile.Data().Marshal())
	}

	lj.jsonDecoder(jsonFile.Data().Marshal())
	mixedData, _ = localizeJSFile(js, lj)

	var htmlFile, _ = jsFile.ParentDirectory().File(jsFileNameWithoutExtension + ".html")
	if protocol.AppDevMode && htmlFile == nil {
		log.Warn("WWW - ", jsFile.MetaData().URI().Name(), "Can't find HTML style file, Mix HTML to JS file skipped")
	}
	if htmlFile != nil {
		var htmlMixedData, _ = localizeHTMLFile(htmlFile.Data().Marshal(), lj)
		for lang, html := range htmlMixedData {
			mixHTMLToJS(mixedData[lang], html)
		}
	}

	for _, f := range jsFile.ParentDirectory().FindFiles(jsFileNameWithoutExtension+"-template-", 0) {
		var namesPart = strings.Split(f.MetaData().URI().Name(), "-template-")
		var htmlMixedData, _ = localizeHTMLFile(f.Data().Marshal(), lj)
		for lang, f := range htmlMixedData {
			mixHTMLTemplateToJS(mixedData[lang], f, namesPart[1])
		}
	}

	return
}

// localizeJSFile make and returns number of localize file by number of language indicate in JSON localize text
func localizeJSFile(js []byte, lj localize) (mixedData map[string][]byte, err protocol.Error) {
	mixedData = make(map[string][]byte, len(lj))
	for lang, text := range lj {
		mixedData[lang], _ = replaceLocalizeTextInJS(js, text)
	}
	return
}

func replaceLocalizeTextInJS(js []byte, text []string) (mixData []byte, err protocol.Error) {
	mixData = make([]byte, 0, len(js))

	var replacerIndex int
	var bracketIndex int
	for {
		replacerIndex = bytes.Index(js, []byte("LocaleText["))
		if replacerIndex < 0 {
			mixData = append(mixData, js...)
			return
		}
		mixData = append(mixData, js[:replacerIndex]...)
		js = js[replacerIndex:]

		bracketIndex = bytes.IndexByte(js, ']')
		var textIndex, err = strconv.ParseUint(string(js[11:bracketIndex]), 10, 8)
		if err != nil {
			// err = "Index numbers in desire file is not valid, double check your file for bad structures"
		} else {
			mixData = append(mixData, text[textIndex]...)
		}

		js = js[bracketIndex+1:]
	}
}
