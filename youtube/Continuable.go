// SPDX-FileCopyrightText: 2025 Antoni Szymański
// SPDX-License-Identifier: MPL-2.0

package youtube

import (
	"github.com/antoniszymanski/innertube-go/common/utils"
	"github.com/dop251/goja"
)

type Continuable[T any] struct {
	vm   *goja.Runtime
	this *goja.Object
}

func (x *Continuable[T]) Items() ([]T, error) {
	var result utils.Array[T]
	if err := utils.ExportTo(x.vm, x.this.Get("items"), &result); err != nil {
		return nil, err
	}
	return result, nil
}

func (x *Continuable[T]) Next(count int64) ([]T, error) {
	val, err := utils.CallAsync(x.vm, x.this, "next", x.vm.ToValue(count))
	if err != nil {
		return nil, err
	}
	var result utils.Array[T]
	if err = utils.ExportTo(x.vm, val, &result); err != nil {
		return nil, err
	}
	return result, nil
}

func (x *Continuable[T]) FromObject(vm *goja.Runtime, obj *goja.Object) error {
	x.vm = vm
	x.this = obj
	return nil
}
