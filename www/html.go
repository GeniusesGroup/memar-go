/* For license and copyright information please see LEGAL file in repository */

package www

import (
	"bytes"
	"strconv"

	"../json"
	"../minify"
	"../protocol"
)

// mixHTMLToJS add given HTML file to specific part of JS file.
func mixHTMLToJS(js, html []byte) (mixedData []byte, err protocol.Error) {
	html, err = minify.HTML.MinifyBytes(html)
	if err != nil {
		return
	}

	var loc = bytes.Index(js, []byte("HTML: ("))
	if loc < 0 {
		loc = bytes.Index(js, []byte("HTML = ("))
		if loc < 0 {
			// err = "html file can't add to JS file due to bad JS format."
			return
		}
	}

	var graveAccentIndex int = bytes.IndexByte(js[loc:], '`')
	graveAccentIndex += loc + 1

	mixedData = make([]byte, 0, len(js)+len(html))
	mixedData = append(mixedData, js[:graveAccentIndex]...)
	mixedData = append(mixedData, html...)
	mixedData = append(mixedData, js[graveAccentIndex:]...)
	return
}

// mixHTMLTemplateToJS add given HTML template file to specific part of JS file.
func mixHTMLTemplateToJS(js, html []byte, tempName string) (mixedData []byte, err protocol.Error) {
	html, err = minify.HTML.MinifyBytes(html)
	if err != nil {
		return
	}

	var loc = bytes.Index(js, []byte(tempName+`": (`))
	if loc < 0 {
		// err = "html template file can't add to JS file due to bad JS format."
		return
	}

	var graveAccentIndex int = bytes.IndexByte(js[loc:], '`')
	graveAccentIndex += loc + 1

	mixedData = make([]byte, 0, len(js)+len(html))
	mixedData = append(mixedData, js[:graveAccentIndex]...)
	mixedData = append(mixedData, html...)
	mixedData = append(mixedData, js[graveAccentIndex:]...)
	return
}

// localizeHTMLFile make and returns number of localize file by number of language indicate in JSON localize text
func localizeHTMLFile(html []byte, lj localize) (mixedData map[string][]byte, err protocol.Error) {
	mixedData = make(map[string][]byte, len(lj))
	for lang, text := range lj {
		mixedData[lang], err = replaceLocalizeTextInHTML(html, text)
	}
	return
}

func replaceLocalizeTextInHTML(html []byte, text []string) (mixData []byte, err protocol.Error) {
	mixData = make([]byte, 0, len(html))
	var replacerIndex int
	var bracketIndex int
	var textIndex uint64
	var goErr error
	for {
		replacerIndex = bytes.Index(html, []byte("${LocaleText["))
		if replacerIndex < 0 {
			mixData = append(mixData, html...)
			return
		}
		mixData = append(mixData, html[:replacerIndex]...)
		html = html[replacerIndex:]

		bracketIndex = bytes.IndexByte(html, ']')
		textIndex, goErr = strconv.ParseUint(string(html[13:bracketIndex]), 10, 8)
		if goErr != nil {
			// err = "Index numbers in desire file is not valid, double check your file for bad structures"
		} else {
			mixData = append(mixData, text[textIndex]...)
		}

		html = html[bracketIndex+1:]
	}
}

type localize map[string][]string

func (lj *localize) jsonDecoder(data []byte) (err protocol.Error) {
	// TODO::: convert to generated code
	err = json.UnMarshal(data, lj)
	return
}
