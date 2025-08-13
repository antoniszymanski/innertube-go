// SPDX-FileCopyrightText: 2025 Antoni Szymański
// SPDX-License-Identifier: MPL-2.0

package innertube

import (
	"time"

	"github.com/mailru/easyjson"
)

//go:generate go tool easyjson video.go

func (c *Client) Video(id string) (*Video, error) {
	data, err := c.call(&requestConfig{
		Method: "POST",
		Path:   "/player",
		Data:   map[string]string{"videoId": id},
	})
	if err != nil {
		return nil, err
	}

	var raw rawVideo
	if err := easyjson.Unmarshal(data, &raw); err != nil {
		return nil, err
	}

	return raw.toVideo()
}

type Video struct {
	ID            string
	Title         string
	LengthSeconds int64
	AuthorID      string
	Description   string
	Thumbnails    []Thumbnail
	ViewCount     int64
	Author        string
	IsPrivate     bool
	IsLiveContent bool
	PublishDate   time.Time
	LikeCount     int64
}

type Thumbnail struct {
	URL    string `json:"url,required"`
	Width  int    `json:"width,required"`
	Height int    `json:"height,required"`
}

func (raw *rawVideo) toVideo() (*Video, error) {
	var v Video
	var err error
	v.ID = raw.VideoDetails.VideoID
	v.Title = raw.VideoDetails.Title
	v.LengthSeconds, err = atoi(raw.VideoDetails.LengthSeconds)
	if err != nil {
		return nil, err
	}
	v.AuthorID = raw.VideoDetails.ChannelID
	v.Description = raw.VideoDetails.ShortDescription
	v.Thumbnails = raw.VideoDetails.Thumbnail.Thumbnails
	v.ViewCount, err = atoi(raw.VideoDetails.ViewCount)
	if err != nil {
		return nil, err
	}
	v.Author = raw.VideoDetails.Author
	v.IsPrivate = raw.VideoDetails.IsPrivate
	v.IsLiveContent = raw.VideoDetails.IsLiveContent
	v.PublishDate = raw.Microformat.PlayerMicroformatRenderer.PublishDate
	v.LikeCount, err = atoi(raw.Microformat.PlayerMicroformatRenderer.LikeCount)
	if err != nil {
		return nil, err
	}
	return &v, nil
}

//easyjson:json
type rawVideo struct {
	VideoDetails struct {
		VideoID          string `json:"videoId,required"`
		Title            string `json:"title,required"`
		LengthSeconds    string `json:"lengthSeconds,required"`
		ChannelID        string `json:"channelId,required"`
		ShortDescription string `json:"shortDescription,required"`
		Thumbnail        struct {
			Thumbnails []Thumbnail `json:"thumbnails,required"`
		} `json:"thumbnail,required"`
		ViewCount     string `json:"viewCount,required"`
		Author        string `json:"author,required"`
		IsPrivate     bool   `json:"isPrivate,required"`
		IsLiveContent bool   `json:"isLiveContent,required"`
	} `json:"videoDetails,required"`
	Microformat struct {
		PlayerMicroformatRenderer struct {
			PublishDate time.Time `json:"publishDate,required"`
			LikeCount   string    `json:"likeCount,required"`
		} `json:"playerMicroformatRenderer,required"`
	} `json:"microformat,required"`
}
