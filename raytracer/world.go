package raytracer

type Surface interface {
    IntersectsRay(r Ray, tMin float64, tMax float64) (bool, Hit)
}

type World struct {
    Items []Surface
}

func (w World) IntersectsRay(r Ray, tMin float64, tMax float64) (bool, Hit) {
    hitAnything := false
    closest := tMax
    h := Hit{}

    for _, item := range w.Items {
        didHit, tempH := item.IntersectsRay(r, tMin, closest)

        if didHit && tempH.T < closest {
            hitAnything = true
            h = tempH
            closest = h.T
        }
    }
    return hitAnything, h
}

