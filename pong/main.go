package main

import (
	"fmt"

	"github.com/veandco/go-sdl2/sdl"
)

func main() {

	winWidth := 800
	winHeight := 600

	window, err := sdl.CreateWindow("Go SDL2 Window", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, int32(winWidth), int32(winHeight), sdl.WINDOW_SHOWN)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer window.Destroy()

	renderer, err := sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer renderer.Destroy()

	tex, err := sdl.CreateTexture(sdl.PIXELFORMAT_ABGR8888, sdl.TEXTUREACCESS_STREAMING, int32(winWidth), int32(winHeight))
	if err != nil {
		fmt.Println(err)
		return
	}

	defer tex.Destroy()

	pixels := make([]byte, winWidth*winHeight*4)

	sdl.Delay(2000)

}
