package main

import (
	"bytes"
	"fmt"
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"io/ioutil"
	"os"

	"github.com/mattn/go-sixel"
)

func main() {
	fmt.Println("Hello world")
}

func ShowImage(filename string) error {
	var buf, _ = ioutil.ReadFile(filename)
	var img, _, _ = image.Decode(bytes.NewReader(buf))
	return sixel.NewEncoder(os.Stdout).Encode(img)
}
