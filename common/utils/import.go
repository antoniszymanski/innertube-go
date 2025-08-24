// SPDX-FileCopyrightText: 2025 Antoni Szymański
// SPDX-License-Identifier: MPL-2.0

package utils

import "github.com/dop251/goja"

type ToValue interface {
	ToValue(vm *goja.Runtime) (goja.Value, error)
}

func ImportFrom(vm *goja.Runtime, in any) (goja.Value, error) {
	switch i := in.(type) {
	case nil:
		return goja.Undefined(), nil
	case ToValue:
		return i.ToValue(vm)
	default:
		return vm.ToValue(in), nil
	}
}
