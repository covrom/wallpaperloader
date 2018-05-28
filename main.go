package main

import (
	"crypto/rand"
	"encoding/json"
	"errors"
	"io"
	"log"
	"math/big"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

const (
	ltmpy = "/tmp/wallpaperloader.Yandex"
	ltmpb = "/tmp/wallpaperloader.Bing"
	ltmpu = "/tmp/wallpaperloader.Unsplash"
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

func getUnsplash(fn string) error {
	resp, err := http.Get("https://source.unsplash.com/random/1920x1080")
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return errors.New("Unsplash: not downloaded")
	}

	f, err := os.Create(fn)
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = io.Copy(f, resp.Body)
	return err
}

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Usage: filename suffix for saving images")
	}
	fn := os.Args[1]
	for {
		randfs := make([]string, 0, 3)

		if err := getUnsplash(ltmpu); err == nil {
			randfs = append(randfs, ltmpu)
			log.Println("Unsplash OK")
		} else {
			log.Println(err)
		}
		if err := getYa(ltmpy); err == nil {
			randfs = append(randfs, ltmpy)
			log.Println("Yandex OK")
		} else {
			log.Println(err)
		}
		if err := getBing(ltmpb); err == nil {
			randfs = append(randfs, ltmpb)
			log.Println("Bing OK")
		} else {
			log.Println(err)
		}

		idx, err := rand.Int(rand.Reader, big.NewInt(int64(len(randfs))))
		if err == nil {
			fromf := randfs[int(idx.Int64())]
			os.Rename(fromf, fn)
			log.Println("Choose", filepath.Ext(fromf))
		}

		time.Sleep(time.Hour)
	}
}
