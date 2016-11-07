package native

import (
	"fmt"
	"strings"
)

//ruby:nomain
//ruby:module Gorb

type NativeBoolArray struct {
	//ruby:ignore
	List *[]bool
}

func (a *NativeBoolArray) Each(fn func(bool)) {
	for _, v := range *a.List {
		fn(v)
	}
}

//ruby []
func (a *NativeBoolArray) Get(idx int) bool {
	return (*a.List)[idx]
}

//ruby []=
func (a *NativeBoolArray) Set(idx int, value bool) {
	(*a.List)[idx] = value
}

func (a *NativeBoolArray) Push(value bool) {
	*a.List = append(*a.List, value)
}

func (a *NativeBoolArray) Size() int {
	return len(*a.List)
}

func (a *NativeBoolArray) Length() int {
	return a.Size()
}

func (a *NativeBoolArray) String() string {
	return a.Inspect()
}

func (a *NativeBoolArray) Inspect() string {
	mapped := make([]string, a.Size())
	for i, v := range *a.List {
		mapped[i] = fmt.Sprintf("%t", v)
	}
	return "[" + strings.Join(mapped, ", ") + "]"
}
