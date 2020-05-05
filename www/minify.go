/* For license and copyright information please see LEGAL file in repository */

package www

import (
	"regexp"
)

var newLine = regexp.MustCompile(`\r?\n`)
var htmlComment = regexp.MustCompile(`(<!--.*?-->)|(<!--[\w\W\n\s]+?-->)`)

// MinifyHTML use to minify html string!
func MinifyHTML(html string) string {
	html = newLine.ReplaceAllString(html, "")
	html = htmlComment.ReplaceAllString(html, "")
	return html
}

// MinifyCSS use to minify css string!
func MinifyCSS(css string) string {
	css = newLine.ReplaceAllString(css, "")
	return css
}

