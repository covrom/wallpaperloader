package main

import (
	"crypto/rand"
	"log"
	"os"
	"time"

	"github.com/covrom/wallpaperloader/wallsrc"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Usage: filename suffix for saving images")
	}

	fn := os.Args[1]
	ui := &wallsrc.UnsplashImages{}
	// yi := &wallsrc.YandexImages{}
	bi := &wallsrc.BingImages{}

	for {
		randfs := make([]wallsrc.Source, 0, 3)

		if err := ui.Get(); err == nil {
			randfs = append(randfs, ui)
			log.Println("Unsplash OK")
		} else {
			log.Println(err)
		}
		// if err := yi.Get(); err == nil {
		// 	randfs = append(randfs, yi)
		// 	log.Println("Yandex OK")
		// } else {
		// 	log.Println(err)
		// }
		if err := bi.Get(); err == nil {
			randfs = append(randfs, bi)
			log.Println("Bing OK")
		} else {
			log.Println(err)
		}

		if len(randfs) > 0 {
			rnd := make([]byte, 1)
			rand.Read(rnd)
			idx := int(rnd[0]) % len(randfs)

			f, err := os.Create(fn)
			if err == nil {
				randfs[idx].WriteBody(f)
				f.Close()
				log.Println("Writed", randfs[idx])
			}
			for _, src := range randfs {
				src.Clean()
			}
			time.Sleep(time.Hour)
		} else {
			// if all sources is fail, try after 5 minutes
			time.Sleep(5 * time.Minute)
		}
	}
}
