// lissajous服务器,通过请求获取lissajous参数
// eg: http://localhost:8000/?cycles=20
package main
import (
	"log"
	"net/http"
	"io"
	"math"
	"math/rand"
	"image"
	"image/color"
	"image/gif"
	"strconv"
	"strings"
)

var palette = []color.Color{color.White, color.RGBA{0, 0xff, 0, 0xff}}

const (
	 whiteIndex = 0
	 blackIndex = 1
)

func main()  {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		log.Print(err)
		return 
	}
	lissajous(w, r)	
}

func lissajous(out io.Writer, r *http.Request)  {
	const (
		// 完整的x振荡器转数
		// cycles = 5
		// 角分辨率
		res = 0.001
		// 绘图大小[-size...+size]
		size = 100
		// 动画的帧数
		nframes = 64
		// 10ms帧间延迟单元
		delay = 8
	)
	var cycles int
	for k, v := range r.Form {
		// 判断字符串是否相等，不区分大小写
		if strings.EqualFold(k, "cycles") {
			cycles, _ = strconv.Atoi(v[0])
		}
	}	
	// Y振荡器的相对频率 
	freq := rand.Float64() * 3.0
	anim := gif.GIF{LoopCount: nframes}
	// 相位差
	phase := 0.0
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < math.Pi*2*float64(cycles); t += res {
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