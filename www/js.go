/* For license and copyright information please see LEGAL file in repository */

package www

import (
	"bytes"
	"strconv"
	"strings"

	"../assets"
	"../log"
)

// AddJSToJS use to add a JS file to other!
func addJSToJS(ass *assets.Folder, srcJS, desJS *assets.File, inlined map[string]*assets.File) {
	var ok bool
	_, ok = inlined[srcJS.FullName]
	if ok {
		// Inlined before!
		return
	}
	inlined[srcJS.FullName] = srcJS
	desJS.Data = append(desJS.Data, srcJS.Data...)
}

// AddJSToJSRecursively use to add JSS and all import to JS file!
func addJSToJSRecursively(ass *assets.Folder, srcJS, desJS *assets.File, inlined map[string]*assets.File) {
	var ok bool
	_, ok = inlined[srcJS.FullName]
	if ok {
		// Inlined before!
		return
	}
	// Tell other this file will add to desJS later!
	inlined[srcJS.FullName] = srcJS

	var im, st, en int
	var loc, depName, fileName string
	var locPart []string
	var imDep *assets.Folder
	var imFile *assets.File
	for {
		im = bytes.Index(srcJS.Data, []byte("import "))
		if im == -1 {
			break
		}
		// Find start and end of import file location!
		st = im + bytes.IndexByte(srcJS.Data[im:], '\'') + 1
		en = st + bytes.IndexByte(srcJS.Data[st:], '\'')
		loc = string(srcJS.Data[st:en])

		locPart = strings.Split(loc, "/")
		if len(locPart) < 2 {
			// don't parse dynamically import in files
			break
		} else if len(locPart) == 2 && locPart[0] == "." {
			imDep = srcJS.Dep
		} else {
			depName = locPart[len(locPart)-2]
			imDep = ass.GetDependencyRecursively(depName)
			if imDep == nil {
				continue
			}
		}

		copy(srcJS.Data[im-1:], srcJS.Data[en+1:])
		srcJS.Data = srcJS.Data[:len(srcJS.Data)-(en-im)-2]
		// srcJSString = srcJSString[:im] + srcJSString[en+2:]

		fileName = locPart[len(locPart)-1]
		imFile = imDep.GetFile(fileName)
		if imFile != nil {
			addJSToJSRecursively(ass, imFile, desJS, inlined)
			inlined[imFile.FullName] = imFile
		}
	}

	desJS.Data = append(desJS.Data, srcJS.Data...)
}

// localizeJSFile make and returns number of localize file by number of language indicate in JSON localize text
func localizeJSFile(file *assets.File, lj localize) (files map[string]*assets.File) {
	files = make(map[string]*assets.File, len(lj))
	for lang, text := range lj {
		files[lang] = replaceLocalizeTextInJS(file, text, lang)
	}
	return
}

func replaceLocalizeTextInJS(js *assets.File, text []string, lang string) (newFile *assets.File) {
	newFile = js.Copy()
	newFile.Name += "-" + lang
	newFile.FullName = newFile.Name + "." + newFile.Extension
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
			log.Warn("Index numbers in", js.FullName, "is not valid, double check your file for bad structures")
		} else {
			newFile.Data = append(newFile.Data, text[textIndex]...)
		}

		jsData = jsData[bracketIndex+1:]
	}
}
