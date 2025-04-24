package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func folderExists(folderName string, basePath string) bool {
	folderPath := filepath.Join(basePath, folderName)
	info, err := os.Stat(folderPath)
	return err == nil && info.IsDir()
}

func createFolder(folderName string, basePath string) {
	if !folderExists(folderName, basePath) {
		folderPath := filepath.Join(basePath, folderName)
		err := os.Mkdir(folderPath, 0755)
		if err != nil {
			fmt.Println("❌ Error creating folder:", err)
		}
	}
}

func moveFiles(folderName string, basePath string, fileName string) {
	oldPath := filepath.Join(basePath, fileName)
	folderPath := filepath.Join(basePath, folderName)
	newPath := filepath.Join(folderPath, fileName)

	createFolder(folderName, basePath)

	if info, err := os.Stat(folderPath); err == nil && info.IsDir() {
		err := os.Rename(oldPath, newPath)
		if err != nil {
			fmt.Println("Error moving file:", err)
			return
		}
		fmt.Printf("📄 %s moving to %s\n", fileName, folderName)
	}
}

func main() {
	fmt.Println("📁 Folder Organizer Starting...")
	fmt.Println("")

	basePath := "/home/nazma/Documents/go-projects/github.com/nazma98/basic-projects/folder-organizer/sample-testing"

	f, err := os.Open(basePath)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("📁 Directory reading started...")
	files, err := f.ReadDir(0)
	if err != nil {
		fmt.Println(err)
		return
	}

	fileCount := 0

	for _, v := range files {
		if !v.IsDir() {
			fileName := v.Name()
			ext := filepath.Ext(fileName)
			if ext == ".pdf" || ext == ".doc" || ext == ".txt" {
				folderName := "Documents"
				moveFiles(folderName, basePath, fileName)
			} else if ext == ".jpg" || ext == "jpeg" || ext == "png" {
				folderName := "Images"
				moveFiles(folderName, basePath, fileName)
			} else if ext == ".mp3" {
				folderName := "Music"
				moveFiles(folderName, basePath, fileName)
			} else if ext == ".mp4" || ext == ".mov" || ext == ".avi" || ext == ".amv" {
				folderName := "Videos"
				moveFiles(folderName, basePath, fileName)
			} else {
				folderName := "Others"
				moveFiles(folderName, basePath, fileName)
			}
		} else {
			fileCount = fileCount + 1
		}
	}
	if fileCount > 0 {
		fmt.Println("❌ No files found!")
	}

	fmt.Println("")
	fmt.Println("🛑 Closing the directory 📁")
	defer f.Close()
}
