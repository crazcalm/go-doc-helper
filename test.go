/*
From what I can tell, Go packages that have docs either have a
doc.go file, package.go file, or a .go file that has the same
name as the directory name.
*/
package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

const (
	goDoc string = "/home/marcus/.gvm/gos/go1.7/src/cmd/go/doc.go"
)

func walk(root string, paths []string) []string {
	entries, err := ioutil.ReadDir(root)
	if err != nil {
		panic(err)
	}
	for _, item := range entries {
		if item.IsDir() {
			fullPath := filepath.Join(root, item.Name())
			fmt.Println(fullPath)
			paths = append(paths, fullPath)
			paths = walk(fullPath, paths)
		}
	}
	return paths
}

func isPackage(path string) bool {
	var result bool

	dirName := filepath.Dir(path) + ".go"
	fmt.Println(dirName)
	//Check for doc.go, package.go dirName.go files
	files, err := ioutil.ReadDir(path)
	if err != nil {
		panic(err)
	}
	fmt.Println(files)
	for _, file := range files {
		if file.Name() == "doc.go" {
			result = true
		} else if file.Name() == "package.go" {
			result = true
		} else if file.Name() == dirName {
			result = true
		}
	}

	return result
}

func main() {
	sl := filepath.Join(runtime.GOROOT(), "src")
	pk := filepath.Join(os.Getenv("GOPATH"), "src")
	fmt.Println(sl)
	fmt.Println(pk)

	var list []string
	paths := walk(sl, list)
	paths = walk(pk, paths)
	fmt.Println(paths)

	var newList []string
	for _, path := range paths {
		if isPackage(path) {
			newList = append(newList, path)
		}
	}

	fmt.Println("Final list?")
	for _, path := range newList {
		fmt.Println(path)
	}

	// I still need to remove the unneeded part of the path
	var lastList []string
	for _, path := range newList {
		if strings.Contains(path, sl) {
			lastList = append(lastList, path[len(sl)+1:])
		} else if strings.Contains(path, pk) {
			lastList = append(lastList, path[len(pk)+1:])
		}
	}

	for _, path := range lastList {
		fmt.Println(path)
	}
}
