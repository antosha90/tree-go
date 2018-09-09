package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func recursiveDirTreePrint(path string, treePrefix string, showFiles bool, out io.Writer) error {
	files, err := ioutil.ReadDir(path)
	endOfDir := false

	if !showFiles {
		var dirFiles []os.FileInfo
		for _, file := range files {
			if file.IsDir() {
				dirFiles = append(dirFiles, file)
			}
		}
		files = dirFiles
	}

	for index, file := range files {
		filePrefix := "├───"
		endOfDir = index+1 == len(files)
		if endOfDir {
			filePrefix = "└───"
		}

		treeString := treePrefix + filePrefix + file.Name()
		if !file.IsDir() {
			sizePostfix := ""
			switch file.Size() {
			case 0:
				sizePostfix = " (empty)"
			default:
				sizePostfix = fmt.Sprintf(" (%db)", file.Size())
			}
			treeString += sizePostfix
		}

		// fmt.Println(treeString)
		fmt.Fprintf(out, treeString+"\n")

		if file.IsDir() {
			nextTreePrefix := treePrefix + "│\t"
			if endOfDir {
				nextTreePrefix = treePrefix + "\t"
			}
			recursiveDirTreePrint(path+"/"+file.Name(), nextTreePrefix, showFiles, out)
		}
	}
	return err
}

func dirTree(out io.Writer, path string, printFiles bool) error {
	err := recursiveDirTreePrint(path, "", printFiles, out)
	return err
}

func main() {
	out := os.Stdout
	if !(len(os.Args) == 2 || len(os.Args) == 3) {
		panic("usage go run main.go . [-f]")
	}
	path := os.Args[1]
	printFiles := len(os.Args) == 3 && os.Args[2] == "-f"
	err := dirTree(out, path, printFiles)
	if err != nil {
		panic(err.Error())
	}
}
