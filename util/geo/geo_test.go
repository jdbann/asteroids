package geo_test

import (
	"testing"

	"github.com/jdbann/asteroids/util/geo"
	"gotest.tools/v3/assert"
)

func TestRectangle_WrapVec2(t *testing.T) {
	type testCase struct {
		name      string
		rectangle geo.Rectangle
		vec2      geo.Vec2
		want      geo.Vec2
	}

	run := func(t *testing.T, tc testCase) {
		got := tc.rectangle.WrapVec2(tc.vec2)
		assert.DeepEqual(t, got, tc.want)
	}

	testCases := []testCase{
		{
			name:      "up and left of rectangle at origin",
			rectangle: geo.Rectangle{Max: geo.Vec2{X: 10, Y: 10}},
			vec2:      geo.Vec2{X: -2, Y: -2},
			want:      geo.Vec2{X: 8, Y: 8},
		},

		{
			name:      "inside rectangle at origin",
			rectangle: geo.Rectangle{Max: geo.Vec2{X: 10, Y: 10}},
			vec2:      geo.Vec2{X: 5, Y: 5},
			want:      geo.Vec2{X: 5, Y: 5},
		},
		{
			name:      "down and right of rectangle at origin",
			rectangle: geo.Rectangle{Max: geo.Vec2{X: 10, Y: 10}},
			vec2:      geo.Vec2{X: 12, Y: 12},
			want:      geo.Vec2{X: 2, Y: 2},
		},
		{
			name:      "up and left of rectangle in negative space",
			rectangle: geo.Rectangle{Min: geo.Vec2{X: -20, Y: -20}, Max: geo.Vec2{X: -10, Y: -10}},
			vec2:      geo.Vec2{X: -22, Y: -22},
			want:      geo.Vec2{X: -12, Y: -12},
		},
		{
			name:      "inside rectangle in negative space",
			rectangle: geo.Rectangle{Min: geo.Vec2{X: -20, Y: -20}, Max: geo.Vec2{X: -10, Y: -10}},
			vec2:      geo.Vec2{X: -18, Y: -18},
			want:      geo.Vec2{X: -18, Y: -18},
		},
		{
			name:      "down and right of rectangle in negative space",
			rectangle: geo.Rectangle{Min: geo.Vec2{X: -20, Y: -20}, Max: geo.Vec2{X: -10, Y: -10}},
			vec2:      geo.Vec2{X: -8, Y: -8},
			want:      geo.Vec2{X: -18, Y: -18},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			run(t, tc)
		})
	}
}

func TestVec2_InPolygon(t *testing.T) {
	type testCase struct {
		name    string
		vec2    geo.Vec2
		polygon geo.Polygon
		want    bool
	}

	run := func(t *testing.T, tc testCase) {
		got := tc.vec2.InPolygon(tc.polygon)
		assert.Equal(t, got, tc.want)
	}

	testCases := []testCase{
		{
			name: "outside square",
			vec2: geo.Vec2{X: 15, Y: 5},
			polygon: geo.Polygon{
				Vertices: []geo.Vec2{
					{0, 0},
					{0, 10},
					{10, 10},
					{10, 0},
				},
			},
			want: false,
		},
		{
			name: "inside square",
			vec2: geo.Vec2{X: 5, Y: 5},
			polygon: geo.Polygon{
				Vertices: []geo.Vec2{
					{0, 0},
					{0, 10},
					{10, 10},
					{10, 0},
				},
			},
			want: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			run(t, tc)
		})
	}
}
