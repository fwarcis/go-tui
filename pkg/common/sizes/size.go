package sizes

import (
	"errors"
	"fmt"
)

type Size[S ~[]E, E any] struct {
	len    int
	nToCap int
}

func NewSize[S ~[]E, E any](
	len, additionalCap int,
) (Size[S, E], error) {
	var err error

	if len < 0 {
		err = errors.Join(err, newError(
			"NewSize", "%w",
			NewNegativeLenError(len),
		))
	}
	if additionalCap < 0 {
		err = errors.Join(err, newError(
			"NewSize", "%w",
			NewNegativeCapError(additionalCap),
		))
	}

	if err != nil {
		return Size[S, E]{}, err
	}
	return Size[S, E]{
		len:    len,
		nToCap: additionalCap,
	}, nil
}

type NegativeLenError struct {
	Value int
}

func NewNegativeLenError(val int) NegativeLenError {
	return NegativeLenError{Value: val}
}

func (e NegativeLenError) Error() string {
	return fmt.Sprintf(
		"len out of range [%d]", e.Value,
	)
}

type NegativeCapError struct {
	Value int
}

func NewNegativeCapError(val int) NegativeCapError {
	return NegativeCapError{Value: val}
}

func (e NegativeCapError) Error() string {
	return fmt.Sprintf(
		"cap out of range [%d]", e.Value,
	)
}

func (s Size[S, _]) Apply(slice S) S {
	var nToCap int
	if s.len < cap(slice) {
		nToCap = s.nToCap
	} else {
		nToCap = s.len - cap(slice) + s.nToCap
	}
	return append(
		slice[:cap(slice)],
		make(S, nToCap)...,
	)[:s.len]
}

func (s Size[_, _]) Len() int {
	return s.len
}

func (s Size[_, _]) NToCap() int {
	return s.nToCap
}
