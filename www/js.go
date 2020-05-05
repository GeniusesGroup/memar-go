/* For license and copyright information please see LEGAL file in repository */

package www

import (
	"fmt"
	"os"
	"regexp"
	"strings"

	"../assets"
)

var jsComment = regexp.MustCompile(`(/\*.*?\*/)|(/\*[\w\W\n\s]+?\*/)`)

// MinifyJS use to minify js string!
func MinifyJS(js string) string {
	return jsComment.ReplaceAllString(js, "")
}

// AddHTMLToJS use to add minified HTML to related JS file in GUI architecture!
func addHTMLToJS(ass *assets.Folder, htmlFile *assets.File) {
	var htmlFileString string = string(htmlFile.Data)

	var n = strings.Split(htmlFile.Name, "-template-")

	// Get related js file
	var jsFile = ass.GetFile(n[0] + ".js")
	if jsFile == nil {
		return
	}
	var jsFileString string = string(jsFile.Data)

	// Minify html
	htmlFileString = MinifyHTML(htmlFileString)

	// add template to js
	var loc int
	if len(n) > 1 {
		var tempName = n[1]
		loc = strings.Index(jsFileString, `"`+tempName+`": (`)
	} else {
		loc = strings.Index(jsFileString, "HTML: (")
	}

	if loc < 0 {
		fmt.Fprintf(os.Stderr, "%v\n", "Following html file can't add to related JS file")
		fmt.Fprintf(os.Stderr, "%v\n", htmlFile.FullName)
		fmt.Fprintf(os.Stderr, "%v\n", n[0]+".js")
		return
	}

	jsFileString = jsFileString[:loc] + strings.Replace(jsFileString[loc:], "``", "`"+htmlFileString+"`", 1)
	jsFile.Data = []byte(jsFileString)
}

// AddCSSToJS use to add CSS to JS file!
func addCSSToJS(ass *assets.Folder, cssFile *assets.File) {
	var cssFileString string = string(cssFile.Data)

	// Get related js file
	var jsFile = ass.GetFile(cssFile.Name + ".js")
	if jsFile == nil {
		return
	}
	var jsFileString string = string(jsFile.Data)

	// Minify css
	cssFileString = MinifyCSS(cssFileString)

	// add page css
	jsFileString = strings.Replace(jsFileString, `CSS: "",`, `CSS: '`+cssFileString+`',`, 1)
	jsFile.Data = []byte(jsFileString)
}

// AddJSToJS use to add a JS file to other!
func addJSToJS(ass *assets.Folder, srcJS, desJS *assets.File, inlined map[string]*assets.File) {
	var ok bool
	_, ok = inlined[srcJS.FullName]
	if ok {
		// Inlined before!
		return
	}
	inlined[srcJS.FullName] = srcJS
	var srcJSString string = string(srcJS.Data)
	srcJSString = MinifyJS(srcJSString)
	desJS.DataString = desJS.DataString + srcJSString
}

// AddJSToJSRecursively use to add JSS and all import to JS file!
func addJSToJSRecursively(ass *assets.Folder, srcJS, desJS *assets.File, inlined map[string]*assets.File) {
	var ok bool
	_, ok = inlined[srcJS.FullName]
	if ok {
		// Inlined before!
		return
	}

	var srcJSString string = string(srcJS.Data)
	srcJSString = MinifyJS(srcJSString)
	// Tell other this file will add to desJS later!
	inlined[srcJS.FullName] = srcJS

	var im, st, en int
	var loc, depName, fileName string
	var locPart []string
	var imDep *assets.Folder
	var imFile *assets.File
	for {
		im = strings.Index(srcJSString, "import ")
		if im < 0 {
			break
		}
		// Find start and end of import file location!
		st = im + strings.IndexAny(srcJSString[im:], "'") + 1
		en = st + strings.IndexAny(srcJSString[st:], "'")
		loc = srcJSString[st:en]

		locPart = strings.Split(loc, "/")
		if len(locPart) < 2 {
			// don't parse dynamically import in files
			break
		} else if locPart[0] == "." {
			imDep = srcJS.Dep
		} else {
			depName = locPart[len(locPart)-2]
			imDep = ass.GetDependencyRecursively(depName)
			if imDep == nil {
				continue
			}
		}

		srcJSString = srcJSString[:im] + srcJSString[en+2:]

		fileName = locPart[len(locPart)-1]
		imFile = imDep.GetFile(fileName)

		if imFile != nil {
			addJSToJSRecursively(ass, imFile, desJS, inlined)
			inlined[imFile.FullName] = imFile
		}
	}

	desJS.DataString = desJS.DataString + srcJSString
}
