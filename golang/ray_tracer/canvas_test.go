package ray_tracer

import "testing"

func TestCanvas(t *testing.T) {
	width, height := 10, 20
	canvas := NewCanvas(width, height)
	if canvas.Width() != width {
		t.Errorf("canvas.Width() != %d", width)
	}
	if canvas.Height() != height {
		t.Errorf("canvas.Height() != %d", height)
	}
	black := NewColor(0, 0, 0)
	for i := 0; i < canvas.Height(); i++ {
		for j := 0; j < canvas.Width(); j++ {
			if canvas[i][j] != black {
				t.Errorf("canvas[%d][%d] != %+v", i, j, black)
			}
		}
	}
	x, y := 2, 3
	red := NewColor(1, 0, 0)
	canvas[x][y] = red
	for i := 0; i < canvas.Height(); i++ {
		for j := 0; j < canvas.Width(); j++ {
			if i == x && j == y {
				if canvas[i][j] != red {
					t.Errorf("canvas[%d][%d] != %+v", i, j, red)
				}
			} else if canvas[i][j] != black {
				t.Errorf("canvas[%d][%d] != %+v", i, j, black)
			}
		}
	}
}
