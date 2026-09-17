package asrts

var CanPanic = false

func MayReturnIfNotNil[V, I any](
	val V, indicator I,
) (V, I) {
	if CanPanic && any(indicator) != nil {
		panic("")
	}

	return val, indicator
}

func MayBeNil[V, I any](
	val V, indicator I,
) V {
	if CanPanic && any(indicator) != nil {
		panic("")
	}

	return val
}

func MayReturnIfFalse[V any](
	val V, ok bool,
) (V, bool) {
	if CanPanic && !ok {
		panic("")
	}

	return val, ok
}

func MayBeFalse[V any](
	val V, ok bool,
) V {
	if CanPanic && !ok {
		panic("")
	}

	return val
}
