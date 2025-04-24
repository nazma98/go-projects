package main

import (
	 "fmt"
	 "os"
)

func main() {
	fmt.Println("📁 Folder Organizer Starting...")

	f, err := os.Open("/home/nazma/Documents/go-projects/github.com/nazma98/basic-projects/folder-organizer")
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
        fmt.Println(v.Name(), v.IsDir())
    }

	fmt.Println("Closing the directory 📁")
	defer f.Close()
}