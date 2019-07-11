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
		markup float64
	}

	// fill in decision table with rules
	row(condition{Price: le(100.10)}, action{markup: 1.5})

	// Act
	res := apply(condition{Price: 100.00})
	assert.NotNil(t, res)
	assert.Equal(t, float64(1.5), res.(action).markup)

	res = apply(condition{Price: 100.20})
	assert.Nil(t, res)
}

func Benchmark_apply_le_1(b *testing.B) {
	type condition struct {
		Affiliate interface{}
		Channel   interface{}
		Price     interface{}
	}

	type action struct {
		markup float64
	}

	for i := 0; i < 50; i++ {
		row(condition{Affiliate: ANY, Channel: ANY, Price: le(100.10)}, action{markup: 1.5})
	}
	// fill in decision table with rules

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		apply(condition{Price: 200.00})
	}
}
