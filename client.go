// SPDX-FileCopyrightText: 2025 Antoni Szymański
// SPDX-License-Identifier: MPL-2.0

package innertube

import (
	"bytes"
	"io"
	"net/http"
	"net/url"

	"github.com/antoniszymanski/innertube-go/internal"
	"github.com/go-json-experiment/json"
)

type Client struct {
	Language string // default: en
	Region   string // default: US

	UserAgent  string
	HTTPClient *http.Client
}

type requestConfig struct {
	Method string
	Path   string
	Params url.Values
	Input  map[string]string
	Output any
}

func (c *Client) call(config *requestConfig) error {
	if config.Params == nil {
		config.Params = make(url.Values)
	}
	config.Params.Set("key", "AIzaSyAO_FJ2SlqU8Q4STEHLGCilw_Y9_11qcW8")
	config.Params.Set("prettyPrint", "false")
	query := config.Params.Encode()

	var body io.Reader
	if config.Input != nil {
		payload := internal.Payload{
			Context: internal.Context{
				Client: internal.Client{
					ClientName:    "WEB",
					ClientVersion: "2.20201209.01.00",
					Hl:            "en",
					Gl:            "US",
				},
			},
			Data: config.Input,
		}
		if c.Language != "" {
			payload.Context.Client.Hl = c.Language
		}
		if c.Region != "" {
			payload.Context.Client.Gl = c.Region
		}

		bodyData, err := json.Marshal(&payload)
		if err != nil {
			return err
		}
		body = bytes.NewReader(bodyData)
	}

	url := "https://www.youtube.com/youtubei/v1" + config.Path + "?" + query
	req, err := http.NewRequest(config.Method, url, body)
	if err != nil {
		return err
	}

	if config.Input != nil {
		req.Header.Set("Content-Type", "application/json")
	}
	if c.UserAgent != "" {
		req.Header.Set("User-Agent", c.UserAgent)
	} else {
		req.Header.Set("User-Agent", pkgPath+" "+Version())
	}

	var resp *http.Response
	if c.HTTPClient != nil {
		resp, err = c.HTTPClient.Do(req)
	} else {
		resp, err = http.DefaultClient.Do(req)
	}
	if err != nil {
		return err
	}
	defer resp.Body.Close() //nolint:errcheck

	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		err = newError(resp)
	} else if config.Output != nil {
		err = json.UnmarshalRead(resp.Body, config.Output)
	}
	return err
}

func newError(resp *http.Response) Error {
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		body = nil
	}
	return Error{
		StatusCode: resp.StatusCode,
		Message:    bytes2string(body),
	}
}

type Error struct {
	StatusCode int
	Message    string
}

func (e Error) Error() string {
	statusText := http.StatusText(e.StatusCode)
	sz := 3 + 1 + len(statusText)
	if e.Message != "" {
		sz += 3 + quotedLen(e.Message)
	}
	dst := make([]byte, 0, sz)
	dst = appendInt(dst, e.StatusCode)
	dst = append(dst, ' ')
	dst = append(dst, statusText...)
	if e.Message != "" {
		dst = append(dst, " - "...)
		dst = appendQuote(dst, e.Message)
	}
	return bytes2string(dst)
}
