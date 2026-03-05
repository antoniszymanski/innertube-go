// SPDX-FileCopyrightText: 2025 Antoni Szymański
// SPDX-License-Identifier: MPL-2.0

package internal

import (
	"reflect"

	"github.com/grafana/sobek"
)

type ToValue interface {
	ToValue(vm *sobek.Runtime) (sobek.Value, error)
}

func ImportFrom(vm *sobek.Runtime, in any) (sobek.Value, error) {
	if in == nil || try(reflect.ValueOf(in).IsNil) {
		return sobek.Undefined(), nil
	}
	switch i := in.(type) {
	case ToValue:
		return i.ToValue(vm)
	default:
		return vm.ToValue(in), nil
	}
}

func try[T any](f func() T) T {
	defer func() { _ = recover() }()
	return f()
}

func Try[T any](vm *sobek.Runtime, f func() T) (T, *sobek.Exception) {
	var x T
	ex := vm.Try(func() {
		x = f()
	})
	return x, ex
}
