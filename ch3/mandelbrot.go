// 使用complex128复数算法来生成一个Mandelbrot图像(png格式)
package main

import (
	"image/color"
	"math/cmplx"
)

// func main() {
// 	const (
// 		// 每个点的两个嵌套的循环对应 -2到+2区间的复数平面
// 		xmin, ymin, xmax, ymax = -2, -2, +2, +2
// 		// 图片大小 1024*1024
// 		width, height = 1024, 1024
// 	)

// 	img := image.NewRGBA(image.Rect(0, 0, width, height))
// 	for py := 0; py < height; py++ {
// 		y := float64(py)/height*(ymax-ymin) + ymin
// 		for px := 0; px < width; px++ {
// 			x := float64(px)/width*(xmax-xmin) + xmin
// 			z := complex(x, y)
// 			// 图像的点(px,py) 代表复数z的值.
// 			img.Set(px, py, mandelbrot(z))
// 		}
// 	}
// 	png.Encode(os.Stdout, img) // 忽略错误
// }

func mandelbrot(z complex128) color.Color {
	const iterations = 200
	const contrast = 15

	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			return color.Gray{255 - contrast*n}
		}
	}
	return color.Black
}
