package main

import (
	"crypto/rand"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"

	"github.com/covrom/wallpaperloader/wallsrc"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Usage: filename for saving images")
	}

	fn := os.Args[1]
	dir := filepath.Dir(fn)
	if err := os.MkdirAll(dir, 0755); err != nil {
		log.Fatal(err)
	}
	blurIdx := strings.LastIndexByte(fn, '/')
	blurfn := fn[:blurIdx] + "/blur" + fn[blurIdx:]
	dirblur := filepath.Dir(blurfn)
	if err := os.MkdirAll(dirblur, 0755); err != nil {
		log.Fatal(err)
	}

	sources := []wallsrc.Source{
		&wallsrc.UnsplashImagesRnd{},
		&wallsrc.UnsplashImagesWallpaper{},
		// &wallsrc.YandexImages{},
		&wallsrc.BingImages{},
	}

	for {
		randfs := make([]wallsrc.Source, 0, 3)

		for _, src := range sources {
			if err := src.Get(); err == nil {
				randfs = append(randfs, src)
				log.Println(src, "OK")
			} else {
				src.Clean()
				log.Println(err)
			}
		}

		if len(randfs) > 0 {
			rnd := make([]byte, 1)
			rand.Read(rnd)
			idx := int(rnd[0]) % len(randfs)

			f, err := os.OpenFile(fn, os.O_WRONLY|os.O_CREATE|os.O_TRUNC|os.O_SYNC, 0755)
			if err == nil {
				randfs[idx].WriteBody(f)
				f.Close()

				// convert "$FILE" -level -50%,100%,0.6 -filter Gaussian -resize 20% -define filter:sigma=2.5 -resize 500% -fill white -gravity center "$IMAGE".jpg
				if err := exec.Command(
					"convert",
					fn,
					"-level",
					"-50%,100%,0.6",
					"-filter",
					"Gaussian",
					"-resize",
					"20%",
					"-define",
					"filter:sigma=2.5",
					"-resize",
					"500%",
					"-fill",
					"white",
					"-gravity",
					"center",
					blurfn,
				).Run(); err != nil {
					log.Println(err)
				}

				log.Println("Writed", randfs[idx])
			} else {
				log.Println(err)
			}

			f = nil

			for _, src := range randfs {
				src.Clean()
			}

			time.Sleep(time.Hour)
		} else {
			// if all sources fail, try after 5 minutes
			time.Sleep(5 * time.Minute)
		}
	}
}
