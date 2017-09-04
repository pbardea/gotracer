package raytracer

import (
    v "../vector"
)

type Camera struct {
    lowerLeft, h, v, origin v.Vector
}

func NewCamera() Camera {
    c := Camera{
        lowerLeft: v.Vector{-1.0, -1.0, -1.0},
        h: v.Vector{2.0, 0.0, 0.0},
        v: v.Vector{0.0, 2.0, 0.0},
        origin: v.Vector{0.0, 0.0, 0.0},
    }

    return c
}

func (c Camera) RayAt(x float64, y float64) Ray {
    position := c.h.Scale(x).Add(c.v.Scale(1-y))
    direction := c.lowerLeft.Add(position)

    return Ray{c.origin, direction}
}

