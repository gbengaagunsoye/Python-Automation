package main

import (
	"fmt"
	"os"
	"path/filepath"
)

// var i int = 42

// var (
// 	firstName string = "Gbenga"
// 	lastName  string = "Agunsoye"
// 	age       int    = 29
// 	season    int    = 11
// )

// var (
// 	counter int = 0
// )

// func main() {
// 	// // var i int
// 	// // i = 42
// 	// var j int = 27
// 	// // k := 99
// 	// fmt.Printf("%v, %T", i, i)

// 	s := [5]int{3, 4, 5, 7, 8}
// 	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
// 		s[i], s[j] = s[j], s[i]
// 	}

// 	fmt.Println(s)
// 	// fmt.Println(reversedNumbers)

// }

func main() {
	// file, _ := os.OpenFile("file.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	// file.Write([]byte("append data\n"))

	// file.Close()

	// dpath := "./dir"
	// fname := "file.txt"
	// _ = os.MkdirAll(dpath, 0777)
	// fpath := filepath.Join(dpath, fname)
	// file, _ := os.Create(fpath)
	// file.Close()

	// _ = os.Remove(fpath)
	// _ = os.RemoveAll(dpath)

	//GET FILE INFORMATION

	// file, _ := os.Stat("Main.go")
	// f1 := fmt.Println
	// f1("File Name: ", file.Name())
	// f1("File Size: ", file.Size())
	// f1("Last modified: ", file.ModTime())
	// f1("Is a Directory: ", file.IsDir())

	// ff := fmt.Printf
	// ff("Permission 9-bit: %s\n", file.Mode())
	// ff("Permission 3-digit: %o\n", file.Mode())
	// ff("Permission 4-digit: %04o\n", file.Mode())

	// files, _ := ioutil.ReadDir(".")
	// for _, files := range files {
	// 	fmt.Println(files.Name(), files.ModTime())
	// }

	// device := map[string]int{

	// 	"Router1": 90,
	// }

	// fmt.Println(device)

	scan := func(
		path string, i os.FileInfo, _ error) error {
		fmt.Println(i.IsDir(), path)
		return nil
	}
	_ = filepath.Walk(".", scan)
}
