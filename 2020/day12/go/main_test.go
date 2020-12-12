package main

import (
	"testing"
)

func TestRotate(t *testing.T) {

	tcs := []struct {
		h, v, dir int

		eh, ev int
	}{
		{
			// N, E -> E, S
			h: 10, v: 1,
			eh: 1, ev: -10,
			dir: 1,
		},
		{
			// E, S -> S, W
			h: 1, v: -10,
			eh: -10, ev: -1,
			dir: 1,
		},
		{
			// S, W -> W, N
			h: -10, v: -1,
			eh: -1, ev: 10,
			dir: 1,
		},
		{
			// W, N -> N, E
			h: -1, v: 10,
			eh: 10, ev: 1,
			dir: 1,
		},

		{
			// N, E -> W, N
			h: 10, v: 1,
			eh: -1, ev: 10,
			dir: -1,
		},
		{
			// W, N -> S, W
			h: -1, v: 10,
			eh: -10, ev: -1,
			dir: -1,
		},
		{
			// S, W -> E, S
			h: -10, v: -1,
			eh: 1, ev: -10,
			dir: -1,
		},
		{
			// E, S -> N, E
			h: 1, v: -10,
			eh: 10, ev: 1,
			dir: -1,
		},
	}

	for i, tc := range tcs {
		v, h := rotate(90, tc.dir, tc.v, tc.h)
		if v != tc.ev {
			t.Errorf("%d) v %d != %d", i+1, tc.ev, v)
		}
		if h != tc.eh {
			t.Errorf("%d) h %d != %d", i+1, tc.eh, h)
		}
	}
}
