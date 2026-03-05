// SPDX-FileCopyrightText: 2025 Antoni Szymański
// SPDX-License-Identifier: MPL-2.0

package youtube

import (
	"github.com/antoniszymanski/innertube-go/internal"
	"github.com/grafana/sobek"
)

type VideoRelated struct {
	VideoCompact    *VideoCompact
	PlaylistCompact *PlaylistCompact
}

func (x *VideoRelated) FromValue(vm *sobek.Runtime, val sobek.Value) error {
	module := vm.Get("youtubei").ToObject(vm)
	switch {
	case vm.InstanceOf(val, module.Get("VideoCompact").ToObject(vm)):
		x.VideoCompact = &VideoCompact{}
		return internal.ExportTo(vm, val, x.VideoCompact)
	case vm.InstanceOf(val, module.Get("PlaylistCompact").ToObject(vm)):
		x.PlaylistCompact = &PlaylistCompact{}
		return internal.ExportTo(vm, val, x.PlaylistCompact)
	default:
		panic("unreachable")
	}
}
