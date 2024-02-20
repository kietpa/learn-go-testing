package example_test

import (
	"fmt"
	_ "image/png"
	"io"
	example "test/2-godoc-example"
	"testing"
)

func ExampleImage_crop() {
	var f io.Reader

	img, err := example.Decode(f)
	if err != nil {
		panic(err)
	}

	err = example.Crop(img, 0, 0, 20, 20)
	if err != nil {
		panic(err)
	}
	var out io.Writer
	example.Encode(img, ".png", out)

	fmt.Println("See out.png for image")
	// Output:
	// See out.png for image
}

func TestDecode(t *testing.T) {

}

func TestCrop(t *testing.T) {

}
