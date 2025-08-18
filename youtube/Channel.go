// SPDX-FileCopyrightText: 2025 Antoni Szymański
// SPDX-License-Identifier: MPL-2.0

package youtube

import (
	"github.com/antoniszymanski/innertube-go/common/shared"
	"github.com/antoniszymanski/innertube-go/common/utils"
	"github.com/dop251/goja"
)

type Channel struct {
	BaseChannel
	// How many videos does this channel have
	VideoCount   string `js:"videoCount"`
	Banner       shared.Thumbnails
	MobileBanner shared.Thumbnails
	TvBanner     shared.Thumbnails
	Shelves      shared.Shelves
}

func (x *Channel) FromObject(vm *goja.Runtime, obj *goja.Object) error {
	if err := vm.ExportTo(obj, x); err != nil {
		return err
	}
	if err := utils.ExportTo(vm, obj, &x.BaseChannel); err != nil {
		return err
	}
	if err := utils.ExportTo(vm, obj.Get("shelves"), &x.Shelves); err != nil {
		return err
	}
	return nil
}
