package parser

func NoErr(err error) {
	if err != nil {
		panic(err)
	}
}

func Must[T any](v T, err error) T {
	NoErr(err)
	return v
}

func ReverseSlice(b []byte) []byte {
	var new_b = make([]byte, len(b))
	for i, v := range b {
		new_b[len(b)-i-1] = v
	}

	return new_b
}
