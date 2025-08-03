// SPDX-FileCopyrightText: 2025 Antoni Szymański
// SPDX-License-Identifier: MPL-2.0

package innertube

import "time"

func (c *Client) GetVideo(id string) (*Video, error) {
	var raw rawVideo
	if err := c.call(&requestConfig{
		Method: "POST",
		Path:   "/player",
		Input:  map[string]string{"videoId": id},
		Output: &raw,
	}); err != nil {
		return nil, err
	}
	return raw.toVideo()
}

type Video struct {
	ID            string      `json:"id"`
	Title         string      `json:"title"`
	LengthSeconds int64       `json:"lengthSeconds"`
	ChannelID     string      `json:"authorId"`
	Description   string      `json:"description"`
	Thumbnails    []Thumbnail `json:"thumbnails"`
	ViewCount     int64       `json:"viewCount"`
	Author        string      `json:"author"`
	IsPrivate     bool        `json:"isPrivate"`
	IsLiveContent bool        `json:"isLiveContent"`
	PublishDate   time.Time   `json:"publishDate"`
	LikeCount     int64       `json:"likeCount"`
}

type Thumbnail struct {
	URL    string `json:"url"`
	Width  int    `json:"width"`
	Height int    `json:"height"`
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
	v.ChannelID = raw.VideoDetails.ChannelID
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

type rawVideo struct {
	VideoDetails struct {
		VideoID          string `json:"videoId"`
		Title            string `json:"title"`
		LengthSeconds    string `json:"lengthSeconds"`
		ChannelID        string `json:"channelId"`
		ShortDescription string `json:"shortDescription"`
		Thumbnail        struct {
			Thumbnails []Thumbnail `json:"thumbnails"`
		} `json:"thumbnail"`
		ViewCount     string `json:"viewCount"`
		Author        string `json:"author"`
		IsPrivate     bool   `json:"isPrivate"`
		IsLiveContent bool   `json:"isLiveContent"`
	} `json:"videoDetails"`
	Microformat struct {
		PlayerMicroformatRenderer struct {
			PublishDate time.Time `json:"publishDate"`
			LikeCount   string    `json:"likeCount"`
		} `json:"playerMicroformatRenderer"`
	} `json:"microformat"`
}
