package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func ReadFile() {

}

func ReadDir(dirname string) []os.FileInfo {
	files, err := ioutil.ReadDir(dirname)
	if err != nil {
		log.Fatal(err)
	}
	return files
}

func findImportantFiles(fileList []os.FileInfo, fileName string) int {
	for i, f := range fileList {
		if f.Name() == fileName {
			return i
		}
	}
	return -1
}

func removeFiles(fileList []os.FileInfo) []os.FileInfo {
	i := findImportantFiles(fileList, "main.exe")
	if i == -1 {
		log.Fatal("Error, file not found")
	}
	fileList = remove(fileList, i)

	i = findImportantFiles(fileList, "main.go")
	if i == -1 {
		log.Fatal("Error, file not found")
	}
	fileList = remove(fileList, i)

	return fileList
}

func remove(fileList []os.FileInfo, index int) []os.FileInfo {
	fileList[index] = fileList[len(fileList)-1]
	return fileList[:len(fileList)-1]
}

func main() {
	path, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	// fmt.Println(path)

	files := ReadDir(path) //Gets list of files in directory

	files = removeFiles(files) //removes our important files

	for _, f := range files {
		fmt.Println(f.Name())
	}

	//Now crypto...
}
