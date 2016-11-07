package array

import "fmt"

func ReverseArray(list []string) *[]string {
	out := make([]string, len(list))
	for i, v := range list {
		out[len(list)-1-i] = v
	}
	return &out
}

func MutateArray(list []string, muter func(*[]string)) {
	muter(&list)
	fmt.Println("Go has", list)
}

func MutateIntArray(list []int, muter func(*[]int)) {
	muter(&list)
	fmt.Println("Go has", list)
}
