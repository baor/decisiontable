package main

import (
	"reflect"
)

type decision struct {
	condition interface{}
	action    interface{}
}

var tbl []decision

// ANYtype - "- Any -"
type ANYtype struct{}

// ANY - "- Any -"
var ANY = ANYtype{}

func init() {
	tbl = []decision{}
}

func row(condition interface{}, action interface{}) {
	tbl = append(tbl, decision{
		condition: condition,
		action:    action,
	})
	return
}

func apply(req interface{}) interface{} {
	if len(tbl) == 0 {
		panic("table is empty")
	}

	//reqType := reflect.TypeOf(req)
	reqValue := reflect.ValueOf(req)

	for _, r := range tbl {
		match := true

		cnd := r.condition
		cndType := reflect.TypeOf(cnd)
		cndValue := reflect.ValueOf(cnd)
		cndNumField := cndType.NumField()

		for i := 0; i < cndNumField; i++ {
			if cndValue.Field(i).Type() == reflect.TypeOf(ANY) {
				continue
			}

			cndFieldName := cndType.Field(i).Name
			cndFieldValue := cndValue.Field(i)
			reqFieldValue := reqValue.FieldByName(cndFieldName)

			if cndFieldValue.Kind() == reflect.Func {
				res := cndFieldValue.Call([]reflect.Value{reqFieldValue})
				if len(res) == 0 {
					panic("No results!")
				}
				if res[0].Bool() {
					continue
				}
			}

			if eq(cndFieldValue, reqFieldValue) {
				continue
			}
			match = false
		}

		if match {
			return r.action
		}
	}
	return nil
}

func eq(cnd reflect.Value, req reflect.Value) bool {
	if cnd.Kind() != req.Kind() {
		panic("different types")
	}

	switch cnd.Kind() {
	case reflect.Bool:
		return cnd.Bool() == req.Bool()
	case reflect.Int:
		return cnd.Int() == req.Int()
	case reflect.String:
		return cnd.String() == req.String()

	default:
		panic("unsupported type")
	}
}

func ne(cnd interface{}) interface{} {
	return func(cnd reflect.Value, req reflect.Value) bool {
		if cnd.Kind() != req.Kind() {
			panic("different types")
		}

		switch cnd.Kind() {
		case reflect.Bool:
			return cnd.Bool() != req.Bool()
		case reflect.Int:
			return cnd.Int() != req.Int()
		case reflect.String:
			return cnd.String() != req.String()

		default:
			panic("unsupported type")
		}
	}
}
