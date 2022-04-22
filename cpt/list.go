package cpt

import (
	"crypto/rand"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func randomString(length int) string {
	// Create uuid
	b := make([]byte, length)
	// Populate the byte slice with cryptographically secure random bytes
	_, err := rand.Read(b)
	if err != nil {
		fmt.Println(err)
	}
	// Return the byte slice as a string
	return fmt.Sprintf("%x", b)
}

func List(path string) []File {
	// list the files in the directory
	files, err := ioutil.ReadDir(path)
	if err != nil {
		panic(err)
	}
	// create a slice to hold the file names
	var Files []File
	// loop through the files
	for _, file := range files {
		// generate random name
		name := randomString(32)
		// join path by os.sep
		path := strings.Join([]string{path, file.Name()}, "/")

		// add the file path to the slice and the random name associate with it
		new_file := New(path, file.Name(), Generate(), name)
		// if the file is a directory
		if file.IsDir() {
			// recursively call the function
			Files = append(Files, List(new_file.Path)...)
		} else {
			Files = append(Files, new_file)
		}
	}
	// return the slice of file names
	return Files
}

func MkAll(filepath string) {
	// create each folders in the path that didn't exist
	path := strings.Split(filepath, "/")
	path_string := strings.Join(path[:len(path)-1], "/")
	if err := os.MkdirAll(path_string, os.ModePerm); err != nil {
		panic(err)
	}
}

func SaveFilesInfos(files []File, output string) {
	json_string := "["
	// loop through the files
	for _, file := range files {
		// convert File to json string
		json_string += file.ToJSON() + ","
		// append the json string to the output file
	}
	json_string += "]"
	if err := ioutil.WriteFile(output, []byte(json_string), 0644); err != nil {
		panic(err)
	}
}
