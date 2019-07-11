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
			cndFieldName := cndType.Field(i).Name
			cndFieldValueInterface := cndValue.Field(i).Interface()
			cndFieldValue := reflect.ValueOf(cndFieldValueInterface)
			reqFieldValue := reqValue.FieldByName(cndFieldName)
			//fmt.Printf("cndFieldName: %s, cndFieldValue: %+v, reqFieldValue: %+v\n", cndFieldName, cndFieldValue, reqFieldValue)

			// skip zero and ANY
			// TODO: add tests for zero values
			if cndFieldValueInterface == reflect.Zero(cndValue.Field(i).Type()).Interface() || cndFieldValue.Type() == reflect.TypeOf(ANY) {
				continue
			}

			if cndFieldValue.Kind() == reflect.Func {
				cndFieldValueType := cndFieldValue.Type()
				n := cndFieldValueType.NumIn()
				if n != 1 {
					panic("N is " + string(n))
				}

				res := cndFieldValue.Call([]reflect.Value{reflect.ValueOf(reqFieldValue)})
				if len(res) == 0 {
					panic("No results!")
				}
				if res[0].Bool() {
					continue
				}
				match = false
				break
			}

			if eq(cndFieldValue, reqFieldValue) {
				continue
			}
			match = false
			break
		}

		if match {
			return r.action
		}
	}
	return nil
}

func eq(cnd reflect.Value, req reflect.Value) bool {
	if cnd.Kind() != req.Kind() {
		panic("different kinds: " + cnd.Kind().String() + "!=" + req.Kind().String())
	}

	cndKind := reflect.TypeOf(cnd.Interface()).Kind()
	var res bool
	switch cndKind {
	case reflect.Bool:
		res = cnd.Bool() == req.Bool()
	case reflect.Int:
		res = cnd.Int() == req.Int()
	case reflect.String:
		res = cnd.String() == req.String()
	default:
		panic("unsupported kind " + cndKind.String())
	}
	return res
}

func ne(cnd interface{}) interface{} {
	return func(req reflect.Value) bool {
		cndV := reflect.ValueOf(cnd)
		reqV := reflect.ValueOf(req.Interface())
		if cndV.Kind() != reqV.Kind() {
			panic("different types: " + cndV.Kind().String() + "!=" + reqV.Kind().String())
		}

		cndKind := reflect.TypeOf(cndV.Interface()).Kind()
		var res bool
		switch cndKind {
		case reflect.Bool:
			res = cndV.Bool() != reqV.Bool()
		case reflect.Int:
			res = cndV.Int() != reqV.Int()
		case reflect.String:
			res = cndV.String() != reqV.String()
		default:
			panic("unsupported kind " + cndKind.String())
		}

		return res
	}
}

func le(cnd interface{}) interface{} {
	return func(req reflect.Value) bool {
		cndV := reflect.ValueOf(cnd)
		reqV := reflect.ValueOf(req.Interface())
		if cndV.Kind() != reqV.Kind() {
			panic("different types: " + cndV.Kind().String() + "!=" + reqV.Kind().String())
		}

		cndKind := reflect.TypeOf(cndV.Interface()).Kind()
		var res bool
		switch cndKind {
		case reflect.Int:
			res = reqV.Int() <= cndV.Int()
		case reflect.Float32:
			res = reqV.Float() <= cndV.Float()
		case reflect.Float64:
			res = reqV.Float() <= cndV.Float()
		default:
			panic("unsupported kind " + cndKind.String())
		}

		return res
	}
}

func ge(cnd interface{}) interface{} {
	return func(req reflect.Value) bool {
		cndV := reflect.ValueOf(cnd)
		reqV := reflect.ValueOf(req.Interface())
		if cndV.Kind() != reqV.Kind() {
			panic("different types: " + cndV.Kind().String() + "!=" + reqV.Kind().String())
		}

		cndKind := reflect.TypeOf(cndV.Interface()).Kind()
		var res bool
		switch cndKind {
		case reflect.Int:
			res = reqV.Int() >= cndV.Int()
		case reflect.Float32:
			res = reqV.Float() >= cndV.Float()
		case reflect.Float64:
			res = reqV.Float() >= cndV.Float()
		default:
			panic("unsupported kind " + cndKind.String())
		}

		return res
	}
}
