// SPDX-FileCopyrightText: 2025 Antoni Szymański
// SPDX-License-Identifier: MPL-2.0

package youtube

import (
	"github.com/antoniszymanski/innertube-go/internal"
	"github.com/antoniszymanski/innertube-go/shared"
	"github.com/grafana/sobek"
)

type Channel struct {
	BaseChannel
	// How many videos does this channel have
	VideoCount   string `js:"videoCount"`
	Banner       shared.Thumbnails
	MobileBanner shared.Thumbnails
	TvBanner     shared.Thumbnails
	Shelves      []shared.Shelf
}

func (x *Channel) FromObject(vm *sobek.Runtime, obj *sobek.Object) error {
	if err := vm.ExportTo(obj, x); err != nil {
		return err
	}
	if err := internal.ExportTo(vm, obj, &x.BaseChannel); err != nil {
		return err
	}
	if err := internal.ExportTo(vm, obj.Get("shelves"), &x.Shelves); err != nil {
		return err
	}
	return nil
}
