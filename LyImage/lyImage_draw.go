package main

import (
	"fmt"
	"image"
	"image/color"
	"image/jpeg"
	"image/png"
	"log"
	"os"
)

func main() {
	DrawVerticalLine()
	DrawHorizLine()
}

//画垂直线
func DrawVerticalLine() {
	const (
		dx = 300
		dy = 500
	)

	// 需要保存的文件
	imgName := "Vertical"
	imgfile, _ := os.Create(fmt.Sprintf("%0s.png", imgName))
	defer imgfile.Close()

	// 新建一个 指定大小的 RGBA位图
	img := image.NewNRGBA(image.Rect(0, 0, dx, dy))

	for y := 0; y < dy; y++ {
		for x := 0; x < dx; x++ {

			if x%8 == 0 {
				// 设置某个点的颜色，依次是 RGBA
				img.Set(x, y, color.RGBA{uint8(x % 256), uint8(y % 256), 0, 255})
			}
		}
	}

	// 以PNG格式保存文件
	err := png.Encode(imgfile, img)
	if err != nil {
		log.Fatal(err)
	}
}

//画水平线
func DrawHorizLine() {
	const (
		dx = 600
		dy = 800
	)
	//创建保存文件
	imgName := "Line"
	imgFile, _ := os.Create(fmt.Sprintf("%0s.jpg", imgName))
	defer imgFile.Close()
	//新建一个 指定大小的 RGBA位图
	img := image.NewRGBA(image.Rect(0, 0, dx, dy))
	for i := 0; i < dx; i++ {
		img.Set(i, 200, color.RGBA{uint8(i % 256), uint8(i % 256), 0, 255})
	}
	//Encode函数 将采用JPEG 4:2:0基线格式和指定的编码质量将图像写入w。
	//如果o为nil将使用DefaultQuality。
	err := jpeg.Encode(imgFile, img, nil)
	if err != nil {
		log.Fatal(err)
	}
}

//求绝对值
func Abs(x int) int {
	if x >= 0 {
		return x
	}
	return -x
}

/*
func drawline(x0, y0, x1, y1 int, brush Putpixel) {
    dx := abs(x1 - x0)
    dy := abs(y1 - y0)
    sx, sy := 1, 1
    if x0 >= x1 {
    sx = -1
    }
    if y0 >= y1 {
    sy = -1
    }
    err := dx - dy

    for {
    brush(x0, y0)
    if x0 == x1 && y0 == y1 {
        return
    }
    e2 := err * 2
    if e2 > -dy {
        err -= dy
        x0 += sx
    }
    if e2 < dx {
        err += dx
        y0 += sy
    }
    }
}

func main() {

    const (
    dx = 300
    dy = 500
    )

    // 需要保存的文件

    // 新建一个 指定大小的 RGBA位图
    img := image.NewNRGBA(image.Rect(0, 0, dx, dy))

    drawline(5, 5, dx-8, dy-10, func(x, y int) {
    img.Set(x, y, color.RGBA{uint8(x), uint8(y), 0, 255})
    })

    // 左右都画一条竖线
    for i := 0; i < dy; i++ {
    img.Set(0, i, color.Black)
    img.Set(dx-1, i, color.Black)
    }

    imgcounter := 250
    imgfile, _ := os.Create(fmt.Sprintf("%03d.png", imgcounter))
    defer imgfile.Close()

    // 以PNG格式保存文件
    err := png.Encode(imgfile, img)
    if err != nil {
    log.Fatal(err)
    }
}
*/
