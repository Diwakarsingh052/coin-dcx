package sum

func SumInt(vs []int) int {

	sum := 0
	if vs == nil {
		return 0
	}
	for _, v := range vs {
		sum = v + sum
	}
	return sum

}

// go test -run SumInt/one -v // it is a pattern matching // all tests which have the prefix will run
