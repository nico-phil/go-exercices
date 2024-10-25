package main

import (
	"archive/zip"
	"bytes"
	"fmt"
)

func SentinalErrorExample() {

	data := []byte("this is not a zip file")
	notZipfile := bytes.NewReader(data)

	_, err := zip.NewReader(notZipfile, int64(len(data)))
	if err != nil {
		if err == zip.ErrFormat {
			fmt.Println("told you")
		}
	}

}

