package main

import (
	"image/png"
	"log"
	"os"

	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/qr"
)

func main() {
	qrCode, err := qr.Encode("https://www.linkedin.com/in/akhil-c-1381a4246/", qr.L, qr.Auto)
	if err != nil {
		log.Fatal(err)
	}
	qrCode, err = barcode.Scale(qrCode, 256, 256)
	if err != nil {
		log.Fatal(err)
	}
	file, err := os.Create("qrcode.png")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	err = png.Encode(file, qrCode)
	if err != nil {
		log.Fatal(err)
	}
}
