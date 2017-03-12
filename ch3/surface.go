// 通过浮点计算生成的图形
// 带有两个参数的z = f(x, y)函数的三维形式，使用了可缩放矢量图形（SVG）格式输出，SVG是一个用于矢量线绘制的XML标准
//
// sin(r)/r函数的输出图形，其中r是sqrt(xx+yy)
package main

import (
	"fmt"
	"math"
)

const (
	// 画板大小,单位像素
	width, height = 600, 320
	// 网格单元数
	cells = 100
	// 轴范围 (-xyrange..+xyrange)
	xyrange = 30.0
	// 每个x/y单元像素
	xyscale = width / 2 / xyrange
	// 每个z单元像素, 0.4是任意的缩放系数
	zscale = height * 0.4
	// x,y轴角 (30°)
	angle = math.Pi / 6
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle)

func main()  {
	fmt.Printf("<svg xmlns='http://www.w3.org/2000/svg' "+
	    "style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", width, height)
		for i := 0; i < cells; i++ {
			for j := 0; j < cells; j++ {
				// 计算多边形ABCD
				// B对应(i,j)顶点位置，A、C、D是其相邻的顶点
				ax, ay := corner(i+1, j)
				bx, by := corner(i, j)
				cx, cy := corner(i, j+1)
				dx, dy := corner(i+1, j+1)
				// 输出SVG绘图指令
				fmt.Printf("<polygon points='%g,%g %g,%g %g,%g %g,%g'/>\n",
				    ax, ay, bx, by, cx, cy, dx, dy)
			}
		}
		fmt.Println("</svg>")
}

func corner(i, j int) (float64, float64) {
	// 网格顶点(i,j)对应的参数坐标(x,y)
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	// 计算表面高度
    z := f(x, y)

    // 将(x,y,z)等角投影的绘画到二维svg画布上(sx,sy)
	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy
}

func f(x, y float64) float64 {
	// 点(x,y)与(0.0)的距离
	r := math.Hypot(x, y)
	return math.Sin(r)/r
}