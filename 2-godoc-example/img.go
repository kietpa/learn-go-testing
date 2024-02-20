package example

import "io"

type Image struct{}

// Decode images
func Decode(r io.Reader) (*Image, error) {
	return &Image{}, nil
}

// Crop images
func Crop(img *Image, x1, y1, x2, y2 int) error {
	return nil
}

// Encode images

func Encode(img *Image, format string, w io.Writer) error {
	return nil
}
