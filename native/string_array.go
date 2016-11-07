package native

import (
	"fmt"
	"strings"
)

//ruby:nomain
//ruby:module Gorb

type NativeStringArray struct {
	//ruby:ignore
	List *[]string
}

func (a *NativeStringArray) Each(fn func(string)) {
	for _, v := range *a.List {
		fn(v)
	}
}

//ruby []
func (a *NativeStringArray) Get(idx int) string {
	return (*a.List)[idx]
}

//ruby []=
func (a *NativeStringArray) Set(idx int, value string) {
	(*a.List)[idx] = value
}

func (a *NativeStringArray) Push(value string) {
	*a.List = append(*a.List, value)
}

func (a *NativeStringArray) Size() int {
	return len(*a.List)
}

func (a *NativeStringArray) Length() int {
	return a.Size()
}

func (a *NativeStringArray) String() string {
	return a.Inspect()
}

func (a *NativeStringArray) Inspect() string {
	mapped := make([]string, a.Size())
	for i, v := range *a.List {
		mapped[i] = fmt.Sprintf("%q", v)
	}
	return "[" + strings.Join(mapped, ", ") + "]"
}
