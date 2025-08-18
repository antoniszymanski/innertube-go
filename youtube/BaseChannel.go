// SPDX-FileCopyrightText: 2025 Antoni Szymański
// SPDX-License-Identifier: MPL-2.0

package youtube

import (
	"github.com/antoniszymanski/innertube-go/common/shared"
	"github.com/antoniszymanski/innertube-go/common/utils"
	"github.com/dop251/goja"
)

type BaseChannel struct {
	ID string `js:"id"`
	// The channel's name
	Name string `js:"name"`
	// The channel's handle start with @
	Handle string `js:"handle"`
	// The channel's description
	Description string `js:"description"`
	// Thumbnails of this Channel
	Thumbnails shared.Thumbnails `js:"thumbnails"`
	// How many subscribers does this channel have.
	//
	// This is not the exact amount, but a literal string like `"1.95M subscribers"`
	SubscriberCount string `js:"subscriberCount"`
	// Continuable of videos
	Videos ChannelVideos
	// Continuable of shorts
	Shorts ChannelShorts
	// Continuable of live
	Live ChannelLive
	// Continuable of playlists
	Playlists ChannelPlaylists
	// Continuable of posts
	Posts ChannelPosts
	// The URL of the channel page
	URL string `js:"url"`
}

type (
	ChannelVideos    = Continuable[VideoCompact]
	ChannelShorts    = Continuable[VideoCompact]
	ChannelLive      = Continuable[VideoCompact]
	ChannelPlaylists = Continuable[PlaylistCompact]
	ChannelPosts     = Continuable[Post]
)

func (x *BaseChannel) FromObject(vm *goja.Runtime, obj *goja.Object) error {
	if err := vm.ExportTo(obj, x); err != nil {
		return err
	}
	if err := utils.ExportTo(vm, obj.Get("videos"), &x.Videos); err != nil {
		return err
	}
	if err := utils.ExportTo(vm, obj.Get("shorts"), &x.Shorts); err != nil {
		return err
	}
	if err := utils.ExportTo(vm, obj.Get("live"), &x.Live); err != nil {
		return err
	}
	if err := utils.ExportTo(vm, obj.Get("playlists"), &x.Playlists); err != nil {
		return err
	}
	if err := utils.ExportTo(vm, obj.Get("posts"), &x.Posts); err != nil {
		return err
	}
	return nil
}
