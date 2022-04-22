package cpt

import (
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
	return fmt.Sprintf(`{"path": "%s", "name": "%s", "key": "%s", "uuid": "%s"}`, file.Path, file.Name, string(file.Key), file.UUID)
}
