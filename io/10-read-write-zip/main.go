package main

import (
	"archive/zip"
	"bytes"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	var buf bytes.Buffer

	zipW := zip.NewWriter(&buf)
	f, err := zipW.Create("file.txt")

	if err != nil {
		panic(err)
	}

	_, err = f.Write([]byte("this is my file content"))

	if err != nil {
		panic(err)
	}

	err = zipW.Close()

	if err != nil {
		panic(err)
	}

	ioutil.WriteFile("text.zip", buf.Bytes(), os.ModePerm)

	zipR, err := zip.OpenReader("text.zip")

	for _, file := range zipR.File {
		r, _ := file.Open()
		io.Copy(os.Stdout, r)
	}
}
