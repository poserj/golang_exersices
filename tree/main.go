package main

import (
	"fmt"
	"io/fs"
	"io"
	"log"
	"os"
	"path/filepath"
	"sort"
)

func sortFilesInPlace(files []fs.DirEntry) []fs.DirEntry {
	sort.Slice(files, func(i, j int) bool {
		return files[i].Name() < files[j].Name()
	})
	return files
}

func filePrintAttr(file fs.DirEntry, out io.Writer) {
	f, _ := file.Info()
	fmt.Fprint(out, " (")
	if f.Size() == 0{
	fmt.Fprint(out, f.Size())
	fmt.Fprintln(out, "b)")
	}else{
	fmt.Fprintln(out, "empty)")
	}
	
}

func getTab(s string, first bool) string {
	if first {
		return ""
	} else {
		return s + "\t"
	}
}

func getCurFiles(dirname string, first bool, tab_before string, printFiles bool, out io.Writer) {
	tab := getTab(tab_before, first)
	files, err := os.ReadDir(dirname)
	if err != nil {
		log.Fatal("Cant read folder")
	}
	sortFilesInPlace(files)
	for _, file := range files {
		fmt.Fprint(out, tab+"├───"+file.Name())
		if !file.IsDir() {
			filePrintAttr(file, out)

		} else {
			fmt.Fprintln(out,"")
			rec_folder := filepath.Join(dirname, file.Name())
			if printFiles {
				getCurFiles(rec_folder, false, tab, true, out)
			}
		}

	}
}

func dirTree(out io.Writer, path string, printFiles bool) error {

	getCurFiles(path, printFiles, "", printFiles, out)
	return nil

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
