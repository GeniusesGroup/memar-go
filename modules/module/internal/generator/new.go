/* For license and copyright information please see the LEGAL file in the code repository */

package generator

import (
	"bytes"
	"go/format"

	error_p "memar/process/error/protocol"
)

type New struct {
	basePath    string
	packageName string

	bb bytes.Buffer

	filesData [][]byte
}

//memar:impl memar/computer/capsule/protocol.LifeCycle
func (g *New) Init(basePath string) (err error_p.Error) {
	g.basePath = basePath

	g.filesData = make([][]byte, 0, len(g.FilesPath()))

	// TODO::: ready strings

	err = g.makeFiles()
	if err != nil {
		return
	}

	return
}

func (g *New) DirsPath() []string {
	return []string{
		g.basePath + "/errors",
		g.basePath + "/internal/storage",
		g.basePath + "/pages",
		g.basePath + "/protocol",
		g.basePath + "/services",
		g.basePath + "/widgets",
	}
}

func (g *New) FilesPath() []string {
	return []string{
		g.basePath + "/module-cmd.go",
		g.basePath + "/module-gui.go",
		g.basePath + "/module-server.go",
		g.basePath + "/module.go",
		g.basePath + "/locale.eng.go",

		g.basePath + "/protocol/data_model.go",
		// g.basePath + "/protocol/bs-.go",
		// g.basePath + "/protocol/dt-.go",
	}
}

func (g *New) FilesData() [][]byte {
	return g.filesData
}

func (g *New) makeFiles() (err error_p.Error) {
	err = g.makeFile_ModuleCMD()
	if err != nil {
		return
	}
	err = g.makeFile_ModuleGUI()
	if err != nil {
		return
	}
	err = g.makeFile_ModuleServer()
	if err != nil {
		return
	}
	err = g.makeFile_Module()
	if err != nil {
		return
	}
	err = g.makeFile_LocaleENG()
	if err != nil {
		return
	}

	err = g.makeFile_Protocol_DataModel()
	if err != nil {
		return
	}

	return
}

func (g *New) makeFile_ModuleCMD() (err error_p.Error) {
	var goErr error
	var fileData []byte

	goErr = temp_file_ModuleCMD.Execute(&g.bb, g)
	if goErr != nil {
		// err =
		return
	}
	fileData, goErr = format.Source(g.bb.Bytes())
	if goErr != nil {
		// err =
		return
	}
	g.filesData = append(g.filesData, fileData)
	g.bb.Reset()
	return
}

func (g *New) makeFile_ModuleGUI() (err error_p.Error) {
	var goErr error
	var fileData []byte

	goErr = temp_file_ModuleGUI.Execute(&g.bb, g)
	if goErr != nil {
		// err =
		return
	}
	fileData, goErr = format.Source(g.bb.Bytes())
	if goErr != nil {
		// err =
		return
	}
	g.filesData = append(g.filesData, fileData)
	g.bb.Reset()
	return
}

func (g *New) makeFile_ModuleServer() (err error_p.Error) {
	var goErr error
	var fileData []byte

	goErr = temp_file_ModuleServer.Execute(&g.bb, g)
	if goErr != nil {
		// err =
		return
	}
	fileData, goErr = format.Source(g.bb.Bytes())
	if goErr != nil {
		// err =
		return
	}
	g.filesData = append(g.filesData, fileData)
	g.bb.Reset()
	return
}

func (g *New) makeFile_Module() (err error_p.Error) {
	var goErr error
	var fileData []byte

	goErr = temp_file_Module.Execute(&g.bb, g)
	if goErr != nil {
		// err =
		return
	}
	fileData, goErr = format.Source(g.bb.Bytes())
	if goErr != nil {
		// err =
		return
	}
	g.filesData = append(g.filesData, fileData)
	g.bb.Reset()
	return
}

func (g *New) makeFile_LocaleENG() (err error_p.Error) {
	var goErr error
	var fileData []byte

	goErr = temp_file_LocaleENG.Execute(&g.bb, g)
	if goErr != nil {
		// err =
		return
	}
	fileData, goErr = format.Source(g.bb.Bytes())
	if goErr != nil {
		// err =
		return
	}
	g.filesData = append(g.filesData, fileData)
	g.bb.Reset()
	return
}

func (g *New) makeFile_Protocol_DataModel() (err error_p.Error) {
	var goErr error
	var fileData []byte

	goErr = temp_file_Protocol_DataModel.Execute(&g.bb, g)
	if goErr != nil {
		// err =
		return
	}
	fileData, goErr = format.Source(g.bb.Bytes())
	if goErr != nil {
		// err =
		return
	}
	g.filesData = append(g.filesData, fileData)
	g.bb.Reset()
	return
}
