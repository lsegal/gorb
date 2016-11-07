package native

import (
	"fmt"
	"strings"
)

//ruby:nomain
//ruby:module Gorb

type NativeFloatArray struct {
	//ruby:ignore
	List *[]float64
}

func (a *NativeFloatArray) Each(fn func(float64)) {
	for _, v := range *a.List {
		fn(v)
	}
}

//ruby []
func (a *NativeFloatArray) Get(idx int) float64 {
	return (*a.List)[idx]
}

//ruby []=
func (a *NativeFloatArray) Set(idx int, value float64) {
	(*a.List)[idx] = value
}

func (a *NativeFloatArray) Push(value float64) {
	*a.List = append(*a.List, value)
}

func (a *NativeFloatArray) Size() int {
	return len(*a.List)
}

func (a *NativeFloatArray) Length() int {
	return a.Size()
}

func (a *NativeFloatArray) String() string {
	return a.Inspect()
}

func (a *NativeFloatArray) Inspect() string {
	mapped := make([]string, a.Size())
	for i, v := range *a.List {
		mapped[i] = fmt.Sprintf("%f", v)
	}
	return "[" + strings.Join(mapped, ", ") + "]"
}
