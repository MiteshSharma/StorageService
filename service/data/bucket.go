package data

import (
	"time"
	"fmt"
	"strconv"
)

type Bucket struct {
	Name string
	Size int64
	LastModifiedAt time.Time
}

func NewBucket(name string, size int64, modificationAt time.Time) *Bucket {
	return &Bucket{
		Name: name,
		Size: size,
		LastModifiedAt: modificationAt}
}

func (c Bucket) Print() {
	fmt.Println("Bucket name: "+ c.Name +" Size: "+ strconv.FormatInt(c.Size, 10)+" Modification date ", c.LastModifiedAt)
}
