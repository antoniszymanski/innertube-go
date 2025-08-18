// SPDX-FileCopyrightText: 2025 Antoni Szymański
// SPDX-License-Identifier: MPL-2.0

package shared

import (
	"strings"

	"github.com/antoniszymanski/innertube-go/common/utils"
	"github.com/dop251/goja"
)

type Thumbnails []Thumbnail

func (x *Thumbnails) FromObject(vm *goja.Runtime, obj *goja.Object) error {
	return (*utils.Array[Thumbnail])(x).FromObject(vm, obj)
}

type Thumbnail struct {
	URL    string `js:"url"`
	Width  int64  `js:"width"`
	Height int64  `js:"height"`
}

func (x *Thumbnail) FromValue(vm *goja.Runtime, val goja.Value) error {
	if err := vm.ExportTo(val, x); err != nil {
		return err
	}
	switch {
	case strings.HasPrefix(x.URL, "//"):
		x.URL = "https:" + x.URL
	case !strings.HasPrefix(x.URL, "https://"):
		x.URL = "https://" + x.URL
	}
	return nil
}
