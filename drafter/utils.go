package drafter

import (
	"errors"
	"math/rand"
	"slices"
)

func shuffle[T any](slice []T) ([]T, error) {
	var out []T
	if len(slice) < 1 {
		return out, errors.New("empty slice to shuffle")
	}
	for i := range rand.Perm(len(slice)) {
		out = append(out, slice[i])
	}
	return out, nil
}

func filter[T any](ss []T, test func(T) bool) []T {
	var out = []T{}
	for _, it := range ss {
		if test(it) {
			out = append(out, it)
		}
	}
	return out
}

func remove[T comparable](ss []T, t T) []T {
	out := slices.Clone(ss)
	return slices.DeleteFunc(out, func(k T) bool { return t == k })
}

func randomOf[T comparable](ss []T) (T, error) {
	var out T
	if !(len(ss) > 0) {
		return out, errors.New("empty slice to pick from")
	}
	k := rand.Intn(len(ss))
	return ss[k], nil
}
