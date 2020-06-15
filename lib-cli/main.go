/* For license and copyright information please see LEGAL file in repository */

package main

import (
	"fmt"
	"path"
	"runtime"
	"time"

	ag "../Achaemenid-generator"
	gg "../Ganjine-generator"
	"../assets"
	"../syllab"
	wg "../www"
)

const (
	// version must update in each release!
	version = "v0.6.7"
)

// Some mutable folder name!
const (
	FolderNameAPIs = "apis"
	FolderNameDB   = "db"
	FolderNameGUI  = "gui"

	FolderNameServices  = "services"
	FolderNameDataStore = "datastore"

	FolderNameGGPages     = "pages"
	FolderNameGUILandings = "landings"
	FolderNameGUIWidgets  = "widgets"

	FolderNameJSSDK = "sdk-js"
	FolderNameGOSDK = "sdk-go"
)

var (
	// ServiceRootLocation is location of repository root folder.
	ServiceRootLocation string
)

func init() {
	// Indicate ServiceRootLocation
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		buildLog("No caller information, So we can't specify service root location")
		return
	}
	ServiceRootLocation = path.Dir(path.Dir(path.Dir(filename)))
}

func main() {
	defer saveLog()

	// Add some generic data to first of file.
	buildLog("Generator version:", version)

	// print contact information to code-generate.log
	buildLog("You may help us and create issue:")
	buildLog("https://github.com/SabzCity/libgo/issues/new")
	buildLog("For more information, see:")
	buildLog("https://github.com/SabzCity/libgo/")

	start := time.Now()
	buildLog("Code generate start at", start)

	buildLog("Service root location is at", ServiceRootLocation)

	// Parse repository
	var repo = assets.NewFolder("")
	var err error
	err = repo.ReadRepositoryFromFileSystem(ServiceRootLocation)
	if err != nil {
		buildLog("Read repository face this error:", err)
	}

	buildLog("Enter desire chaparkhane CLI service ID. You can select:")
	buildLog("0  : Nothing DO!!! Prevent mistakes!!!")
	buildLog("1  : Quit without save changes")
	buildLog("2  : Save changes and quit")
	buildLog("3  : Save changes without quit")
	buildLog("*************** Common Services *************** ")
	buildLog("10  : Add project template to repository")
	buildLog(" *************** Achaemenid Services *************** ")
	buildLog("30  : Add new Achaemenid service file to apis/services folder")
	buildLog("31  : Update exiting Achaemenid file in apis/services folder")
	buildLog("32  : Make Achaemenid service GO-SDK")
	buildLog("33  : Make Achaemenid service JS-SDK")
	buildLog("40  : Make www assets file from gui folder")
	buildLog(" *************** Ganjine Services *************** ")
	buildLog("50  : Add new Ganjine file to apis/datastore folder")
	buildLog("51  : Update exiting Ganjine file in apis/datastore folder")
	buildLog(" *************** Syllab Services *************** ")
	buildLog("70  : Update Syllab encoder||decoder methods in given file name by safe manner")
	buildLog("71  : Update Syllab encoder||decoder methods in given file name by unsafe manner")
	buildLog(" *************** GUI Services *************** ")
	buildLog(" *************** JSON Services *************** ")
	buildLog("----------------------------------------")
Choose:
	var requestedService int
	fmt.Scanln(&requestedService)
	buildLog("You choose ", requestedService)

	switch requestedService {
	case 0:
		buildLog("Nothing DO in this ID to prevent often mistakes enter multi times!!!")
		goto Choose
	case 1:
		buildLog("See you soon!")
		goto Exit
	case 2:
		err = repo.WriteRepositoryToFileSystem(ServiceRootLocation)
		if err != nil {
			buildLog("Unable to write to repo:", err)
		}
		buildLog("All changes write to disk as you desire!")
		buildLog("See you soon!")
		goto Exit
	case 3:
		err = repo.WriteRepositoryToFileSystem(ServiceRootLocation)
		if err != nil {
			buildLog("Unable to write to repo:", err)
		}
		buildLog("All changes write to disk as you desire!")

	// *************** Common Services ***************
	case 10:
		_, err := MakeNewProject(&MakeNewProjectReq{Repo: repo})
		if err != nil {
			buildLog("Add project template to repository face this error:", err)
		}
		buildLog("Add project template had been succeed!!")
	case 11:
		// res, err := ag.UpdateProjectTemplate(&ReqUpdateProjectTemplate002{readRepo})
		// if err != nil {
		// 	buildLog("Update project template error:", err)
		// }
		// repo = res.Repo

	// *************** Achaemenid Services ***************
	case 30:
		buildLog("Enter desire service name in ```kebab-case``` like ```register-new-person```")
		var serviceName string
		fmt.Scanln(&serviceName)
		buildLog("Desire name: ", serviceName)

		var file = assets.File{
			Name: serviceName,
		}
		err = ag.MakeNewServiceFile(&file)
		if err != nil {
			buildLog("Add new Achaemenid service template face this error:", err)
			break
		}
		repo.Dependencies[FolderNameAPIs].Dependencies[FolderNameServices].SetFile(&file)
		buildLog("Add new Achaemenid service had been succeed!!\n")
	case 31:
		buildLog("Enter desire full file name with extension!")
		var fileName string
		fmt.Scanln(&fileName)
		buildLog("Desire file name: ", fileName)

		var file = repo.GetFileRecursively(fileName)
		if file == nil {
			buildLog("Desire file name not exist in repo!!")
			break
		}

		err = ag.UpdateServiceFile(file)
		if err != nil {
			buildLog("Update Achaemenid service file face this error:", err)
		}
		buildLog("Update exiting Achaemenid file had been succeed!!\n")
	case 32:
		buildLog("Enter desire full file name with extension!")
		var fileName string
		fmt.Scanln(&fileName)
		buildLog("Desire file name: ", fileName)

		var file = repo.GetFileRecursively(fileName)
		if file == nil {
			buildLog("Desire file name not exist in repo!!")
			break
		}

		file, err = ag.MakeGoSDK(file)
		if err != nil {
			buildLog("Make Achaemenid service GO-SDK face this error:", err)
		}
		repo.Dependencies[FolderNameGOSDK].SetFile(file)
		buildLog("Make Achaemenid service Go-SDK had been succeed!!\n")
	case 33:
		buildLog("Enter desire full file name with extension!")
		var fileName string
		fmt.Scanln(&fileName)
		buildLog("Desire file name: ", fileName)

		var file = repo.GetFileRecursively(fileName)
		if file == nil {
			buildLog("Desire file name not exist in repo!!")
			break
		}

		file, err = ag.MakeJSSDK(file)
		if err != nil {
			buildLog("Make Achaemenid service JS-SDK face this error:", err)
		}
		repo.Dependencies[FolderNameJSSDK].SetFile(file)
		buildLog("Make Achaemenid service JS-SDK had been succeed!!\n")
	case 40:
		var file = assets.File{}
		err = wg.MakeAssetsFile(repo, &file)
		if err != nil {
			buildLog("Make www assets file from gui folder face this error:", err)
			break
		}
		repo.SetFile(&file)
		buildLog("Make www assets file from gui folder had been succeed!!\n")

	// *************** Ganjine Services ***************
	case 50:
		buildLog("Enter desire structure name in ```kebab-case``` like ```person-authentication```")
		var sName string
		fmt.Scanln(&sName)
		buildLog("Desire name: ", sName)

		var file = assets.File{
			Name: sName,
		}
		err = gg.MakeNewDatastoreFile(&file)
		if err != nil {
			buildLog("Add new Ganjine file template face this error:", err)
			break
		}
		repo.Dependencies[FolderNameAPIs].Dependencies[FolderNameDataStore].SetFile(&file)
		buildLog("Add new structure had been succeed!!\n")
	case 51:
		buildLog("Enter desire full file name with extension!")
		var fileName string
		fmt.Scanln(&fileName)
		buildLog("Desire file name: ", fileName)

		var file = repo.GetFileRecursively(fileName)
		if file == nil {
			buildLog("Desire file name not exist in repo!!")
			break
		}

		err = gg.UpdateDatastoreFile(file)
		if err != nil {
			buildLog("Update Structure file face this error:", err)
		}
		buildLog("Update exiting Ganjine file had been succeed!!\n")

	// *************** Syllab Services ***************
	case 70:
		buildLog("Enter desire full file name with extension!")
		var fileName string
		fmt.Scanln(&fileName)
		buildLog("Desire file name: ", fileName)

		var file = repo.GetFileRecursively(fileName)
		if file == nil {
			buildLog("Desire file name not exist in repo!!")
			break
		}

		err = syllab.CompleteEncoderMethodSafe(file)
		if err != nil {
			buildLog("Update Syllab encoder||decoder safe face this error:", err)
		}
		buildLog("Update Syllab encoder||decoder safe had been succeed!!\n")
	case 71:
		buildLog("Enter desire full file name with extension!")
		var fileName string
		fmt.Scanln(&fileName)
		buildLog("Desire file name: ", fileName)

		var file = repo.GetFileRecursively(fileName)
		if file == nil {
			buildLog("Desire file name not exist in repo!!")
			break
		}

		err = syllab.CompleteEncoderMethodUnsafe(file)
		if err != nil {
			buildLog("Update Syllab encoder||decoder unsafe face this error:", err)
		}
		buildLog("Update Syllab encoder||decoder unsafe had been succeed!!\n")

	default:
		buildLog("Nothing DO in given ID to prevent often mistakes enter bad ID!!!")
		goto Choose
	}

	buildLog("----------------------------------------")
	buildLog("Enter new desire chaparkhane service ID:")
	goto Choose

Exit:
	defer buildLog("CLI app run duration:", time.Since(start))
}
