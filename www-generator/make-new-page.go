/* For license and copyright information please see LEGAL file in repository */

package wg

import (
	"bytes"
	"mime"
	"text/template"
	"unsafe"

	"../assets"
)

// MakeNewPage generate 4 files for a page that page name will be js.Name
func MakeNewPage(js, html, css, json *assets.File) (err error) {
	js.FullName = js.Name + ".js"
	js.Extension = "js"
	js.MimeType = mime.TypeByExtension("js")
	js.State = assets.StateChanged
	var jsBuf = new(bytes.Buffer)
	err = jsPageFile.Execute(jsBuf, js.Name)
	if err != nil {
		return
	}
	js.Data = jsBuf.Bytes()

	html.FullName = js.Name + ".html"
	html.Name = js.Name
	html.Extension = "html"
	html.MimeType = mime.TypeByExtension("html")
	html.State = assets.StateChanged
	html.Data = *(*[]byte)(unsafe.Pointer(&htmlPageFile))

	css.FullName = js.Name + ".css"
	css.Name = js.Name
	css.Extension = "css"
	css.MimeType = mime.TypeByExtension("css")
	css.State = assets.StateChanged
	css.Data = *(*[]byte)(unsafe.Pointer(&cssPageFile))

	json.FullName = js.Name + ".json"
	json.Name = js.Name
	json.Extension = "json"
	json.MimeType = mime.TypeByExtension("json")
	json.State = assets.StateChanged
	json.Data = *(*[]byte)(unsafe.Pointer(&jsonPageFile))

	return
}

var jsPageFile = template.Must(template.New("jsPageFile").Parse(`/* For license and copyright information please see LEGAL file in repository */

Application.Pages["{{.}}"] = {
    ID: "{{.}}",
    RecordID: null,
    Condition: {},
    State: "",
    Robots: "all",
    Info: {
        Name: "LocaleText[0]",
        ShortName: "LocaleText[1]",
        Tagline: "LocaleText[2]",
        Slogan: "LocaleText[3]",
        Description: "LocaleText[4]",
        Tags: ["LocaleText[5]"]
    },
    Icon: "{{.}}",
    Related: ["", ""],
    HTML: () => ` + "``," + `
    CSS: '',
    Templates: {},
}

Application.Pages["{{.}}"].ConnectedCallback = function () {
    window.document.body.innerHTML = this.HTML()

    // Application.Widgets["hamburger-menu"].ConnectedCallback()
    // Application.Widgets["user-menu"].ConnectedCallback()
    // Application.Widgets["service-menu"].ConnectedCallback()
}

Application.Pages["{{.}}"].DisconnectedCallback = function () {
}
`))

var htmlPageFile = `<!-- For license and copyright information please see LEGAL file in repository -->

<header>
</header>

<main>
    <header>
    </header>

    <footer>
    </footer>
</main>`

var cssPageFile = `/* For license and copyright information please see LEGAL file in repository */
`

var jsonPageFile = `{
    "en": [
        "",
        "",
        "",
        "",
        "",
        "",
        ""
    ],
    "fa": [
        "",
        "",
        "",
        "",
        "",
        "",
        ""
    ]
}`
