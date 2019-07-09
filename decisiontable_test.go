package main

import (
	"reflect"
	"testing"
)

// func Test_row(t *testing.T) {
// 	type args struct {
// 		condition interface{}
// 		action    interface{}
// 	}
// 	tests := []struct {
// 		name string
// 		args args
// 	}{
// 		// TODO: Add test cases.
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			row(tt.args.condition, tt.args.action)
// 		})
// 	}
// }

func Test_apply(t *testing.T) {
	type args struct {
		req interface{}
	}
	tests := []struct {
		name string
		args args
		want interface{}
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := apply(tt.args.req); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("apply() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_eq(t *testing.T) {
	type args struct {
		cnd reflect.Value
		req reflect.Value
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := eq(tt.args.cnd, tt.args.req); got != tt.want {
				t.Errorf("eq() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_ne(t *testing.T) {
	type args struct {
		cnd interface{}
	}
	tests := []struct {
		name string
		args args
		want interface{}
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ne(tt.args.cnd); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ne() = %v, want %v", got, tt.want)
			}
		})
	}
}
