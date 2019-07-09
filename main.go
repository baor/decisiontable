package main

import (
	"fmt"
)

// domain is a structure of incoming request to which rules are applied
type domain struct {
	affiliate   string
	channel     string
	origin      string
	destination string
	markup      float64
}

func applyRules(r domain) domain {
	// condition is a flat structure which describes condition columns in decision table
	type condition struct {
		affiliate     interface{}
		channel       interface{}
		origin        interface{}
		destination   interface{}
		originCountry interface{}
		price         interface{}
	}

	// action is a structure which describes a response in case of match in the table
	type action struct {
		markup float64
	}

	// fill in decision table with rules
	row(condition{affiliate: "Aff1", channel: ANY, origin: "AMS", destination: "LON"}, action{markup: 2.0})
	row(condition{affiliate: "Aff2", channel: ANY, origin: "AMS", destination: "LON"}, action{})
	row(condition{affiliate: ANY, channel: ANY, origin: "AMS", destination: "PAR"}, action{markup: 1.0})
	row(condition{affiliate: ANY, channel: ANY, origin: "AMS", destination: "LON", price: 11.22}, action{markup: 1.0})
	row(condition{affiliate: ANY, channel: ANY, origin: "AMS", destination: ne("LON"), price: 11.22}, action{markup: 1.0})

	// apply decision table to the request
	res := apply(r)

	// if respose is nil, no actions were applied
	if res == nil {
		fmt.Println("action is nil")
		return r
	}

	// if response is not nil, match happened
	act := res.(action)
	fmt.Printf("action: %+v \n", act)

	// assemble action data to response
	r.markup = act.markup
	return r
}

func main() {
	r1 := domain{
		affiliate:   "Aff1",
		origin:      "AMS",
		destination: "LON",
	}
	fmt.Println("r1")
	fmt.Printf("Before rule: %+v \n", r1)
	r1 = applyRules(r1)
	fmt.Printf("After rule: %+v \n", r1)

	r2 := domain{
		affiliate:   "Aff1",
		origin:      "NONE",
		destination: "LON",
	}
	fmt.Println("r2")
	fmt.Printf("Before rule: %+v \n", r2)
	r2 = applyRules(r2)
	fmt.Printf("After rule: %+v \n", r2)
}
