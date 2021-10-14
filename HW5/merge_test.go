package main

import (
	"math"
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Test struct {
	Funcs []MathFunction
} 

func TestMerge(t *testing.T) {

	testCase := &Test {
		Funcs: []MathFunction {
			{
				func(i int) int {
					return i * i
				},
			},
			{
				func(i int) int {
					return i + i
				},
			},
			{
				func(i int) int {
					return int(math.Pow(2, float64(i)))
				},
			},
		},
	}

	expected := func() []int {
		var slice []int
		for _, t := range(testCase.Funcs) {
			ch := Generator(t)
			for x := range(ch) {
				slice = append(slice, x)
			}
		}
		return slice
	} ()

	result := MergeResult(testCase.Funcs)

	// Transform to general form
	sort.Ints(expected)
	sort.Ints(result)

	assert.Equal(t, expected, result, "The results should be the same")
}
