// SPDX-FileCopyrightText: 2025 Antoni Szymański
// SPDX-License-Identifier: MPL-2.0

package internal

import (
	"reflect"
	"strings"

	"github.com/dop251/goja"
)

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
	}
	if typ := reflect.TypeOf(out); typ.Kind() == reflect.Pointer {
		typ = typ.Elem()
		out := func() reflect.Value { return reflect.ValueOf(out).Elem() }
		switch {
		case typ.Kind() == reflect.Slice:
			return exportToSlice(vm, in, out())
		case isOption(typ):
			return exportToOption(vm, in, out())
		}
	}
	return vm.ExportTo(in, out)
}

func ToObject(vm *goja.Runtime, val goja.Value) (*goja.Object, error) {
	obj, ex := Try(vm, func() *goja.Object { return val.ToObject(vm) })
	if ex != nil {
		return nil, ex
	}
	return obj, nil
}

func exportToSlice(vm *goja.Runtime, in goja.Value, out reflect.Value) (err error) {
	elemType := out.Type().Elem()
	vm.ForOf(in, func(val goja.Value) bool {
		elemPtr := reflect.New(elemType)
		if err = ExportTo(vm, val, elemPtr.Interface()); err != nil {
			return false
		}
		out = reflect.Append(out, elemPtr.Elem())
		return true
	})
	return
}

func isOption(typ reflect.Type) bool {
	if typ.PkgPath() != "github.com/antoniszymanski/option-go" {
		return false
	}
	name := typ.Name()
	return strings.HasPrefix(name, "Option[") && strings.HasSuffix(name, "]")
}

func exportToOption(vm *goja.Runtime, in goja.Value, out reflect.Value) error {
	if goja.IsNull(in) || goja.IsUndefined(in) {
		out.Set(reflect.Zero(out.Type()))
		return nil
	}
	if err := vm.ExportTo(in, field(out, 2).Interface()); err != nil {
		out.Set(reflect.Zero(out.Type()))
		return err
	}
	field(out, 1).Elem().SetBool(true)
	return nil
}

func field(v reflect.Value, i int) reflect.Value {
	f := v.Field(i)
	return reflect.NewAt(f.Type(), f.Addr().UnsafePointer())
}
