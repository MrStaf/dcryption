package main

import (
	"strings"

	"github.com/MrStaf/dcryption/cpt"
)

func main() {
	// // Generate a new key
	// key := string(cpt.Generate())
	// // Encrypt the file
	// cpt.Encrypt("./test.txt", key, "./test.enc")
	// cpt.Decrypt("./test.enc", key, "./test.dec")
	output_encrypted := "./encrypted"
	// output_decrypted := "./decrypted"
	filepath := "./test"
	file_info := "data.json"

	// List the files in the directory
	files := cpt.List(filepath)
	// Create the output file info
	output_file_infos := strings.Join([]string{output_encrypted, file_info}, "/")
	cpt.MkAll(output_file_infos)
	cpt.SaveFilesInfos(files, output_file_infos)

	// // Print the files
	// for _, file := range files {
	// 	// fmt.Println(file.Path, file.Name, file.Key, file.UUID)
	// 	// Create output file
	// 	output := strings.Join([]string{output_encrypted, file.UUID + ".enc"}, "/")

	// 	// Encrypt the file
	// 	cpt.Encrypt(file.Path, file.Key, output)
	// 	output_d := strings.Replace(file.Path, filepath, output_decrypted, 1)
	// 	fmt.Println(output, output_d)
	// 	// Decrypt the file
	// 	cpt.Decrypt(output, file.Key, output_d)
	// }
}
