package geo

import "math"

type Polygon struct {
	Vertices []Vec2
}

func (p Polygon) BoundingBox() Rectangle {
	bounds := Rectangle{
		Min: Vec2{
			X: p.Vertices[0].X,
			Y: p.Vertices[0].Y,
		},
		Max: Vec2{
			X: p.Vertices[0].X,
			Y: p.Vertices[0].Y,
		},
	}

	for _, v := range p.Vertices {
		bounds.Min.X = min(bounds.Min.X, v.X)
		bounds.Min.Y = min(bounds.Min.Y, v.Y)

		bounds.Max.X = max(bounds.Max.X, v.X)
		bounds.Max.Y = max(bounds.Max.Y, v.Y)
	}

	return bounds
}

func (p Polygon) Edge(i int) (Vec2, Vec2) {
	j := i - 1
	if j < 0 {
		j = len(p.Vertices) - 1
	}

	return p.Vertices[i], p.Vertices[j]
}

func (p Polygon) Edges() int {
	return len(p.Vertices)
}

func (p Polygon) Rotate(rad float32) Polygon {
	out := Polygon{
		Vertices: make([]Vec2, p.Edges()),
	}

	for i, v := range p.Vertices {
		out.Vertices[i] = v.Rotate(rad)
	}

	return out
}

func (p Polygon) Translate(delta Vec2) Polygon {
	out := Polygon{
		Vertices: make([]Vec2, p.Edges()),
	}

	for i, v := range p.Vertices {
		out.Vertices[i] = v.Add(delta)
	}

	return out
}

type Rectangle struct {
	Min, Max Vec2
}

func (r Rectangle) Dx() float32 {
	return r.Max.X - r.Min.X
}

func (r Rectangle) Dy() float32 {
	return r.Max.Y - r.Min.Y
}

func (r Rectangle) Inset(x, y float32) Rectangle {
	if r.Dx() < 2*x {
		r.Min.X = (r.Min.X + r.Max.X) / 2
		r.Max.X = r.Min.X
	} else {
		r.Min.X += x
		r.Max.X -= x
	}
	if r.Dy() < 2*y {
		r.Min.Y = (r.Min.Y + r.Max.Y) / 2
		r.Max.Y = r.Min.Y
	} else {
		r.Min.Y += y
		r.Max.Y -= y
	}
	return r
}

func (r Rectangle) Overlaps(s Rectangle) bool {
	return r.Min.X < s.Max.X && s.Min.X < r.Max.X && r.Min.Y < s.Max.Y && s.Min.Y < r.Max.Y
}

func (r Rectangle) WrapVec2(v Vec2) Vec2 {
	v.X = float32(float64(v.X) - float64(r.Max.X-r.Min.X)*math.Floor(float64((v.X-r.Min.X)/(r.Max.X-r.Min.X))))
	v.Y = float32(float64(v.Y) - float64(r.Max.Y-r.Min.Y)*math.Floor(float64((v.Y-r.Min.Y)/(r.Max.Y-r.Min.Y))))
	return v
}

type Vec2 struct {
	X, Y float32
}

func (v Vec2) Add(delta Vec2) Vec2 {
	return Vec2{
		X: v.X + delta.X,
		Y: v.Y + delta.Y,
	}
}

func (v Vec2) Rotate(rad float32) Vec2 {
	return Vec2{
		X: float32(math.Cos(float64(rad)))*v.X - float32(math.Sin(float64(rad)))*v.Y,
		Y: float32(math.Sin(float64(rad)))*v.X + float32(math.Cos(float64(rad)))*v.Y,
	}
}
