package main

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
)

func main() {
	folderName := "folder"
	fileExtension := make(map[string]int)
	readFolder(folderName, fileExtension)
	fmt.Println(fileExtension)

}
//Read folder add file stats to map, This does not used go-routines and channels
func readFolder(folderName string, fileExtension map[string]int) {
	files, err := ioutil.ReadDir(folderName)
	if err != nil {
		fmt.Println(err, " Unable to read")
	}
	for _, file := range files {
		if file.IsDir() {
			readFolder(folderName+"/"+file.Name(), fileExtension)
		} else {
			extension := filepath.Ext(file.Name())
			if _, ok := fileExtension[extension]; ok {
				//do something here
				fileExtension[extension]++
			} else {
				fileExtension[extension] = 1
			}

		}
	}
}
