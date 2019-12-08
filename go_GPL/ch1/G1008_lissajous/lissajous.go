package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"log"
	"math"
	"math/rand"
	"net/http"
	"os"
	"time"
)

var palette = []color.Color{color.White, color.Black}

const (
	// 画板中的第一种颜色
	whiteIndex = 0
	// 画板中的下一种颜色
	blackIndex = 1
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	if len(os.Args) > 1 && os.Args[1] == "web" {
		handler := func(w http.ResponseWriter, r *http.Request) {
			lissajous(w)
		}
		http.HandleFunc("/", handler)
		log.Fatal(http.ListenAndServe("localhost:8080", nil))
		return
	}
	lissajous(os.Stdout)
}

func lissajous(out io.Writer) {
	const (
		// 完整的x振荡器变化的个数
		cycles = 5
		// 角度分辨率
		res = 0.001
		// 图像画布包含[-size..+size]
		size = 100
		// 动画中的帧数
		nframes = 64
		// 以10毫秒为单位的帧间延时
		delay = 8
	)
	// y振荡器的相对频率
	freq := rand.Float64() * 3.0
	anim := gif.GIF{LoopCount: nframes}
	// phase difference
	phase := 0.0
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), blackIndex)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	// 注意：忽略编码错误
	gif.EncodeAll(out, &anim)
}
