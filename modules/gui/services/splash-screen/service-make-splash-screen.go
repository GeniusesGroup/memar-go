/* For license and copyright information please see the LEGAL file in the code repository */

package www

import (
	"bytes"
	"text/template"

	"memar/protocol"
)

// MakeSplashFiles make splash screen that use as landing page or first screen user see when open GUI app.
func MakeSplashFiles(html, css, json protocol.File) (err error) {
	html.Rename("splash.html")
	var htmlBuf = new(bytes.Buffer)
	if err = splashHTMLFileTemplate.Execute(htmlBuf, ""); err != nil {
		return
	}
	html.Data().Unmarshal(htmlBuf.Bytes())

	css.Rename("splash.css")
	var cssBuf = new(bytes.Buffer)
	if err = splashCSSFileTemplate.Execute(cssBuf, ""); err != nil {
		return
	}
	css.Data().Unmarshal(cssBuf.Bytes())

	json.Rename("splash.json")
	var jsonBuf = new(bytes.Buffer)
	if err = splashJSONFileTemplate.Execute(jsonBuf, ""); err != nil {
		return
	}
	json.Data().Unmarshal(jsonBuf.Bytes())
	return
}

var splashHTMLFileTemplate = template.Must(template.New("splashHTMLFile").Parse(`
<!-- For license and copyright information please see the LEGAL file in the code repository -->

<main>
    <div>
        <img src="/images/app-icon-128x128.png" alt="Platform logo" />
        <p>
            POWERED BY <br />
            <a href="/cloud">SABZCITY PLATFORM</a>
        </p>

        <noscript>Please Enable Javascript to be able to use this web app.</noscript>
    </div>
</main>
`))

var splashCSSFileTemplate = template.Must(template.New("splashCSSFile").Parse(`
/* For license and copyright information please see the LEGAL file in the code repository */

main {
    position: fixed;
    width: 100%;
    height: 100%;
    top: 0;
    left: 0;
    background: rgba(0, 0, 0, 0.5)
}

div {
    position: absolute;
    height: 200px;
    max-width: 400px;
    top: 0;
    bottom: 0;
    right: 0;
    left: 0;
    margin: auto;
    text-align: center;
    background: #ffffff;
    box-shadow: 0 2px 4px -1px rgba(0, 0, 0, 0.2), 0 4px 5px 0 rgba(0, 0, 0, 0.14), 0 1px 10px 0 rgba(0, 0, 0, 0.12);
}

/* img {} */

p {
    color: #92989b;
    font-family: sans-serif;
    font-size: 9pt
}

a {
    color: #616161;
    text-decoration: none
}
`))

var splashJSONFileTemplate = template.Must(template.New("splashJSONFile").Parse(`
{
	"en": []
}
`))
