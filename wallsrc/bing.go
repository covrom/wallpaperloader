package wallsrc

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

type BingImages struct {
	body bytes.Buffer
}

func (img *BingImages) Get() error {
	resp, err := http.Get("https://www.bing.com/HPImageArchive.aspx?format=js&idx=0&n=1&mkt=ru_RU")
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	var j struct {
		Images []struct {
			URL string `json:"url"`
		} `json:"images"`
	}

	dec := json.NewDecoder(resp.Body)
	err = dec.Decode(&j)
	if err != nil {
		return err
	}
	if len(j.Images) > 0 {
		resp, err := http.Get("https://www.bing.com" + j.Images[0].URL)
		if err != nil {
			return err
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			return errors.New("Bing: not downloaded")
		}

		_, err = img.body.ReadFrom(resp.Body)
		return err
	}
	return errors.New("No images")

}

func (img *BingImages) WriteBody(w io.Writer) error {
	_, err := img.body.WriteTo(w)
	return err
}

func (img *BingImages) String() string {
	return "Bing daily wallpaper"
}

func (img *BingImages) Clean() {
	img.body = bytes.Buffer{}
}
