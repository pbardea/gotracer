package main

import (
    "math"
    "math/rand"

    img "./img"
    v "./vector"
    rt "./raytracer"
    g "./geometry"
    m "./material"
)

func getColor(r rt.Ray, world rt.Surface, depth int) v.Vector {
    const maxBounce = 20
    didHit, h := world.IntersectsRay(r, 0.001, math.MaxFloat64)

    if didHit {
        if depth < maxBounce {
            bouncedRay, bounced := h.Material.Scatter(h)
            if bounced {
                return h.Material.Color().Mult(getColor(bouncedRay, world, depth + 1))
            } else {
                return h.Material.Color()
            }
        }
        return v.Vector{0.0, 0.0, 0.0}
    }
	return v.Vector{1, 1, 1}
}

func main() {
    const (
        w = 200
        h = 200

        // Anti-aliasing sampling rate
        aas = 50
    )

    c := rt.NewCamera()

    s := g.Sphere{v.Vector{0, 0, -1}, 0.5, m.Diffuse{v.Vector{1, 0.2, 0.3}}}
    p := g.Plane{v.Vector{0, -1, 0}, v.Vector{0, 1, 0}, m.Diffuse{v.Vector{1, 0.8, 0.8}}}

    world := rt.World{[]rt.Surface{s, p}}
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
                rgb = rgb.Add(getColor(r, world, 0))
            }
            rgb = rgb.Scale(1.0 / float64(aas))
            pixels[y][x] = rgb
        }
    }
    i := img.Img{img.Dim {w,h}, pixels}
    i.Render("out/basic.png")
}

