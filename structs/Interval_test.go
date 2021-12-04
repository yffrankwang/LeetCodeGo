package structs

import (
	"testing"
)

func Test_Interval2Ints(t *testing.T) {
	actual := Interval2Ints(Interval{Start: 1, End: 2})
	expected := []int{1, 2}
	assertEqual(t, expected, actual)
}

func Test_IntervalSlice2Intss(t *testing.T) {
	actual := IntervalSlice2Intss(
		[]Interval{
			{
				Start: 1,
				End:   2,
			},
			{
				Start: 3,
				End:   4,
			},
		},
	)
	expected := [][]int{
		{1, 2},
		{3, 4},
	}

	assertEqual(t, expected, actual)
}
func Test_Intss2IntervalSlice(t *testing.T) {
	expected := []Interval{
		{
			Start: 1,
			End:   2,
		},
		{
			Start: 3,
			End:   4,
		},
	}
	actual := Intss2IntervalSlice([][]int{
		{1, 2},
		{3, 4},
	},
	)

	assertEqual(t, expected, actual)
}
