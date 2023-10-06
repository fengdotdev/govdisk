package disk

import (
	"fmt"

	"github.com/fengdotdev/goerrorsplus/e"
)

type Disk struct {
	//path string folderid
	folders map[string]string

	//folderID / foldername
	folderIdAndName map[string]string

	//fileID / filename
	fileIdAndName map[string]string

	//folderID / path
	folderLocation map[string]string

	//fileID / FolderID
	fileLocation map[string]string
}


//TESTME
func NewDisk() *Disk {
	return &Disk{
		folders:         make(map[string]string),
		folderIdAndName: make(map[string]string),
		fileIdAndName:   make(map[string]string),
		folderLocation:  make(map[string]string),
		fileLocation:    make(map[string]string),
	}
}


//TESTME
func (d *Disk) CreateFolder(folderName string, path string) *e.ErrorPlus {

	if true{
		return e.NewErrorPlus(fmt.Errorf("folder already exists"), "folder already exists", []string{"disk", "folderOps"}, d.CreateFolder, folderName)
	}


	//create folder
	//add to folders
	//add to folderIdAndName
	//add to folderLocation
	//return folderID
	return  nil
}