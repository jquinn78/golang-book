package math

import "testing"

type testpair struct {
	values  []float64
	average float64
}

type mint struct {
	values []float64
	min    float64
}

type maxt struct {
	values []float64
	max    float64
}

var tests = []testpair{
	{[]float64{1, 2}, 1.5},
	{[]float64{1, 1, 1, 1, 1, 1}, 1},
	{[]float64{-1, 1}, 0},
}

var tMin = []mint{
	{[]float64{5, 6}, 5},
	{[]float64{0, 1}, 0},
	{[]float64{-1, 1}, -1},
}

var tMax = []maxt{
	{[]float64{5, 6}, 6},
	{[]float64{0, 1}, 1},
	{[]float64{-1, 1}, 1},
}

func TestAverage(t *testing.T) {

	for _, pair := range tests {

		v := Average(pair.values)
		if v != pair.average {
			t.Error("Found", pair.values,
				"expected", pair.average,
				"got", v)

		}
	}
}

func TestMin(t *testing.T) {
	for _, pair := range tMin {

		v := Min(pair.values)
		if v != pair.min {
			t.Error("Found", pair.values,
				"expected", pair.min,
				"got", v)
		}
	}
}

func TestMax(t *testing.T) {
	for _, pair := range tMax {

		v := Max(pair.values)
		if v != pair.max {
			t.Error("Found", pair.values,
				"expected", pair.max,
				"got", v)
		}
	}
}
