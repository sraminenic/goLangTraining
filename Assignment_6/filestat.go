package main

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"sync"
)

func main() {

	//Folder Source
	folderSRC := "folder"

	extensionCh := make(chan string) //To communicate extensions
	//gorutine will Wait for all gouritine to complete.
	go func() {
		wg := &sync.WaitGroup{}
		wg.Add(1)
		readFolder(folderSRC, wg, extensionCh)
		wg.Wait()
		close(extensionCh)
	}()

	fileExtension := make(map[string]int)
	for fileExtensionMsg := range extensionCh {
		fileExtension[fileExtensionMsg]++
	}
	for extension, count := range fileExtension {
		fmt.Println(extension, count)
	}

}

/**
folderSRC - Folder source or folder name if it is in same location of application running
wg - Wait for all gouritine to complete.
extensionCh - unbuffered channel, to send if extension of file.
This method read all the folders, add extensions via channel communication
*/
func readFolder(folderSRC string, wg *sync.WaitGroup, extensionCh chan string) {
	defer wg.Done()
	files, err := ioutil.ReadDir(folderSRC)
	if err != nil {
		fmt.Println(err, " Unable to read")
	}
	for _, file := range files {
		//if directory then again run go routine
		if file.IsDir() {
			wg.Add(1)
			go readFolder(folderSRC+"/"+file.Name(), wg, extensionCh)
		} else {
			extension := filepath.Ext(file.Name())
			extensionCh <- extension
		}
	}
}
