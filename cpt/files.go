package cpt

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
)

// Create a file class
type File struct {
	// The file path
	Path string
	// The file name
	Name string
	// The file key
	Key []byte
	// The file uuid
	UUID string
	// Is a directory
	IsDir bool
}

func New(Path string, Name string, Key []byte, UUID string) File {
	return File{Path: Path, Name: Name, Key: Key, UUID: UUID}
}

func (file File) ToJSON() string {
	// Convert key to Base64
	key_base64 := base64.StdEncoding.EncodeToString(file.Key)
	return fmt.Sprintf(`{"path": "%s", "name": "%s", "key": "%s", "uuid": "%s"}`, file.Path, file.Name, key_base64, file.UUID)
}

func ParseJson(json_string string) []File {
	var list_files []File
	fmt.Println(json_string)
	err := json.Unmarshal([]byte(json_string), &list_files)
	if err != nil {
		fmt.Println(err)
	}
	return list_files
}
