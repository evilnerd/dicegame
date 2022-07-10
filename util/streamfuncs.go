package util

func Map[T any, U any](in []T, convert func(T) U) []U {
	out := make([]U, 0)
	for _, i := range in {
		out = append(out, convert(i))
	}
	return out
}
