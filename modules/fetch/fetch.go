// SPDX-FileCopyrightText: 2025 Antoni Szymański
// SPDX-License-Identifier: MPL-2.0

package fetch

import (
	"compress/gzip"
	"context"
	"errors"
	"io"
	"net/http"
	"strings"

	"github.com/antoniszymanski/innertube-go/internal"
)

type Request struct {
	URL     string      `js:"url"`
	Body    string      `js:"body"`
	Headers http.Header `js:"headers"`
	Method  string      `js:"method"`
}

type response struct {
	Body       string              `js:"body"`
	Headers    map[string][]string `js:"headers"`
	OK         bool                `js:"ok"`
	Redirected bool                `js:"redirected"`
	Status     int                 `js:"status"`
	StatusText string              `js:"statusText"`
	URL        string              `js:"url"`
}

type redirectedKey struct{}

var client = &http.Client{
	Transport: &http.Transport{
		DisableCompression: true,
	},
	CheckRedirect: func(req *http.Request, via []*http.Request) error {
		if len(via) >= 10 {
			return errors.New("stopped after 10 redirects")
		}
		*req = *req.WithContext(context.WithValue(req.Context(), redirectedKey{}, true))
		return nil
	},
}

func fetch(r Request) (*response, error) {
	req, err := http.NewRequest(
		r.Method,
		r.URL,
		strings.NewReader(r.Body),
	)
	if err != nil {
		return nil, err
	}
	req.Header = make(http.Header)
	for key, value := range r.Headers {
		key = http.CanonicalHeaderKey(key)
		if key != "Accept-Encoding" {
			req.Header[key] = value
		}
	}
	req.Header["Accept-Encoding"] = []string{"gzip"}

	*req = *req.WithContext(context.WithValue(req.Context(), redirectedKey{}, false))

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close() //nolint:errcheck

	var bodyReader io.Reader
	switch resp.Header.Get("Content-Encoding") {
	case "gzip":
		rc, err := gzip.NewReader(resp.Body)
		if err != nil {
			return nil, err
		}
		defer rc.Close() //nolint:errcheck
		bodyReader = rc
	default:
		bodyReader = resp.Body
	}

	bodyData, err := io.ReadAll(bodyReader)
	if err != nil {
		return nil, err
	}

	return &response{
		Body:       internal.BytesToString(bodyData),
		Headers:    resp.Header,
		OK:         200 <= resp.StatusCode && resp.StatusCode <= 299,
		Redirected: resp.Request.Context().Value(redirectedKey{}).(bool),
		Status:     resp.StatusCode,
		StatusText: http.StatusText(resp.StatusCode),
		URL:        resp.Request.URL.String(),
	}, nil
}
