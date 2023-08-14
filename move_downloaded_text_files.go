package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

var (
	errInvalidArg = errors.New("invalid arg")
	dest          = ""
	src           = ""
)

func main() {
	// check for dest
	checkDir(dest)

	files, _ := filepath.Glob(src)

	listOfAllowedExtensions := map[string]map[string]bool{
		"documents": {
			".pdf":  true,
			".docx": true,
		},
		"images": {
			".jpg": true,
			".png": true,
		},
		"movies": {
			".mp4": true,
		},
	}

	for i := 0; i < len(files); i++ {
		file := files[i]
		extension := strings.ToLower(filepath.Ext(file))
		for key, val := range listOfAllowedExtensions {
			if _, ok := val[extension]; ok {
				checkDir(fmt.Sprintf("%s/%s", dest, key))
				// Now copy all the files
				err := os.Rename(fmt.Sprintf("%s/%s", src, file), fmt.Sprintf("%s/%s/%s", dest, key, file))
				if err != nil {
					log.Println("Error moving: ", err)
				}
			}
		}
	}

}

/*
getArgs gets all system arguments

Format:
  - dest=*
  - src=*
*/
func getArgs() map[string]string {
	args := os.Args
	if len(args) == 0 {
		panic(fmt.Errorf("args not supplied"))
	}
	for _, arg := range args {
		if strings.Contains(arg, "src") {
			src = validateArg(arg)
		}
		if strings.Contains(arg, "dest") {
			dest = validateArg(arg)
		}
	}
}

func validateArg(arg string) string {
	d := strings.Split(arg, "=")
	if len(d) != 2 {
		panic(errInvalidArg)
	}
	if !strings.Contains(arg, "/") {
		panic(errInvalidArg)
	}
	return d[1]
}

func checkDir(dir string) {
	if _, err := os.Stat(dir); err != nil {
		return
	} else {
		if errors.Is(err, os.ErrNotExist) {
			os.MkdirAll(dir, 0644)
		}
	}
}
