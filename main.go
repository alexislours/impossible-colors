package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"log"
	"os"
	"strconv"
	"time"
)

func main() {
	if len(os.Args) < 5 {
		log.Fatalf("Error: You must give the width, heigh and two color hex.")
	}
	fmt.Printf("Generating image...\n")
	start := time.Now()

	width, _ := strconv.Atoi(os.Args[1])
	height, _ := strconv.Atoi(os.Args[2])
	color1, _ := parseHexColor(os.Args[3])
	color2, _ := parseHexColor(os.Args[4])

	imgPath := fmt.Sprint(width) + "x" + fmt.Sprint(height) + "_" + os.Args[3] + "_" + os.Args[4] + ".png"
	img := createImage(width, height, color1, color2)

	out, err := os.Create(imgPath)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	png.Encode(out, img)
	end := time.Since(start)
	fmt.Printf("\nCreated image \"%v\". Took %.2fs\n", imgPath, end.Seconds())
}

func createImage(width int, height int, color1 color.RGBA, color2 color.RGBA) *image.RGBA {
	rect := image.Rect(0, 0, width, height)
	img := image.NewRGBA(rect)
	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			if x%2 == 0 {
				img.Set(x, y, color1)
				y++
				img.Set(x, y, color2)
			} else {
				img.Set(x, y, color2)
				y++
				img.Set(x, y, color1)
			}
		}
	}
	return img
}

func parseHexColor(s string) (c color.RGBA, err error) {
	if s[:1] != "#" {
		s = "#" + s
	}
	c.A = 0xff
	switch len(s) {
	case 7:
		_, err = fmt.Sscanf(s, "#%02x%02x%02x", &c.R, &c.G, &c.B)
	case 4:
		_, err = fmt.Sscanf(s, "#%1x%1x%1x", &c.R, &c.G, &c.B)
		c.R *= 17
		c.G *= 17
		c.B *= 17
	default:
		log.Fatalf("%v is not a valid hex code.", s)

	}
	return
}
