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

func ExportTo(vm *goja.Runtime, in goja.Value, out any) error {
	if in == nil {
		return nil
	}
	switch i := out.(type) {
	case FromValue:
		return i.FromValue(vm, in)
	case FromObject:
		obj, err := ToObject(vm, in)
		if err != nil {
			return err
		}
		return i.FromObject(vm, obj)
	default:
		return vm.ExportTo(in, out)
	}
}

func ToObject(vm *goja.Runtime, val goja.Value) (*goja.Object, error) {
	obj, ex := Try(vm, func() *goja.Object { return val.ToObject(vm) })
	if ex != nil {
		return nil, ex
	}
	return obj, nil
}
