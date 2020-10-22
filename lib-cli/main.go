/* For license and copyright information please see LEGAL file in repository */

package main

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	ag "../Achaemenid-generator"
	gg "../Ganjine-generator"
	"../assets"
	"../json"
	"../log"
	"../syllab"
	wg "../www-generator"
)

const (
	// it is better to update version in each release!
	version = "v0.7.6"
)

// Some mutable folder name!
const (
	FolderNameDataStore = "datastore"
	FolderNameGUI       = "gui"
	FolderNameServices  = "services"
	FolderNameSecret    = "secret"

	FolderNameGUIPages    = "pages"
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
	var ex, err = os.Executable()
	if err != nil {
		log.Fatal(err)
		return
	}
	ServiceRootLocation = filepath.Dir(ex)

	log.Init("lib-cli", ServiceRootLocation, 10*60)
}

// TODO::: improve architecture by split main function!
func main() {
	defer log.SaveToStorage()

	// Add some generic data to first of file.
	log.Info("Generator version:", version)

	// print contact information to code-generate.log
	log.Info("You may help us and create issue:")
	log.Info("https://github.com/SabzCity/libgo/issues/new")
	log.Info("For more information, see:")
	log.Info("https://github.com/SabzCity/libgo/")

	start := time.Now()
	log.Info("Code generate start at", start)

	log.Info("Service root location is at", ServiceRootLocation)

	// Parse repository
	var repo = assets.NewFolder("")
	var err error
	err = repo.ReadRepositoryFromFileSystem(ServiceRootLocation, false)
	if err != nil {
		log.Fatal("Read repository face this error:", err)
	}

	log.Info("-------------------------------------------------------------------------------------")
	log.Info("****************************** Available CLI services *******************************")
	log.Info("-------------------------------------------------------------------------------------")
	log.Info("0  : Nothing DO!!! Prevent mistakes!!!")
	log.Info("1  : Quit without save changes")
	log.Info("2  : Save changes and quit")
	log.Info("3  : Save changes without quit")
	log.Info("*************** Common Services *************** ")
	log.Info("10  : Add project template to repository")
	log.Info(" *************** Achaemenid Services *************** ")
	log.Info("30  : Add new Achaemenid service file to ./services/ folder")
	log.Info("31  : Update exiting Achaemenid file in ./services/ folder")
	log.Info("32  : Make Achaemenid service GO-SDK")
	log.Info("33  : Make Achaemenid service JS-SDK")
	log.Info("40  : Make www assets file from gui folder")
	log.Info(" *************** Ganjine Services *************** ")
	log.Info("50  : Add new Ganjine file to ./datastore/ folder")
	log.Info("51  : Update exiting Ganjine file in ./datastore/ folder")
	log.Info(" *************** GUI Services *************** ")
	log.Info("70  : Add new GUI raw page")
	log.Info(" *************** Syllab Services *************** ")
	log.Info("80  : Update Syllab encoder||decoder methods in given file name")
	log.Info(" *************** JSON Services *************** ")
	log.Info("90  : Update JSON encoder||decoder methods in given file name")

Choose:
	log.Info("-------------------------------------------------------------------------------------")
	log.Info("Enter desire CLI service ID:")
	var requestedService uint
	_, err = fmt.Scanln(&requestedService)
	if err != nil {
		log.Warn("Unable to parse given ID:", err)
		goto Choose
	}
	log.Info("You choose:", requestedService)

	switch requestedService {
	case 0:
		log.Info("Nothing DO in this ID to prevent often mistakes enter multi times!!!")
		goto Choose
	case 1:
		goto Exit
	case 2:
		err = repo.WriteRepositoryToFileSystem(ServiceRootLocation)
		if err != nil {
			log.Warn("Unable to write to repo:", err)
		}
		log.Info("All changes write to disk as you desire!")
		goto Exit
	case 3:
		err = repo.WriteRepositoryToFileSystem(ServiceRootLocation)
		if err != nil {
			log.Warn("Unable to write to repo:", err)
		}
		log.Info("All changes write to disk as you desire!")

	// *************** Common Services ***************
	case 10:
		_, err := MakeNewProject(&MakeNewProjectReq{Repo: repo})
		if err != nil {
			log.Warn("Add project template to repository face this error:", err)
		}
		log.Info("Add project template had been succeed!!")
	case 11:
		// res, err := ag.UpdateProjectTemplate(&ReqUpdateProjectTemplate002{readRepo})
		// if err != nil {
		// 	log.Warn("Update project template error:", err)
		// }
		// repo = res.Repo

	// *************** Achaemenid Services ***************
	case 30:
		log.Info("Enter desire service name in ```kebab-case``` like ```register-new-person```")
		var serviceName string
		fmt.Scanln(&serviceName)
		log.Info("Desire name: ", serviceName)

		var file = assets.File{
			Name: serviceName,
		}
		err = ag.MakeNewServiceFile(&file)
		if err != nil {
			log.Warn("Add new Achaemenid service template face this error:", err)
			break
		}
		repo.Dependencies[FolderNameServices].SetFile(&file)
		log.Info("Add new Achaemenid service had been succeed!!")
	case 31:
		log.Info("Enter desire full file name with extension!")
		var fileName string
		fmt.Scanln(&fileName)
		log.Info("Desire file name: ", fileName)

		var file = repo.GetFileRecursively(fileName)
		if file == nil {
			log.Warn("Desire file name not exist in repo!!")
			break
		}

		err = ag.UpdateServiceFile(file)
		if err != nil {
			log.Warn("Update Achaemenid service file face this error:", err)
		}
		log.Info("Update exiting Achaemenid file had been succeed!!")
	case 32:
		log.Info("Enter desire full file name with extension!")
		var fileName string
		fmt.Scanln(&fileName)
		log.Info("Desire file name: ", fileName)

		var file = repo.GetFileRecursively(fileName)
		if file == nil {
			log.Warn("Desire file name not exist in repo!!")
			break
		}

		file, err = ag.MakeGoSDK(file)
		if err != nil {
			log.Warn("Make Achaemenid service GO-SDK face this error:", err)
		}
		repo.Dependencies[FolderNameGOSDK].SetFile(file)
		log.Info("Make Achaemenid service Go-SDK had been succeed!!")
	case 33:
		log.Info("Enter desire full file name with extension!")
		var fileName string
		fmt.Scanln(&fileName)
		log.Info("Desire file name: ", fileName)

		var file = repo.GetFileRecursively(fileName)
		if file == nil {
			log.Warn("Desire file name not exist in repo!!")
			break
		}

		file, err = ag.MakeJSSDK(file)
		if err != nil {
			log.Warn("Make Achaemenid service JS-SDK face this error:", err)
		}
		repo.Dependencies[FolderNameJSSDK].SetFile(file)
		log.Info("Make Achaemenid service JS-SDK had been succeed!!")
	case 40:
		var file = assets.File{}
		err = wg.MakeAssetsFile(repo, &file)
		if err != nil {
			log.Warn("Make www assets file from gui folder face this error:", err)
			break
		}
		repo.SetFile(&file)
		log.Info("Make www assets file from gui folder had been succeed!!")

	// *************** Ganjine Services ***************
	case 50:
		log.Info("Enter desire structure name in ```kebab-case``` like ```person-authentication```")
		var sName string
		fmt.Scanln(&sName)
		log.Info("Desire name: ", sName)

		var file = assets.File{
			Name: sName,
		}
		err = gg.MakeNewDatastoreFile(&file)
		if err != nil {
			log.Warn("Add new Ganjine file template face this error:", err)
			break
		}
		repo.Dependencies[FolderNameDataStore].SetFile(&file)
		log.Info("Add new structure had been succeed!!")
	case 51:
		log.Info("Enter desire full file name with extension!")
		var fileName string
		fmt.Scanln(&fileName)
		log.Info("Desire file name: ", fileName)

		var file = repo.GetFileRecursively(fileName)
		if file == nil {
			log.Warn("Desire file name not exist in repo!!")
			break
		}

		err = gg.UpdateDatastoreFile(file)
		if err != nil {
			log.Warn("Update Structure file face this error:", err)
		}
		log.Info("Update exiting Ganjine file had been succeed!!")

		// *************** GUI Services ***************
	case 70:
		log.Info("Enter desire full page name like 'store'")
		var pageName string
		fmt.Scanln(&pageName)
		log.Info("Desire page name: ", pageName)

		var jsFile = assets.File{Name: pageName}
		var htmlFile assets.File
		var cssFile assets.File
		var jsonFile assets.File

		err = wg.MakeNewPage(&jsFile, &htmlFile, &cssFile, &jsonFile)
		if err != nil {
			log.Warn("Add new GUI raw page face this error:", err)
		}
		repo.Dependencies[FolderNameGUI].Dependencies[FolderNameGUIPages].SetFile(&jsFile)
		repo.Dependencies[FolderNameGUI].Dependencies[FolderNameGUIPages].SetFile(&htmlFile)
		repo.Dependencies[FolderNameGUI].Dependencies[FolderNameGUIPages].SetFile(&cssFile)
		repo.Dependencies[FolderNameGUI].Dependencies[FolderNameGUIPages].SetFile(&jsonFile)
		log.Info("Add new GUI raw page had been succeed!!")

	// *************** Syllab Services ***************
	case 80:
		log.Info("Enter desire full file name with extension:")
		var cmd string
		fmt.Scanln(&cmd)
		log.Info("Desire file name: ", cmd)

		var file = repo.GetFileRecursively(cmd)
		if file == nil {
			log.Warn("Desire file name not exist in repo!!")
			break
		}

		var sgo = syllab.GenerationOptions{
			ForceUpdate: true,
		}
		log.Info("Use unsafe codes means don't copy data from given payload||buffer and just point to it for decoding fields! buffer can't GC until decoded struct free!")
		log.Info("Use unsafe codes Y|N :")
		fmt.Scanln(&cmd)
		log.Info("You choose: ", cmd)
		if cmd == "y" || cmd == "Y" {
			sgo.UnSafe = true
		}

		err = syllab.CompleteMethods(file, &sgo)
		if err != nil {
			log.Warn("Update Syllab encoder||decoder face this error:", err)
		}
		log.Info("Update Syllab encoder||decoder had been succeed!!")

		// *************** JSON Services ***************
	case 90:
		log.Info("Enter desire full file name with extension:")
		var cmd string
		fmt.Scanln(&cmd)
		log.Info("Desire file name: ", cmd)

		var file = repo.GetFileRecursively(cmd)
		if file == nil {
			log.Warn("Desire file name not exist in repo!!")
			break
		}

		var jgo = json.GenerationOptions{
			ForceUpdate: true,
		}

		log.Info("Just accept minifed encoded data to decode Y|N :")
		fmt.Scanln(&cmd)
		log.Info("You choose: ", cmd)
		if cmd == "y" || cmd == "Y" {
			jgo.Minifed = true
		}

		log.Info("Strict mode to decode data Y|N :")
		fmt.Scanln(&cmd)
		log.Info("You choose: ", cmd)
		if cmd == "y" || cmd == "Y" {
			jgo.Strict = true
		}

		log.Info("Use unsafe codes Y|N :")
		fmt.Scanln(&cmd)
		log.Info("You choose: ", cmd)
		if cmd == "y" || cmd == "Y" {
			jgo.UnSafe = true
		}

		err = json.CompleteMethods(file, &jgo)
		if err != nil {
			log.Warn("Update JSON encoder||decoder safe face this error:", err)
		}
		log.Info("Update JSON encoder||decoder safe had been succeed!!")

	default:
		log.Info("Nothing DO in given ID to prevent often mistakes enter bad ID!!!")
		goto Choose
	}

	goto Choose

Exit:
	log.Info("-------------------------------------------------------------------------------------")
	log.Info("CLI app run duration:", time.Since(start))
	log.Info("See you soon!")
}
