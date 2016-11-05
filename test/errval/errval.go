package errval

import "errors"

func Flip(n int) (int, error) {
	if n == 0 {
		return 0, errors.New("oops!")
	}
	return 10 - n, nil
}
