package geometry

import (
    rt "../raytracer"
    v "../vector"
)

type Plane struct {
    Origin, Normal v.Vector
}

func (p Plane) IntersectsRay(r rt.Ray, tMin float64, tMax float64) (bool, rt.Hit) {
    denom := r.Direction.Dot(p.Normal)
    if denom > -0.001 && denom < 0.001 {
        return false, rt.Hit {}
    }

    t := p.Origin.Sub(r.Origin).Dot(p.Normal) / r.Direction.Dot(p.Normal)
    if t > tMin && t < tMax {
        normalDir := 1.0
        if r.Direction.Dot(p.Normal) > 0 {
            normalDir = -1.0
        }
        return true, rt.Hit {
            r.Point(t),
            p.Normal.Scale(normalDir),
            r,
            t,
        }
    }
    return false, rt.Hit{}
}

