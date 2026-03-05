// SPDX-FileCopyrightText: 2025 Antoni Szymański
// SPDX-License-Identifier: MPL-2.0

package youtube

import (
	"errors"

	"github.com/antoniszymanski/innertube-go/internal"
	"github.com/antoniszymanski/innertube-go/modules/youtubei"
	"github.com/antoniszymanski/innertube-go/shared"
	"github.com/grafana/sobek"
)

type Client struct {
	vm   *sobek.Runtime
	this *sobek.Object
}

func NewClient(options *shared.ClientOptions) (Client, error) {
	vm := internal.NewVM()
	youtubei.Enable(vm)

	construct, ex := internal.Try(vm, func() sobek.Value {
		return vm.Get("youtubei").ToObject(vm).Get("Client")
	})
	if ex != nil {
		return Client{}, ex
	}

	arg, err := internal.ImportFrom(vm, options)
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
	if err := internal.ExportTo(c.vm, val, &result); err != nil {
		return shared.OAuthProps{}, err
	}
	return result, nil
}

func (c Client) GetVideo(id string) (VideoResult, error) {
	val, err := internal.CallAsync(c.vm, c.this, "getVideo", c.vm.ToValue(id))
	if err != nil {
		return VideoResult{}, err
	}
	if sobek.IsUndefined(val) {
		return VideoResult{}, ErrVideoNotFound
	}

	var result VideoResult
	if err := internal.ExportTo(c.vm, val, &result); err != nil {
		return VideoResult{}, err
	}
	return result, nil
}

var ErrVideoNotFound = errors.New("video not found")

type VideoResult struct {
	Video     *Video
	LiveVideo *LiveVideo
}

func (x *VideoResult) FromValue(vm *sobek.Runtime, val sobek.Value) error {
	module := vm.Get("youtubei").ToObject(vm)
	switch {
	case vm.InstanceOf(val, module.Get("Video").ToObject(vm)):
		x.Video = &Video{}
		return internal.ExportTo(vm, val, x.Video)
	case vm.InstanceOf(val, module.Get("LiveVideo").ToObject(vm)):
		x.LiveVideo = &LiveVideo{}
		return internal.ExportTo(vm, val, x.LiveVideo)
	default:
		panic("unreachable")
	}
}

func (c Client) GetChannel(id string) (*Channel, error) {
	val, err := internal.CallAsync(c.vm, c.this, "getChannel", c.vm.ToValue(id))
	if err != nil {
		return nil, err
	}
	if sobek.IsUndefined(val) {
		return nil, ErrChannelNotFound
	}

	var channel Channel
	if err := internal.ExportTo(c.vm, val, &channel); err != nil {
		return nil, err
	}
	return &channel, nil
}

var ErrChannelNotFound = errors.New("channel not found")
