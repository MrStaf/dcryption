package main

import (
	"fmt"

	"github.com/MrStaf/dcryption/cpt"
	"github.com/MrStaf/dcryption/win"
)

func main() {
	// // Generate a new key
	key := string(cpt.Generate())
	fmt.Println(key)
	// Save the key to a file
	// cpt.Save([]byte(key), "key.txt")
	// // Encrypt the file
	// cpt.Encrypt("./test.txt", key, "./test.enc")
	// cpt.Decrypt("./test.enc", key, "./test.dec")
	// output_encrypted := "./encrypted"
	// // output_decrypted := "./decrypted"
	// filepath := "./test"
	// file_info := "data.json"
	// cpt.EncryptFolder(filepath, output_encrypted, file_info)
	// output_d := strings.Replace(file.Path, filepath, output_decrypted, 1)
	// cpt.Decrypt(output, file.Key, output_d)
	w := win.Init()
	win.Run(&w)
}
