package main

import (
    img "./img"
    v "./vector"
)

func main() {
    const (
        w = 1000
        h = 1000
    )

    pixels := make([][]v.Vector, h)
    for y := 0; y < h; y++ {
        pixels[y] = make([]v.Vector, w)
        py := float64(y) / float64(h)
        for x := 0; x < w; x++ {
            px := float64(x) / float64(w)
            pixels[y][x] = v.Vector{px, py, 0.2}
        }
    }
    i := img.Img{img.Dim {w,h}, pixels}
    i.Render("out/gradient.png")
}
