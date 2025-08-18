// SPDX-FileCopyrightText: 2025 Antoni Szymański
// SPDX-License-Identifier: MPL-2.0

package utils

import "github.com/dop251/goja"

type FromValue interface {
	FromValue(vm *goja.Runtime, val goja.Value) error
}

type FromObject interface {
	FromObject(vm *goja.Runtime, obj *goja.Object) error
}

func ExportTo(vm *goja.Runtime, val goja.Value, target any) error {
	if val == nil {
		return nil
	}
	switch i := target.(type) {
	case FromValue:
		return i.FromValue(vm, val)
	case FromObject:
		obj, err := ToObject(vm, val)
		if err != nil {
			return err
		}
		return i.FromObject(vm, obj)
	default:
		return vm.ExportTo(val, target)
	}
}

func ToObject(vm *goja.Runtime, val goja.Value) (*goja.Object, error) {
	var obj *goja.Object
	if ex := vm.Try(func() { obj = val.ToObject(vm) }); ex != nil {
		return nil, ex
	}
	return obj, nil
}
