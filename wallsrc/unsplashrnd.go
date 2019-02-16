package wallsrc

import (
	"bytes"
	"errors"
	"io"
	"net/http"
)

type UnsplashImagesRnd struct {
	body bytes.Buffer
}

func (img *UnsplashImagesRnd) Get() error {
	resp, err := http.Get("https://source.unsplash.com/random/1920x1080")
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return errors.New(img.String() + ": not downloaded")
	}

	_, err = img.body.ReadFrom(resp.Body)
	return err
}

func (img *UnsplashImagesRnd) WriteBody(w io.Writer) error {
	_, err := img.body.WriteTo(w)
	return err
}

func (img *UnsplashImagesRnd) String() string {
	return "Unsplash random"
}

func (img *UnsplashImagesRnd) Clean() {
	img.body = bytes.Buffer{}
}
