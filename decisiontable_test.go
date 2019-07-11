package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_apply_ne_1(t *testing.T) {
	type condition struct {
		Destination interface{}
	}

	type action struct {
		response bool
	}

	// fill in decision table with rules
	row(condition{Destination: ne("LON")}, action{response: true})

	// Act
	res := apply(condition{Destination: "AMS"})
	assert.NotNil(t, res)
	assert.True(t, res.(action).response)
}

func Test_apply_ne_2(t *testing.T) {
	type condition struct {
		Destination interface{}
	}

	type action struct {
		response bool
	}

	// fill in decision table with rules
	row(condition{Destination: ne("LON")}, action{response: true})

	// Act
	res := apply(condition{Destination: "LON"})
	assert.Nil(t, res)
}

func Test_apply_le_1(t *testing.T) {
	type condition struct {
		Price interface{}
	}

	type action struct {
		Markup float64
	}

	// fill in decision table with rules
	row(condition{Price: le(100.10)}, action{Markup: 1.5})

	// Act
	res := apply(condition{Price: 100.00})
	assert.NotNil(t, res)
	assert.Equal(t, float64(1.5), res.(action).Markup)

	res = apply(condition{Price: 100.20})
	assert.Nil(t, res)
}

func Benchmark_apply_le_1(b *testing.B) {
	type condition struct {
		C1    interface{}
		C2    interface{}
		C3    interface{}
		C4    interface{}
		C5    interface{}
		C6    interface{}
		C7    interface{}
		C8    interface{}
		C9    interface{}
		Price interface{}
	}

	type action struct {
		Markup float64
	}

	for i := 0; i < 50; i++ {
		row(condition{Price: le(100.10)}, action{Markup: 1.5})
	}
	// fill in decision table with rules

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		apply(condition{Price: 200.00})
	}
}
