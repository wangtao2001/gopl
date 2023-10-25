package main

import (
	"image/color"
	"io"
	"math/rand"
	"os"
	"time"
)

var palette = []color.Color{color.Black, color.White}

const (
	widthIndex = 0
	blackIndex = 1
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	show(os.Stdout)
}

func show(out io.Writer) {

}
