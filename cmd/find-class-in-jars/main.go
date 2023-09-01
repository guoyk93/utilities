package main

import (
	"archive/zip"
	"errors"
	"github.com/guoyk93/rg"
	"github.com/guoyk93/utilities"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func findClassInJar(jarPath string, className string) (ok bool, err error) {
	defer rg.Guard(&err)

	className = strings.ReplaceAll(className, ".", "/")

	zf := rg.Must(zip.OpenReader(jarPath))
	defer zf.Close()

	for _, f := range zf.File {
		if strings.Contains(f.Name, className) {
			ok = true
			return
		}
	}

	return
}

func main() {
	var err error
	defer utilities.Exit(&err)
	defer rg.Guard(&err)

	classNames := os.Args[1:]

	if len(classNames) == 0 {
		err = errors.New("no class name specified")
		return
	}

	err = filepath.WalkDir(".", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() {
			return nil
		}
		if strings.ToLower(filepath.Ext(path)) != ".jar" {
			return nil
		}
		for _, className := range classNames {
			found, err := findClassInJar(path, className)
			if err != nil {
				return err
			}
			if found {
				log.Println(">>> " + path)
				log.Println("<<< " + className)
			}
		}
		return nil
	})
}
