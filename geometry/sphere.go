package geometry

import (
    rt "../raytracer"
    v "../vector"
)

type Sphere struct {
    Center v.Vector
    Radius float64
}

func (s Sphere) IntersectsRay(r rt.Ray, tMin float64, tmax float64) (bool, rt.Hit) {
    h := rt.Hit{}
    return false, h
}

