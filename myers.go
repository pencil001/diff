package diff

import (
	"errors"
)

type Myers struct {
	A    []string
	lenA int

	B    []string
	lenB int
}

func NewMyers(a, b []string) *Myers {
	return &Myers{
		A:    a,
		lenA: len(a),
		B:    b,
		lenB: len(b),
	}
}

func (m *Myers) idx(i int) int {
	return i + m.lenA + m.lenB
}

func (m *Myers) getShortestEdit() ([][]int, error) {
	max := m.lenA + m.lenB

	v := make([]int, 2*max+1)
	v[m.idx(1)] = 0

	trace := make([][]int, 0)

	x, y := 0, 0

	for d := 0; d <= max; d++ {
		for k := -d; k <= d; k += 2 {
			if k == -d || (k != d && v[m.idx(k-1)] < v[m.idx(k+1)]) {
				x = v[m.idx(k+1)]
			} else {
				x = v[m.idx(k-1)] + 1
			}

			y = x - k

			for x < m.lenA && y < m.lenB && m.A[x] == m.B[y] {
				x, y = x+1, y+1
			}

			v[m.idx(k)] = x

			if x >= m.lenA && y >= m.lenB {
				trace = append(trace, append([]int(nil), v...))

				return trace, nil
			}
		}

		trace = append(trace, append([]int(nil), v...))
	}

	return nil, errors.New("not found shortest edit")
}

func (m *Myers) backtrack() ([]EditPoint, error) {
	trace, err := m.getShortestEdit()
	if err != nil {
		return nil, err
	}

	x := m.lenA
	y := m.lenB

	points := make([]EditPoint, 0)

	for d := len(trace) - 1; d >= 0; d-- {
		v := trace[d]

		k := x - y
		prevK := k

		if k == -d || (k != d && v[m.idx(k-1)] < v[m.idx(k+1)]) {
			prevK = k + 1
		} else {
			prevK = k - 1
		}

		prevX := v[m.idx(prevK)]
		prevY := prevX - prevK

		for x > prevX && y > prevY {
			points = append([]EditPoint{{PrevX: x - 1, PrevY: y - 1, X: x, Y: y}}, points...)
			x, y = x-1, y-1
		}

		if d > 0 {
			points = append([]EditPoint{{PrevX: prevX, PrevY: prevY, X: x, Y: y}}, points...)
		}

		x, y = prevX, prevY
	}

	return points, nil
}

func (m *Myers) Diff() ([]EditDiff, error) {
	points, err := m.backtrack()
	if err != nil {
		return nil, err
	}

	diffs := make([]EditDiff, 0)

	for _, p := range points {
		if p.X == p.PrevX {
			diffs = append(diffs, EditDiff{
				OpType: EditAdd,
				Old:    "",
				New:    m.B[p.PrevY],
			})
		} else if p.Y == p.PrevY {
			diffs = append(diffs, EditDiff{
				OpType: EditDel,
				Old:    m.A[p.PrevX],
				New:    "",
			})
		} else {
			diffs = append(diffs, EditDiff{
				OpType: EditEq,
				Old:    m.A[p.PrevX],
				New:    m.B[p.PrevY],
			})
		}
	}

	return diffs, nil
}
