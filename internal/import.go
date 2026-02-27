// SPDX-FileCopyrightText: 2025 Antoni Szymański
// SPDX-License-Identifier: MPL-2.0

package internal

import (
	"reflect"

	"github.com/dop251/goja"
)

type ToValue interface {
	ToValue(vm *goja.Runtime) (goja.Value, error)
}

func ImportFrom(vm *goja.Runtime, in any) (goja.Value, error) {
	if in == nil || try(reflect.ValueOf(in).IsNil) {
		return goja.Undefined(), nil
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

func Try[T any](vm *goja.Runtime, f func() T) (T, *goja.Exception) {
	var x T
	ex := vm.Try(func() {
		x = f()
	})
	return x, ex
}
