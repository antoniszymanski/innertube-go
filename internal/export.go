// SPDX-FileCopyrightText: 2025 Antoni Szymański
// SPDX-License-Identifier: MPL-2.0

package internal

import (
	"reflect"

	. "github.com/antoniszymanski/option-go"
	"github.com/grafana/sobek"
)

type FromValue interface {
	FromValue(vm *sobek.Runtime, val sobek.Value) error
}

type FromObject interface {
	FromObject(vm *sobek.Runtime, obj *sobek.Object) error
}

func ExportTo(vm *sobek.Runtime, in sobek.Value, out any) error {
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
		case IsOption(typ):
			return exportToOption(vm, in, out())
		}
	}
	return vm.ExportTo(in, out)
}

func ToObject(vm *sobek.Runtime, val sobek.Value) (*sobek.Object, error) {
	obj, ex := Try(vm, func() *sobek.Object { return val.ToObject(vm) })
	if ex != nil {
		return nil, ex
	}
	return obj, nil
}

func exportToSlice(vm *sobek.Runtime, in sobek.Value, out reflect.Value) (err error) {
	elemType := out.Type().Elem()
	vm.ForOf(in, func(val sobek.Value) bool {
		elemPtr := reflect.New(elemType)
		if err = ExportTo(vm, val, elemPtr.Interface()); err != nil {
			return false
		}
		out = reflect.Append(out, elemPtr.Elem())
		return true
	})
	return
}

func exportToOption(vm *sobek.Runtime, in sobek.Value, out reflect.Value) error {
	if sobek.IsNull(in) || sobek.IsUndefined(in) {
		out.Set(reflect.Zero(out.Type()))
		return nil
	}
	if err := vm.ExportTo(in, field(out, "value").Interface()); err != nil {
		out.Set(reflect.Zero(out.Type()))
		return err
	}
	field(out, "valid").Elem().SetBool(true)
	return nil
}

func field(v reflect.Value, name string) reflect.Value {
	f := v.FieldByName(name)
	return reflect.NewAt(f.Type(), f.Addr().UnsafePointer())
}
