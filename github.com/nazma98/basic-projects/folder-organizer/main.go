package main

import (
	 "fmt"
	 "os"
	 "path/filepath"
)

func main() {
	// defaultFolders := []string{"Music", "Videos", "Docs", "Images", "Others"}

	fmt.Println("📁 Folder Organizer Starting...")
	fmt.Println("")

	basePath := "/home/nazma/Documents/go-projects/github.com/nazma98/basic-projects/folder-organizer/sample-testing"

	f, err := os.Open(basePath)
    if err != nil {
        fmt.Println(err)
        return
    }
    files, err := f.ReadDir(0)
    if err != nil {
        fmt.Println(err)
        return
    }

    for _, v := range files {
		if !v.IsDir() {
			ext := filepath.Ext(v.Name())
			oldPath := filepath.Join(basePath, v.Name())
			fmt.Println(v.Name(), ext)
			if ext == ".pdf" || ext == ".doc" {
				folderName := "Documents"
				folderPath := filepath.Join(basePath, folderName)
				newPath := filepath.Join(folderPath, v.Name())

				if info, err := os.Stat(folderPath); err == nil && info.IsDir() {
						err := os.Rename(oldPath, newPath)
						if err != nil {
							fmt.Println("Error moving file:", err)
							continue
						}
						fmt.Printf("📄 %s moving to %s", v.Name(), folderName)
					} else if os.IsNotExist(err) {
					err := os.Mkdir(folderPath, 0755)
					if err != nil {
						fmt.Println("Error creating folder:", err)
						return
					}
				}
				} 
			}

		}

	fmt.Println("")
	fmt.Println("🛑 Closing the directory 📁")
	defer f.Close()
}