package main

import (
	"reflect"
)

type conditionField struct {
	value  reflect.Value
	isZero bool
	isAny  bool
}

type decision struct {
	condition map[string]conditionField
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

func row(cnd interface{}, action interface{}) {
	cndType := reflect.TypeOf(cnd)
	cndValue := reflect.ValueOf(cnd)
	cndNumField := cndType.NumField()

	// parse condition
	condition := map[string]conditionField{}
	for i := 0; i < cndNumField; i++ {
		cndFieldValueInterface := cndValue.Field(i).Interface()

		name := cndType.Field(i).Name
		cndField := conditionField{
			value: reflect.ValueOf(cndFieldValueInterface),
		}

		//fmt.Printf("cndFieldName: %s, cndFieldValue: %+v, reqFieldValue: %+v\n", cndFieldName, cndFieldValue, reqFieldValue)

		// skip zero and ANY
		// TODO: add tests for zero values

		cndField.isZero = cndFieldValueInterface == reflect.Zero(cndValue.Field(i).Type()).Interface()

		if !cndField.isZero {
			cndField.isAny = cndField.value.Type() == reflect.TypeOf(ANY)
		}

		condition[name] = cndField
	}

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
		for name, cndField := range r.condition {
			// skip zero and ANY
			// TODO: add tests for zero values
			if cndField.isZero || cndField.isAny {
				continue
			}

			reqFieldValue := reqValue.FieldByName(name)

			if cndField.value.Kind() == reflect.Func {
				// get number of arguments of the function
				n := cndField.value.Type().NumIn()
				if n != 1 {
					panic("N is " + string(n))
				}

				res := cndField.value.Call([]reflect.Value{reflect.ValueOf(reqFieldValue)})
				if len(res) == 0 {
					panic("No results!")
				}
				if res[0].Bool() {
					continue
				}
				match = false
				break
			}

			if eq(cndField.value, reqFieldValue) {
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
