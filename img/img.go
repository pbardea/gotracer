package img

import (
    "fmt"
    "image"
    "image/color"
    "image/png"
    "os"

    v "../vector"
)

type Dim struct {
    W int
    H int
}

type Img struct {
    Dim Dim
    Pixels [][]v.Vector
}

func (i Img) Render(filename string) {
    const maxColor = 255.0

    m := image.NewRGBA(image.Rect(0, 0, i.Dim.W, i.Dim.H))
    for x := 0; x < i.Dim.W; x++ {
        for y := 0; y < i.Dim.H; y++ {
            rgb := i.Pixels[y][x]

            c := color.RGBA{
                uint8(maxColor * rgb.X),
                uint8(maxColor * rgb.Y),
                uint8(maxColor * rgb.Z),
                255,
            }
            m.Set(x, y, c)
        }
    }

    f, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE, 0600)
    if err != nil {
        fmt.Println(err)
        return
    }
    defer f.Close()
    png.Encode(f, m)
}
