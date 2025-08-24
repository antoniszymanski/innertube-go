// SPDX-FileCopyrightText: 2025 Antoni Szymański
// SPDX-License-Identifier: MPL-2.0

package youtube

import (
	"fmt"

	"github.com/antoniszymanski/innertube-go/common/shared"
	"github.com/antoniszymanski/innertube-go/common/utils"
	"github.com/antoniszymanski/innertube-go/internal"
	"github.com/antoniszymanski/innertube-go/modules/youtubei"
	"github.com/dop251/goja"
)

type Client struct {
	vm   *goja.Runtime
	this *goja.Object
}

func NewClient(options *shared.ClientOptions) (Client, error) {
	vm := internal.NewVM()
	youtubei.Enable(vm)

	var construct goja.Value
	ex := vm.Try(func() {
		construct = vm.Get("youtubei").ToObject(vm).Get("Client")
	})
	if ex != nil {
		return Client{}, ex
	}

	arg, err := utils.ImportFrom(vm, options)
	if err != nil {
		return Client{}, err
	}
	client, err := vm.New(construct, arg)
	if err != nil {
		return Client{}, err
	}
	return Client{vm: vm, this: client}, nil
}

func (c Client) OAuth() (shared.OAuthProps, error) {
	val := c.this.Get("oauth")
	var result shared.OAuthProps
	if err := utils.ExportTo(c.vm, val, &result); err != nil {
		return shared.OAuthProps{}, err
	}
	return result, nil
}

func (c Client) GetVideo(id string) (VideoResult, error) {
	val, err := utils.CallAsync(c.vm, c.this, "getVideo", c.vm.ToValue(id))
	if err != nil {
		return VideoResult{}, err
	}
	if goja.IsUndefined(val) {
		return VideoResult{}, ErrChannelNotFound{id}
	}

	var result VideoResult
	if err := utils.ExportTo(c.vm, val, &result); err != nil {
		return VideoResult{}, err
	}
	return result, nil
}

type ErrVideoNotFound struct {
	VideoID string
}

func (e ErrVideoNotFound) Error() string {
	return fmt.Sprintf("video %q not found", e.VideoID)
}

type VideoResult struct {
	Video     *Video
	LiveVideo *LiveVideo
}

func (x *VideoResult) FromValue(vm *goja.Runtime, val goja.Value) error {
	module := vm.Get("youtubei").ToObject(vm)
	switch {
	case vm.InstanceOf(val, module.Get("Video").ToObject(vm)):
		x.Video = &Video{}
		return utils.ExportTo(vm, val, x.Video)
	case vm.InstanceOf(val, module.Get("LiveVideo").ToObject(vm)):
		x.LiveVideo = &LiveVideo{}
		return utils.ExportTo(vm, val, x.LiveVideo)
	default:
		panic("unreachable")
	}
}

func (c Client) GetChannel(id string) (*Channel, error) {
	val, err := utils.CallAsync(c.vm, c.this, "getChannel", c.vm.ToValue(id))
	if err != nil {
		return nil, err
	}
	if goja.IsUndefined(val) {
		return nil, ErrChannelNotFound{id}
	}

	var channel Channel
	if err := utils.ExportTo(c.vm, val, &channel); err != nil {
		return nil, err
	}
	return &channel, nil
}

type ErrChannelNotFound struct {
	ChannelID string
}

func (e ErrChannelNotFound) Error() string {
	return fmt.Sprintf("channel %q not found", e.ChannelID)
}
