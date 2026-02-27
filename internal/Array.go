// SPDX-FileCopyrightText: 2025 Antoni Szymański
// SPDX-License-Identifier: MPL-2.0

package internal

import "github.com/dop251/goja"

type Array[T any] []T

func (x *Array[T]) FromValue(vm *goja.Runtime, iterable goja.Value) error {
	var err error
	vm.ForOf(iterable, func(val goja.Value) bool {
		var elem T
		if err = ExportTo(vm, val, &elem); err != nil {
			return false
		}
		*x = append(*x, elem)
		return true
	})
	return err
}
