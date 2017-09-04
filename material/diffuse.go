package material

import (
    v "../vector"
    rt "../raytracer"
)

type Diffuse struct {
    C v.Vector
}

func (m Diffuse) Scatter(h rt.Hit) (rt.Ray, bool) {
    direction := h.Normal.Add(v.VectorInUnitSphere())
    return rt.Ray{h.Point, direction, m.C}, false
}

