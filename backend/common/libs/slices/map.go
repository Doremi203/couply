package slices

func Map[From any, To any](from []From, f func(From) To) []To {
	ret := make([]To, 0, len(from))
	for _, item := range from {
		ret = append(ret, f(item))
	}

	return ret
}
