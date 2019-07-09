/*
* it is an exmaple of applying rules in go
* pros:
* 1. testing
* 2. CI/CD development process
* 3. best practicies from software development are applied
* 4. shared work/ visibility with dev
* Cons:
* 1. Visibility with business
* 2. table visual shift
 */
package main

import (
	"fmt"
)

// domain is a structure of incoming request to which rules are applied
type domain struct {
	affiliate     string
	channel       string
	origin        string
	destination   string
	baggageMarkup float64
}

func applyBaggageRules(r domain) domain {
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
		isBaggageEnabled bool
		baggageMarkup    float64
	}

	// fill in decision table with rules
	row(condition{affiliate: "CheapticketsNL", channel: ANY, origin: "AMS", destination: "LON"}, action{isBaggageEnabled: true, baggageMarkup: 2.0})
	row(condition{affiliate: "BudgetAirNL", channel: ANY, origin: "AMS", destination: "LON"}, action{isBaggageEnabled: false})
	row(condition{affiliate: ANY, channel: ANY, origin: "AMS", destination: "PAR"}, action{isBaggageEnabled: true, baggageMarkup: 1.0})
	row(condition{affiliate: ANY, channel: ANY, origin: "AMS", destination: "LON", price: 11.22}, action{isBaggageEnabled: true, baggageMarkup: 1.0})
	row(condition{affiliate: ANY, channel: ANY, origin: "AMS", destination: ne("LON"), price: 11.22}, action{isBaggageEnabled: true, baggageMarkup: 1.0})

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
	r.baggageMarkup = act.baggageMarkup
	return r
}

func main() {
	r1 := domain{
		affiliate:   "Vayama",
		origin:      "AMS",
		destination: "LON",
	}
	fmt.Println("r1")
	fmt.Printf("Before rule: %+v \n", r1)
	r1 = applyBaggageRules(r1)
	fmt.Printf("After rule: %+v \n", r1)

	r2 := domain{
		affiliate:   "Vayama",
		origin:      "NONE",
		destination: "LON",
	}
	fmt.Println("r2")
	fmt.Printf("Before rule: %+v \n", r2)
	r2 = applyBaggageRules(r2)
	fmt.Printf("After rule: %+v \n", r2)
}
