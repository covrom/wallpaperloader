package wallsrc

import (
	"bytes"
	"errors"
	"io"
	"net/http"
)

type UnsplashImages struct {
	body bytes.Buffer
}

func (img *UnsplashImages) Get() error {
	resp, err := http.Get("https://source.unsplash.com/random/1920x1080")
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return errors.New("Unsplash: not downloaded")
	}

	_, err = img.body.ReadFrom(resp.Body)
	return err
}

func (img *UnsplashImages) WriteBody(w io.Writer) error {
	_, err := img.body.WriteTo(w)
	return err
}

func (img *UnsplashImages) String() string {
	return "Unsplash random"
}

func (img *UnsplashImages) Clean() {
	img.body = bytes.Buffer{}
}
