package geo

type Polygon struct {
	Vertices []Vec2
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

func (p Polygon) Translate(delta Vec2) Polygon {
	out := Polygon{
		Vertices: make([]Vec2, p.Edges()),
	}

	for i, v := range p.Vertices {
		out.Vertices[i] = v.Add(delta)
	}

	return out
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
