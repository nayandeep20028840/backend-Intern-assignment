package utils

import (
	"image"
	"net/http"
	_ "image/jpeg"
	_ "image/png"
)

func DownloadImage(url string) (image.Image, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	img, _, err := image.Decode(resp.Body)
	return img, err
}

func CalculatePerimeter(img image.Image) int {
	bounds := img.Bounds()
	width := bounds.Max.X - bounds.Min.X
	height := bounds.Max.Y - bounds.Min.Y
	return 2 * (width + height)
}
