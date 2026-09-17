package asrts

func Equal[T comparable](left, right T) {
	if left != right {
		panic("")
	}
}

func NotEqual[T comparable](left, right T) {
	if left == right {
		panic("")
	}
}

func Must[V any](val V, err error) V {
	if err != nil {
		panic(err)
	}
	return val
}

func ReturnIfNil[V any, I any](
	val V, indicator I,
) (V, I) {
	if any(indicator) != nil {
		panic("")
	}

	return val, indicator
}

func MustBeNil[V any](
	val V, indicator any,
) V {
	if indicator != nil {
		panic("")
	}

	return val
}

func ReturnIfTrue[V any](
	val V, ok bool,
) (V, bool) {
	if !ok {
		panic("")
	}

	return val, ok
}

func MustBeTrue[V any](
	val V, ok bool,
) V {
	if !ok {
		panic("")
	}

	return val
}
