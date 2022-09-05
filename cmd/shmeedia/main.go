package main

import (
	"bytes"
	"errors"
	"fmt"
	"image/jpeg"
	"image/png"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/nfnt/resize"
)

func main() {
	files, err := ioutil.ReadDir("input/")
	if err != nil {
		panic(err)
	}

	for _, file := range files {
		if file.IsDir() {
			continue
		}
		path := filepath.Join("input/", file.Name())
		dest, err := CopyImage(path)
		if err != nil {
			panic(err)
		}
		log.Printf("copied %s to %s", path, dest)
	}
}

func imageType(sourceFile string) string {
	extension := strings.ToLower(filepath.Ext(sourceFile))
	if extension == ".jpeg" || extension == ".jpg" {
		return "jpg"
	}
	return "png"
}

func CopyImage(sourceFile string) (string, error) {
	datestamp := time.Now().Format("2006-01-02")
	destFile := fmt.Sprintf("output/%s-%s", datestamp, filepath.Base(sourceFile))

	// read in file
	data, err := ioutil.ReadFile(sourceFile)
	if err != nil {
		return destFile, err
	}

	// open new file
	out, err := os.Create(destFile)
	if err != nil {
		return destFile, err
	}
	defer out.Close()

	imageType := imageType(sourceFile)
	switch imageType {
	case "jpg":
		image, err := jpeg.Decode(bytes.NewReader(data))
		if err != nil {
			return destFile, err
		}
		croppedImage := resize.Thumbnail(800, 800, image, resize.NearestNeighbor)
		jpeg.Encode(out, croppedImage, nil)
	case "png":
		image, err := png.Decode(bytes.NewReader(data))
		croppedImage := resize.Thumbnail(800, 800, image, resize.NearestNeighbor)
		if err != nil {
			return destFile, err
		}
		png.Encode(out, croppedImage)
	default:
		err = errors.New(fmt.Sprintf("don't know how to encode %s", sourceFile))
		return destFile, err
	}

	return destFile, nil
}
