package main

import (
	"fmt"
	"image/png"
	"math/rand"
	"os"
	"time"

	"github.com/kbinani/screenshot"
	"gopkg.in/toast.v1"
)

const (
	delayKeyfetchMS = 30
)

func main() {
	kl := NewKeylogger()
	for {
		key := kl.GetKey()
		if !key.Empty {
			n := screenshot.NumActiveDisplays()

			for i := 0; i < n; i++ {
				t := time.Now()
				rand.Seed(time.Now().UnixNano())
				randomid := (rand.Intn(100-1+1) + 1)
				bounds := screenshot.GetDisplayBounds(i)
				img, _ := screenshot.CaptureRect(bounds)
				userPath := os.Getenv("USERPROFILE")
				fileName := fmt.Sprintf("%d_%d_%s.png", i, randomid, t.Format("20060102150405"))
				filePath := fmt.Sprintf("%s\\Pictures\\%d_%d_%s.png", userPath, i, randomid, t.Format("20060102150405"))

				file, _ := os.Create(filePath)
				png.Encode(file, img)
				file.Close()
				notification := toast.Notification{
					AppID:               "Go PrtSc",
					Title:               "Screenshot taken",
					Icon:                filePath,
					Message:             fileName,
					ActivationArguments: filePath,
				}
				notification.Push()
			}

		}

		time.Sleep(delayKeyfetchMS * time.Millisecond)
	}

}
