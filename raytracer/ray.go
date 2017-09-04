package raytracer

import (
    v "../vector"
)

type Material interface {
    Scatter(h Hit) (Ray, bool)
}

type Ray struct {
    Origin, Direction, Color v.Vector
}

func (r Ray) Point(t float64) v.Vector {
    d := r.Direction.Scale(t)
    o := r.Origin
    return o.Add(d)
}

type Hit struct {
    Point, Normal v.Vector
    R Ray
    T float64
    Material Material
}

