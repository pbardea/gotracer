package vector

import (
    "math"
    "math/rand"
)

type Vector struct {
    X, Y, Z float64
}

func (v Vector) Add(w Vector) Vector {
    return Vector{v.X + w.X, v.Y + w.Y, v.Z + w.Z}
}

func (v Vector) Sub(w Vector) Vector {
    return Vector{v.X - w.X, v.Y - w.Y, v.Z - w.Z}
}

func (v Vector) Dot(w Vector) float64 {
    return v.X*w.X + v.Y*w.Y + v.Z*w.Z
}

func (v Vector) Mag() float64 {
    return math.Sqrt(v.Dot(v))
}

func (v Vector) Scale(t float64) Vector {
    return Vector{v.X * t, v.Y * t, v.Z * t}
}

func (v Vector) Mult(w Vector) Vector {
    return Vector{v.X * w.X, v.Y * w.Y, v.Z * w.Z}
}

func (v Vector) Normalize() Vector {
    l := v.Mag()
    return Vector{v.X / l, v.Y / l, v.Z / l}
}

func (v Vector) Translate(s float64) Vector {
    return Vector{v.X + s, v.Y + s, v.Z + s}
}

func (v Vector) ReflectAbout(n Vector) Vector {
    return v.Sub(n.Scale(2.0 * v.Dot(n)))
}

func RandomVector() Vector {
    x := rand.Float64()
    y := rand.Float64()
    z := rand.Float64()
    xSign := rand.Int() % 2 == 0
    ySign := rand.Int() % 2 == 0
    zSign := rand.Int() % 2 == 0
    if !xSign { x *= -1  }
    if !ySign { y *= -1  }
    if !zSign { z *= -1  }
    return Vector{x, y, z}
}

var UnitVector = Vector{1, 1, 1}

func VectorInUnitSphere() Vector {
    for {
        r := Vector{rand.Float64(), rand.Float64(), rand.Float64()}
        p := r.Scale(2.0).Sub(UnitVector)
        if p.Dot(p) >= 1.0 {
            return p
        }
    }
}
