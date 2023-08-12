package main

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
)

var dest = "/Users/rajanand/Documents/Downloaded"

func main() {

	_, errFileDoesNotExist := os.Stat(dest);
	if errors.Is(errFileDoesNotExist, os.ErrNotExist) {
		fmt.Println("Creating directory");
		os.MkdirAll(dest, os.ModePerm);
	}

	files, _ := filepath.Glob("/Users/rajanand/Downloads/*")

	listOfAllowedExtensions := map[string]bool{".pdf" : true, ".docx" : true}

	for i := 0; i < len(files); i++ {
		file := files[i]
		extension := filepath.Ext(file)
		if listOfAllowedExtensions[extension] {
			os.Rename(file, dest + "/" + filepath.Base(file))
		}
	}

}

