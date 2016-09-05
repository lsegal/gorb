package data

import "fmt"

type RGB struct {
	R, G, B int
}

type HSV struct {
	H, S, V float64
}

func (h *HSV) String() string {
	return fmt.Sprintf("HSV(h=%f, s=%f, v=%f)", h.H, h.S, h.V)
}

func (h *HSV) Inspect() string {
	return h.String()
}
