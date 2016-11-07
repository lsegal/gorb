package array

func ReverseArray(list []string) []string {
	out := make([]string, len(list))
	for i, v := range list {
		out[len(list)-1-i] = v
	}
	return out
}
