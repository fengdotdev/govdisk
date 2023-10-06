package main

import (
	"fmt"

	"github.com/fengdotdev/govdisk/v"
)



const(
	rootFolder = "/sandboxdata/"
)

func main() {


	fileName:="somefile.txt"

	path := "/some/path/to/"

	fullPath := path + fileName

	disk := v.NewDisk()

	fmt.Println(fullPath)

	err := disk.CreateFolder("some","/father/")
	if err != nil {
		fmt.Println(err.V())
	}


	// crud operations CREATE WRITE READ DELETE *UPDATE *APPEND

	// Logistics operations MOVE COPY RENAME DELETE EXIST LIST
		// --- AT FOLDER LEVEL  Create Move Copy Rename Delete Exist List 



	fmt.Println("Hello World")
}