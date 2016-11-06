package blocks

func DoWith(val int, dblfn func(int) int) int {
	return dblfn(val)
}
