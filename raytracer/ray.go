package raytracer

import (
    v "../vector"
)

type Ray struct {
    Origin, Direction v.Vector
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
}
