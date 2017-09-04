package material

import (
    v "../vector"
    rt "../raytracer"
)

type Metal struct {
    C v.Vector
    Fuzz float64
}

func (m Metal) Scatter(h rt.Hit) (rt.Ray, bool) {
    direction := h.R.Direction.ReflectAbout(h.Normal)
    fuzzed := v.VectorInUnitSphere().Scale(m.Fuzz)
    bouncedRay := rt.Ray{h.Point, direction.Add(fuzzed)}
    // bounced := direction.Dot(h.Normal) > 0
    return bouncedRay, true
}

func (m Metal) Color() v.Vector {
    return m.C
}

