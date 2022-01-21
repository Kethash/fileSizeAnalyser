package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func DirSize(path string) (float64, error) {
	var size float64
	err := filepath.Walk(path, func(_ string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			size += float64(info.Size())
		}
		return err
	})
	return size, err
}

func main() {

	var path string = os.Args[1]
	size, _ := DirSize(path)
	fmt.Printf("The size of the folder is %.2f MB\n", size/1024.0/1024.0)

}
