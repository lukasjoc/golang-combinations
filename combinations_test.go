package combinations

import (
	"reflect"
	"testing"
)

func TestStringCombinationsForStringArrays(t *testing.T) {
	tt := []struct {
		name string
		in   []string
		out  [][]string
	}{
		{
			name: "Empty slice",
			in:   []string{},
			out:  nil,
		},
		{
			name: "Single item",
			in:   []string{"A"},
			out: [][]string{
				{"A"},
			},
		},
		{
			name: "Two items",
			in:   []string{"A", "B"},
			out: [][]string{
				{"A"},
				{"B"},
				{"A", "B"},
			},
		},
		{
			name: "Three items",
			in:   []string{"A", "B", "C"},
			out: [][]string{
				{"A"},
				{"B"},
				{"A", "B"},
				{"C"},
				{"A", "C"},
				{"B", "C"},
				{"A", "B", "C"},
			},
		},
		{
			name: "Four items",
			in:   []string{"A", "B", "C", "D"},
			out: [][]string{
				{"A"},
				{"B"},
				{"A", "B"},
				{"C"},
				{"A", "C"},
				{"B", "C"},
				{"A", "B", "C"},
				{"D"},
				{"A", "D"},
				{"B", "D"},
				{"A", "B", "D"},
				{"C", "D"},
				{"A", "C", "D"},
				{"B", "C", "D"},
				{"A", "B", "C", "D"},
			},
		},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			out := All[string](tc.in)
			if !reflect.DeepEqual(out, tc.out) {
				t.Errorf("error: \nreturn:\t%v\nwant:\t%v", out, tc.out)
			}
		})
	}
}

func TestStringCombinationsNForStringArrays(t *testing.T) {
	tt := []struct {
		name string
		in   []string
		n    int
		out  [][]string
	}{
		{
			name: "Empty slice",
			in:   []string{},
			n:    1,
			out:  nil,
		},
		{
			name: "Single item",
			in:   []string{"A"},
			n:    1,
			out: [][]string{
				{"A"},
			},
		},
		{
			name: "Two items, n = 0",
			in:   []string{"A", "B"},
			n:    0,
			out: [][]string{
				{"A"},
				{"B"},
				{"A", "B"},
			},
		},
		{
			name: "Two items, n = 1",
			in:   []string{"A", "B"},
			n:    1,
			out: [][]string{
				{"A"},
				{"B"},
			},
		}, {
			name: "Two items, n = 2",
			in:   []string{"A", "B"},
			n:    2,
			out: [][]string{
				{"A", "B"},
			},
		},
		{
			name: "Three items, n = 0",
			in:   []string{"A", "B", "C"},
			n:    0,
			out: [][]string{
				{"A"},
				{"B"},
				{"A", "B"},
				{"C"},
				{"A", "C"},
				{"B", "C"},
				{"A", "B", "C"},
			},
		},
		{
			name: "Three items, n = 1",
			in:   []string{"A", "B", "C"},
			n:    1,
			out: [][]string{
				{"A"},
				{"B"},
				{"C"},
			},
		},
		{
			name: "Three items, n = 2",
			in:   []string{"A", "B", "C"},
			n:    2,
			out: [][]string{
				{"A", "B"},
				{"A", "C"},
				{"B", "C"},
			},
		},
		{
			name: "Three items, n = 3",
			in:   []string{"A", "B", "C"},
			n:    3,
			out: [][]string{
				{"A", "B", "C"},
			},
		},
		{
			name: "Three items, n = 4",
			in:   []string{"A", "B", "C"},
			n:    4,
			out: [][]string{
				{"A", "B", "C"},
			},
		},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			out := Combinations[string](tc.in, tc.n)
			if !reflect.DeepEqual(out, tc.out) {
				t.Errorf("error: \nreturn:\t%v\nwant:\t%v", out, tc.out)
			}
		})
	}
}

func TestStringCombinationsNForIntArrays(t *testing.T) {
	tt := []struct {
		name string
		in   []uint64
		n    int
		out  [][]uint64
	}{
		{
			name: "Empty slice",
			in:   []uint64{},
			n:    1,
			out:  nil,
		},
		{
			name: "Single item",
			in:   []uint64{1},
			n:    1,
			out: [][]uint64{
				{1},
			},
		},
		{
			name: "Two items, n = 0",
			in:   []uint64{1, 2},
			n:    0,
			out: [][]uint64{
				{1},
				{2},
				{1, 2},
			},
		},
		{
			name: "Two items, n = 1",
			in:   []uint64{1, 2},
			n:    1,
			out: [][]uint64{
				{1},
				{2},
			},
		}, {
			name: "Two items, n = 2",
			in:   []uint64{1, 2},
			n:    2,
			out: [][]uint64{
				{1, 2},
			},
		},
		{
			name: "Three items, n = 0",
			in:   []uint64{1, 2, 3},
			n:    0,
			out: [][]uint64{
				{1},
				{2},
				{1, 2},
				{3},
				{1, 3},
				{2, 3},
				{1, 2, 3},
			},
		},
		{
			name: "Three items, n = 1",
			in:   []uint64{1, 2, 3},
			n:    1,
			out: [][]uint64{
				{1},
				{2},
				{3},
			},
		},
		{
			name: "Three items, n = 2",
			in:   []uint64{1, 2, 3},
			n:    2,
			out: [][]uint64{
				{1, 2},
				{1, 3},
				{2, 3},
			},
		},
		{
			name: "Three items, n = 3",
			in:   []uint64{1, 2, 3},
			n:    3,
			out: [][]uint64{
				{1, 2, 3},
			},
		},
		{
			name: "Three items, n = 4",
			in:   []uint64{1, 2, 3},
			n:    4,
			out: [][]uint64{
				{1, 2, 3},
			},
		},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			out := Combinations[uint64](tc.in, tc.n)
			if !reflect.DeepEqual(out, tc.out) {
				t.Errorf("error: \nreturn:\t%v\nwant:\t%v", out, tc.out)
			}
		})
	}
}
