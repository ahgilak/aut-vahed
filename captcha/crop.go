package captcha

import (
	"image"
	"image/jpeg"
	"io"
)

type SubImager interface {
	SubImage(r image.Rectangle) image.Image
}

func Part(c image.Image, i int, d int, w io.Writer) {
	rect := image.Rect(d*i, 0, d*(i+1), c.Bounds().Dy())

	jpeg.Encode(w, c.(SubImager).SubImage(rect), &jpeg.Options{Quality: 100})
}
