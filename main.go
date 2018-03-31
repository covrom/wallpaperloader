package main

import (
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"
	"os"
)

func getYa(fn string) error {
	resp, err := http.Get("https://yandex.ru/images/today?size=1920x1080")
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return errors.New("Yandex: not downloaded")
	}

	f, err := os.Create(fn)
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = io.Copy(f, resp.Body)
	return err
}

func getBing(fn string) error {
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

		f, err := os.Create(fn)
		if err != nil {
			return err
		}
		defer f.Close()

		_, err = io.Copy(f, resp.Body)
		return err
	}
	return errors.New("No images")
}

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Usage: filename for saving image")
	}
	if err := getYa(os.Args[1]); err == nil {
		log.Println("Yandex OK")
		return
	} else {
		log.Println(err)
	}
	if err := getBing(os.Args[1]); err == nil {
		log.Println("Bing OK")
		return
	} else {
		log.Println(err)
	}
}
