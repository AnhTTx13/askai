package data

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

// File data will be use with struct of type T
type File[T any] struct {
	FileDir  string
	FileName string
}

func NewFile[T any](file_dir, file_name string) *File[T] {
	// Make directory if not exist
	if err := os.MkdirAll(file_dir, os.ModeDir|0755); err != nil {
		fmt.Println(err.Error())
	}
	return &File[T]{
		FileDir:  file_dir,
		FileName: file_name,
	}
}

// Save data from source `data` to file (json format)
func (f *File[T]) Save(data T) error {
	path := fmt.Sprintf("%s/%s", f.FileDir, f.FileName)

	fileData, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(path, fileData, 0644)
}

// Load data from file to `obj`
func (f *File[T]) Load(obj *T) error {
	path := fmt.Sprintf("%s/%s", f.FileDir, f.FileName)

	file, err := os.OpenFile(path, os.O_CREATE|os.O_RDONLY, 0644)
	defer file.Close()
	if err != nil {
		return err
	}

	data := []byte{}
	buf := make([]byte, 1024)

	for {
		n, err := file.Read(buf)
		if err != nil && err != io.EOF {
			return err
		}
		if n == 0 {
			break
		}
		data = append(data, buf[:n]...)
	}

	return json.Unmarshal(data, obj)
}
