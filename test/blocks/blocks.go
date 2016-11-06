package blocks

func Double(val int, dblfn func(int) int) int {
	return dblfn(val)
}
