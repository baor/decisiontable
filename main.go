package main

import (
	"fmt"
)

// domain is a structure of incoming request to which rules are applied
type domain struct {
	Affiliate   string
	Channel     string
	Origin      string
	Destination string
	Markup      float64
	Price       float64
}

func applyRules(r domain) domain {
	// condition is a flat structure which describes condition columns in decision table
	type condition struct {
		Affiliate     interface{}
		Channel       interface{}
		Origin        interface{}
		Destination   interface{}
		OriginCountry interface{}
		Price         interface{}
	}

	// action is a structure which describes a response in case of match in the table
	type action struct {
		Markup float64
	}

	// fill in decision table with rules
	row(condition{Affiliate: "Aff1", Channel: ANY, Origin: "AMS", Destination: "LON", Price: ANY}, action{Markup: 2.0})
	row(condition{Affiliate: "Aff2", Channel: ANY, Origin: "AMS", Destination: "LON", Price: ANY}, action{})
	row(condition{Affiliate: ANY, Channel: ANY, Origin: "AMS", Destination: ne("LON"), Price: ANY}, action{Markup: 1.0})
	row(condition{Affiliate: ANY, Channel: ANY, Origin: ANY, Destination: "LON", Price: le(11.22)}, action{Markup: 0.5})

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
	r.Markup = act.Markup
	return r
}

func main() {
	r1 := domain{
		Affiliate:   "Aff1",
		Origin:      "AMS",
		Destination: "LON",
	}
	fmt.Println("r1")
	fmt.Printf("Before rule: %+v \n", r1)
	r1 = applyRules(r1)
	fmt.Printf("After rule: %+v \n", r1)

	r2 := domain{
		Affiliate:   "Aff1",
		Origin:      "NONE",
		Destination: "LON",
	}
	fmt.Println("r2")
	fmt.Printf("Before rule: %+v \n", r2)
	r2 = applyRules(r2)
	fmt.Printf("After rule: %+v \n", r2)

	r3 := domain{
		Affiliate:   "Aff1",
		Origin:      "NONE",
		Destination: "LON",
		Price:       11.0,
	}
	fmt.Println("r3")
	fmt.Printf("Before rule: %+v \n", r3)
	r2 = applyRules(r3)
	fmt.Printf("After rule: %+v \n", r3)
}
