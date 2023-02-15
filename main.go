package main

import (
	"bytes"
	_ "embed"
	"github.com/pdfcpu/pdfcpu/pkg/api"
	"io"
	"log"
	"os"
)

//go:embed assets/RG-TEMPLATE.pdf
var pdf []byte

func main() {

	file1 := bytes.NewReader(pdf)
	file2 := bytes.NewReader(pdf)

	readers := []io.ReadSeeker{file1, file2}

	mergedFile := bytes.NewBuffer([]byte{})

	err := api.Merge(readers, mergedFile, nil)
	if err != nil {
		log.Panicln(err)
	}

	writer, err := os.Create("merged-file.pdf")
	if err != nil {
		log.Panicln(err)
	}
	defer writer.Close()

	_, err = writer.Write(mergedFile.Bytes())
	if err != nil {
		log.Panicln(err)
	}
}
