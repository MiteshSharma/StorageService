package data

import (
	"time"
	"fmt"
	"strconv"
)

type File struct {
	Name string
	Size int64
	LastModifiedAt time.Time
}

func NewFile(name string, size int64, modificationAt time.Time) *File {
	return &File{
		Name: name,
		Size: size,
		LastModifiedAt: modificationAt}
}

func (f File) Print() {
	fmt.Println("File name: "+ f.Name +" Size: "+ strconv.FormatInt(f.Size, 10)+" Modification date ", f.LastModifiedAt)
}