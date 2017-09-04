package material

import (
    v "../vector"
    rt "../raytracer"
)

type Reflective struct {
    C v.Vector
    Fuzz float64
}

func (m Reflective) Scatter(h rt.Hit) (rt.Ray, bool) {
    newDirection := h.R.Direction.ReflectAbout(h.Normal)
    fuzzVector := v.RandomVector().Scale(m.Fuzz)
    r := rt.Ray {
        Origin: h.Point,
        Direction: newDirection.Add(fuzzVector).Normalize(),
        Color: m.C.Mult(h.R.Color),
    }
    return r, false
}

