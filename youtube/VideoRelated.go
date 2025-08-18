// SPDX-FileCopyrightText: 2025 Antoni Szymański
// SPDX-License-Identifier: MPL-2.0

package youtube

import (
	"github.com/antoniszymanski/innertube-go/common/utils"
	"github.com/dop251/goja"
)

type VideoRelated struct {
	VideoCompact    *VideoCompact
	PlaylistCompact *PlaylistCompact
}

func (x *VideoRelated) FromValue(vm *goja.Runtime, val goja.Value) error {
	module := vm.Get("youtubei").ToObject(vm)
	switch {
	case vm.InstanceOf(val, module.Get("VideoCompact").ToObject(vm)):
		x.VideoCompact = &VideoCompact{}
		return utils.ExportTo(vm, val, x.VideoCompact)
	case vm.InstanceOf(val, module.Get("PlaylistCompact").ToObject(vm)):
		x.PlaylistCompact = &PlaylistCompact{}
		return utils.ExportTo(vm, val, x.PlaylistCompact)
	default:
		panic("unreachable")
	}
}
