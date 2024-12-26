package fn

func Identity[T any](t T) T {
	return t
}

func Eq[T comparable](eq T) func(t T) bool {
	return func(t T) bool {
		return t == eq
	}
}
