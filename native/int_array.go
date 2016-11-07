package native

import (
	"fmt"
	"strings"
)

//ruby:nomain
//ruby:module Gorb

type NativeIntArray struct {
	//ruby:ignore
	List *[]int
}

func (a *NativeIntArray) Each(fn func(int)) {
	for _, v := range *a.List {
		fn(v)
	}
}

//ruby []
func (a *NativeIntArray) Get(idx int) int {
	return (*a.List)[idx]
}

//ruby []=
func (a *NativeIntArray) Set(idx int, value int) {
	(*a.List)[idx] = value
}

func (a *NativeIntArray) Push(value int) {
	*a.List = append(*a.List, value)
}

func (a *NativeIntArray) Size() int {
	return len(*a.List)
}

func (a *NativeIntArray) Length() int {
	return a.Size()
}

func (a *NativeIntArray) String() string {
	return a.Inspect()
}

func (a *NativeIntArray) Inspect() string {
	mapped := make([]string, a.Size())
	for i, v := range *a.List {
		mapped[i] = fmt.Sprintf("%d", v)
	}
	return "[" + strings.Join(mapped, ", ") + "]"
}
