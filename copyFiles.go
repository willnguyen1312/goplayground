package main

import (
	"io"
	"log"
	"os"
	"path/filepath"
	"regexp"
)

// File struct
type File struct {
	name string
	path string
}

func main() {
	// oldLocation := "./temp/test.txt"
	// newLocation := "./test.txt"
	// Copy(oldLocation, newLocation)
	var files []File
	libRegEx, e := regexp.Compile("(?i).(jpg|jpeg|heic|mp4|mov)$")
	if e != nil {
		log.Fatal(e)
	}
	err := filepath.Walk("./001",
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if libRegEx.MatchString(info.Name()) {
				// println(path)
				files = append(files, File{name: info.Name(), path: path})
			}
			return nil
		})
	if err != nil {
		log.Println(err)
	}

	newLocationPrefix := "./copy/"
	for _, file := range files {
		Copy(file.path, newLocationPrefix+file.name)
	}

}

// Copy the src file to dst. Any existing file will be overwritten and will not
// copy file attributes.
func Copy(src, dst string) error {
	in, err := os.Open(src)
	if err != nil {
		return err
	}
	defer in.Close()

	out, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, in)
	if err != nil {
		return err
	}
	return out.Close()
}
