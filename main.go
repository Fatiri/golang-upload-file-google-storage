package main

import (
	"fmt"

	"github.com/golang-upload-gcs/helper"
)

func main() {
	path := "C:/Users/Ilham Fatiri/OneDrive/Pictures"
	filename := "filename"
	object, err := helper.UploadGCS(path, filename)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(object)
}
