package main

import (
	"flag"
	"fmt"
	"image/color"
	"io"
	"net/http"
	"net/url"
	"os"

	"github.com/yeqown/go-qrcode/writer/standard"

	"github.com/yeqown/go-qrcode/v2"
)

const resultDir = "./qrCodes"

func main() {
	link := flag.String("linkedin", "https://korpitech.com", "Linkedin profile link")
	//logo := flag.String("logo", "https://github.com/PPerminov/qrLogo/raw/main/linkeindLogo.png", "logo to use")
	name := flag.String("name", "linkedinQRCode", "name for the file")
	flag.Parse()

	qrc, err := qrcode.NewWith(*link,
		qrcode.WithEncodingMode(qrcode.EncModeByte),
		qrcode.WithErrorCorrectionLevel(qrcode.ErrorCorrectionHighest),
	)
	if err != nil {
		panic(err)
	}
	err = os.MkdirAll(resultDir, os.ModeDir)
	if err != nil {
		panic(err)
	}
	w, err := standard.New(
		fmt.Sprintf("%s/%s.png", resultDir, *name),
		standard.WithQRWidth(120),
		standard.WithBorderWidth(120),
		standard.WithBorderWidth(90),
		//standard.WithLogoImageFilePNG("/home/pavel/Pictures/Screenshots/logo.png"),
		//standard.WithCircleShape(),
		//standard.WithHalftone(downloadAndSaveOrReturnPath(*logo)),
		//standard.WithHalftone("/home/pavel/work/korpitech/site/assets/logo.png"),
		//standard.WithBgColor(color.RGBA{G: 0, B: 0, R: 0}),
		standard.WithFgColor(color.RGBA{R: 25, G: 78, B: 67}),
		// standard.WithFgColor(color.RGBA{G: 226, B: 22, R: 0, A: 0}),
		//standard.WithFgColor(color.RGBA{G: 129, B: 129, R: 129}),
	)
	if err != nil {
		panic(err)
	}

	if err = qrc.Save(w); err != nil {
		panic(err)
	}
}

func downloadAndSaveOrReturnPath(path string) string {
	u, err := url.Parse(path)
	if err != nil || u.Host == "" {
		return path
	}
	data, err := http.Get(u.String())
	if err != nil {
		panic(err)
	}
	defer data.Body.Close()
	file, err := os.CreateTemp("", "logoImage*.png")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	_, err = io.Copy(file, data.Body)
	if err != nil {
		panic(err)
	}
	return file.Name()
}
