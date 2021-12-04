/* For license and copyright information please see LEGAL file in repository */

package www

import (
	"bytes"
	"strings"
	"text/template"
	"unsafe"
)

type MakeNewPageReq struct {
	Domain      string
	ScopeName   string
	PageVarName string
}

type MakeNewPageRes struct {
	JS, HTML, CSS, JSON []byte
}

// MakeNewPage generate 4 files for a ui page.
func MakeNewPage(req *MakeNewPageReq) (res *MakeNewPageRes, err error) {
	if req.PageVarName == "" {
		req.PageVarName = strings.ReplaceAll(strings.Title(req.ScopeName), "-", " ")
	}

	var jsBuf = new(bytes.Buffer)
	err = jsPageFile.Execute(jsBuf, req)
	if err != nil {
		return
	}
	res = &MakeNewPageRes{
		JS:   jsBuf.Bytes(),
		HTML: (*(*[]byte)(unsafe.Pointer(&htmlPageFile))),
		CSS:  (*(*[]byte)(unsafe.Pointer(&cssPageFile))),
		JSON: (*(*[]byte)(unsafe.Pointer(&jsonPageFile))),
	}

	return
}

var jsPageFile = template.Must(template.New("jsPageFile").Parse(`/* For license and copyright information please see LEGAL file in repository */

import '../libjs/widget-notification/force-leave-page.js'

const page{{.PageVarName}} = {
    URN: {
		URN: "urn:giti:{{.Domain}}:page:{{.ScopeName}}",
		ID: "",
        Domain: "{{.Domain}}",
        Scope: "page",
		Name: "{{.ScopeName}}",
	},
    Icon: "{{.ScopeName}}",
    Info: {
        Name: "LocaleText[0]",
        ShortName: "LocaleText[1]",
        Tagline: "LocaleText[2]",
        Slogan: "LocaleText[3]",
        Description: "LocaleText[4]",
        Tags: ["LocaleText[5]"]
    },
    Robots: "all",
    Related: ["", ""],
    HTML: () => ` + "``," + `
    CSS: '',
    Templates: {},
    Options: {},
	acceptedConditions: {
		"id": "",
		editable: true,
		offset: 0,
	},
	activeState: null,
}
Application.RegisterPage(page{{.PageVarName}})

// function init() {

// }

page{{.PageVarName}}.ActivatePage = async function (state) {
    // this.activeState = state
	// TODO::: Do any logic before page render
	window.document.body.innerHTML = this.HTML()
	// TODO::: Do any logic after page render
}

page{{.PageVarName}}.DeactivatePage = async function () {
	if (this.newList) {
        var forceLeave = await forceLeavePageWidget.ConnectedCallback()
    }
    if (forceLeave) {
		this.newList = null
		// call any widgets DisconnectedCallback e.g. barcodeCameraScannerWidget.DisconnectedCallback()
    }
    return forceLeave
}

page{{.PageVarName}}.OtherAction = async function () {}
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
