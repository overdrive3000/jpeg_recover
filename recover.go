package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatalln("Usage: ./recover image.raw")
	}

	// Open the block image file
	image, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatalln("The file can't be openned:", err)
	}
	defer image.Close()

	// JPEG digital signature
	ds1 := []byte{0xff, 0xd8, 0xff, 0xe0}
	ds2 := []byte{0xff, 0xd8, 0xff, 0xe1}

	// JPEG file count
	jc := 0

	// Flag for opened files
	open := false
	var jpgfile *os.File

	// Needed arrays reading file in 512Byte chunks
	buffer := make([]uint8, 512)

	// Read first block
	binary.Read(image, binary.LittleEndian, &buffer)

	// Read image until EOF
	for {
		err := binary.Read(image, binary.LittleEndian, &buffer)
		if err == io.EOF {
			break
		}
		if bytes.Equal(ds1, buffer[:4]) || bytes.Equal(ds2, buffer[:4]) {
			// Set found jpeg name
			fname := fmt.Sprintf("%03d.jpg", jc)

			if !open {
				jpgfile, err = os.Create(fname)
				if err != nil {
					log.Panicln("File can't be created:", err)
				}
				binary.Write(jpgfile, binary.LittleEndian, buffer)
				open = true
			}
			if open {
				jpgfile.Close()
				jpgfile, err = os.Create(fname)
				binary.Write(jpgfile, binary.LittleEndian, buffer)
				jc++
			}
		} else {
			if open {
				binary.Write(jpgfile, binary.LittleEndian, buffer)
			}
		}
	}

	defer jpgfile.Close()

	fmt.Println("Images where successfully restored")
	os.Exit(0)

}
