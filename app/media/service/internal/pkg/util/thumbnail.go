package util

import (
	"bytes"
	"errors"
	"fmt"
	"image/png"
	"io"

	"github.com/disintegration/imaging"
)

func Thumbnail(imgObject io.Reader) (*bytes.Buffer, error) {
	buf := new(bytes.Buffer)
	img, err := imaging.Decode(imgObject)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Thumbnail image error: %s", err))
	}

	thumbnail := imaging.Thumbnail(img, 240, 240, imaging.CatmullRom)

	if err := png.Encode(buf, thumbnail); err != nil {
		return nil, errors.New(fmt.Sprintf("Thumbnail image error: %s", err))
	}
	return buf, nil
}
