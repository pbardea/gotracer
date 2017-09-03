package main

import (
    "math"
    "math/rand"

    img "./img"
    v "./vector"
    rt "./raytracer"
    g "./geometry"
)

type Surface interface {
    IntersectsRay(r rt.Ray, tMin float64, tMax float64) (bool, rt.Hit)
}

func getColor(r rt.Ray, world Surface, depth int) v.Vector {
    const MaxBounce = 20
    didHit, h := world.IntersectsRay(r, 0.001, math.MaxFloat64)

    if didHit {
        return h.Normal.Translate(1.0).Scale(0.5)
    }
    return v.Vector{1.0, 1.0, 1.0}
}

func main() {
    const (
        w = 1000
        h = 1000

        // Anti-aliasing sampling rate
        aas = 50
    )

    c := rt.NewCamera()
    s := g.Sphere{v.Vector{0.0, 0.0, -5.0}, 0.5}
    pixels := make([][]v.Vector, h)
    for y := 0; y < h; y++ {
        pixels[y] = make([]v.Vector, w)
        py := float64(y) / float64(h)
        for x := 0; x < w; x++ {
            px := float64(x) / float64(w)
            rgb := v.Vector{}
            for sample := 0; sample < aas; sample++ {
                yRand := rand.Float64() / float64(h)
                xRand := rand.Float64() / float64(w)
                r := c.RayAt(px + xRand, py + yRand)
                rgb = rgb.Add(getColor(r, s, 0))
            }
            rgb = rgb.Scale(1.0 / float64(aas))
            pixels[y][x] = rgb
        }
    }
    i := img.Img{img.Dim {w,h}, pixels}
    i.Render("out/ballaa.png")
}

