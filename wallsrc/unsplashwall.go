package wallsrc

import (
	"bytes"
	"errors"
	"io"
	"net/http"
)

type UnsplashImagesWallpaper struct {
	body bytes.Buffer
}

func (img *UnsplashImagesWallpaper) Get() error {
	resp, err := http.Get("https://source.unsplash.com/1920x1080?wallpaper,nature,water,abstract")
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

func (img *UnsplashImagesWallpaper) WriteBody(w io.Writer) error {
	_, err := img.body.WriteTo(w)
	return err
}

func (img *UnsplashImagesWallpaper) String() string {
	return "Unsplash wallpaper,nature,water,abstract"
}

func (img *UnsplashImagesWallpaper) Clean() {
	img.body = bytes.Buffer{}
}
