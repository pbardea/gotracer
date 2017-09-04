package material

import (
    v "../vector"
    rt "../raytracer"
)

type Emitter struct {
    C v.Vector
}

func (m Emitter) Scatter(h rt.Hit) (rt.Ray, bool) {
    return rt.Ray{h.Point, h.Normal}, false
}

func (m Emitter) Color() v.Vector {
    return m.C
}

