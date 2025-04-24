package main

import (
	 "fmt"
	 "os"
	 "path/filepath"
)

func main() {
	fmt.Println("📁 Folder Organizer Starting...")
	fmt.Println("")

	f, err := os.Open("/home/nazma/Documents/go-projects/github.com/nazma98/basic-projects/folder-organizer/sample-testing")
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
			fmt.Println(v.Name(), ext)
		}
    }

	fmt.Println("")
	fmt.Println("🛑 Closing the directory 📁")
	defer f.Close()
}