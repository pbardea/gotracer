package geometry

import (
    "math"

    rt "../raytracer"
    v "../vector"
)

type Sphere struct {
    Center v.Vector
    Radius float64
}

func (s Sphere) IntersectsRay(r rt.Ray, tMin float64, tMax float64) (bool, rt.Hit) {
    // Quadratic formula
    toCenter := r.Origin.Sub(s.Center)
    a := r.Direction.Dot(r.Direction)
    b := 2.0 * toCenter.Dot(r.Direction)
    c := toCenter.Dot(toCenter) - s.Radius*s.Radius
    discr := b*b - 4*a*c

    h := rt.Hit{}

    if discr > 0 {
        t := (-b - math.Sqrt(discr)) / (2*a)
        if t <= 0 {
            t = (-b + math.Sqrt(discr)) / (2*a)
        }
        if t > tMin && t < tMax {
            h.T = t
            h.Point = r.Point(t)
            h.Normal = h.Point.Sub(s.Center).Normalize()
            return true, h
        }
    }
    return false, h
}

