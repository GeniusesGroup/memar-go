/* For license and copyright information please see LEGAL file in repository */

package www

import (
	"bytes"
	"strconv"

	"../file"
	"../giti"
	"../json"
	"../log"
)

// mixHTMLToJS add given HTML file to specific part of JS file.
func mixHTMLToJS(jsFile, htmlFile *file.File) (mixFile *file.File) {
	htmlFile.Minify()
	mixFile = jsFile.Copy()
	var jsData = jsFile.Data()
	var htmlData = htmlFile.Data()

	var loc = bytes.Index(jsData, []byte("HTML: ("))
	if loc < 0 {
		loc = bytes.Index(jsData, []byte("HTML = ("))
		if loc < 0 {
			log.Warn(htmlFile.FullName, "html file can't add to", jsFile.FullName, "file due to bad JS.")
			return
		}
	}

	var graveAccentIndex int = bytes.IndexByte(jsData[loc:], '`')
	graveAccentIndex += loc + 1

	var mixFileData = make([]byte, 0, len(jsData)+len(htmlData))
	mixFileData = append(mixFileData, jsData[:graveAccentIndex]...)
	mixFileData = append(mixFileData, htmlData...)
	mixFileData = append(mixFileData, jsData[graveAccentIndex:]...)

	mixFile.SetData(mixFileData)
	return
}

// mixHTMLTemplateToJS add given HTML template file to specific part of JS file.
func mixHTMLTemplateToJS(jsFile, htmlFile *file.File, tempName string) (mixFile *file.File) {
	htmlFile.Minify()
	mixFile = jsFile.Copy()
	var jsData = jsFile.Data()
	var htmlData = htmlFile.Data()

	var loc = bytes.Index(jsData, []byte(tempName+`": (`))
	if loc < 0 {
		log.Warn(htmlFile.FullName, "html template file can't add to", jsFile.FullName, "file due to bad JS template.")
		return
	}

	var graveAccentIndex int = bytes.IndexByte(jsData[loc:], '`')
	graveAccentIndex += loc + 1

	var mixFileData = make([]byte, 0, len(jsData)+len(htmlData))
	mixFileData = append(mixFileData, jsData[:graveAccentIndex]...)
	mixFileData = append(mixFileData, htmlData...)
	mixFileData = append(mixFileData, jsData[graveAccentIndex:]...)

	mixFile.SetData(mixFileData)
	return
}

// localizeHTMLFile make and returns number of localize file by number of language indicate in JSON localize text
func localizeHTMLFile(htmlFile *file.File, lj localize) (files map[string]*file.File) {
	htmlFile.Minify()

	files = make(map[string]*file.File, len(lj))
	if len(lj) == 0 {
		files[""] = htmlFile
	} else {
		for lang, text := range lj {
			files[lang] = replaceLocalizeTextInHTML(htmlFile, text, lang)
		}
	}
	return
}

func replaceLocalizeTextInHTML(html *file.File, text []string, lang string) (mixFile *file.File) {
	mixFile = html.Copy()
	var htmlData = html.Data()
	var mixFileData = make([]byte, 0, len(htmlData))

	var replacerIndex int
	var bracketIndex int
	var textIndex uint64
	var err error
	for {
		replacerIndex = bytes.Index(htmlData, []byte("${LocaleText["))
		if replacerIndex < 0 {
			mixFileData = append(mixFileData, htmlData...)
			mixFile.SetData(mixFileData)
			return
		}
		mixFileData = append(mixFileData, htmlData[:replacerIndex]...)
		htmlData = htmlData[replacerIndex:]

		bracketIndex = bytes.IndexByte(htmlData, ']')
		textIndex, err = strconv.ParseUint(string(htmlData[13:bracketIndex]), 10, 8)
		if err != nil {
			log.Warn("Index numbers in", html.FullName, "is not valid, double check your file for bad structures")
		} else {
			mixFileData = append(mixFileData, text[textIndex]...)
		}

		htmlData = htmlData[bracketIndex+2:]
	}
}

type localize map[string][]string

func (lj *localize) jsonDecoder(data []byte) (err giti.Error) {
	// TODO::: convert to generated code
	err = json.UnMarshal(data, lj)
	return
}
