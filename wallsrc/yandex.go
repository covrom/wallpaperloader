package wallsrc

import (
	"bytes"
	"errors"
	"io"
	"net/http"
)

type YandexImages struct {
	body bytes.Buffer
}

func (img *YandexImages) Get() error {
	resp, err := http.Get("https://yandex.ru/images/today?size=1920x1080")
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return errors.New("Yandex: not downloaded")
	}

	_, err = img.body.ReadFrom(resp.Body)
	return err
}

func (img *YandexImages) WriteBody(w io.Writer) error {
	_, err := img.body.WriteTo(w)
	return err
}

func (img *YandexImages) String() string {
	return "Yandex Images daily"
}

func (img *YandexImages) Clean() {
	img.body = bytes.Buffer{}
}
