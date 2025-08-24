// SPDX-FileCopyrightText: 2025 Antoni Szymański
// SPDX-License-Identifier: MPL-2.0

package shared

import (
	"github.com/antoniszymanski/innertube-go/common/utils"
	"github.com/dop251/goja"
)

type Shelves []Shelf

func (x *Shelves) FromValue(vm *goja.Runtime, iterable goja.Value) error {
	return (*utils.Array[Shelf])(x).FromValue(vm, iterable)
}

type Shelf struct {
	vm       *goja.Runtime
	Title    string `js:"title"`
	items    []*goja.Object
	Subtitle string `js:"subtitle"`
}

func (x *Shelf) FromObject(vm *goja.Runtime, obj *goja.Object) error {
	if err := vm.ExportTo(obj, x); err != nil {
		return err
	}
	if err := vm.ExportTo(obj.Get("items"), &x.items); err != nil {
		return err
	}
	return nil
}

func ShelfGet[T any](s *Shelf, idx int) (T, error) {
	var target T
	if err := s.vm.ExportTo(s.items[idx], &target); err != nil {
		var zero T
		return zero, err
	}
	return target, nil
}
