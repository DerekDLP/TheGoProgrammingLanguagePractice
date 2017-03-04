// 随机的lissajous图形产生gif动画
// go build
// lissajous >out.gif
package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
	"os"
)
// 原示例
// var palette = []color.Color{color.White, color.Black}
// exercise 1.5 更改颜色:黑色-->绿色
var palette = []color.Color{color.White, color.RGBA{0, 0xff, 0, 0xff}}

// 调色板
const (
	// 原示例
	 whiteIndex = 0
	 blackIndex = 1
)

func main()  {
	lissajous(os.Stdout)
}

func lissajous(out io.Writer)  {
	const (
		// 完整的x振荡器转数
		cycles = 5
		// 角分辨率
		res = 0.001
		// 绘图大小[-size...+size]
		size = 100
		// 动画的帧数
		nframes = 64
		// 10ms帧间延迟单元
		delay = 8
	)
	// Y振荡器的相对频率 
	freq := rand.Float64() * 3.0
	anim := gif.GIF{LoopCount: nframes}
	// 相位差
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
	// 忽略encoding的错误
	gif.EncodeAll(out, &anim)
}