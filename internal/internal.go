// SPDX-FileCopyrightText: 2025 Antoni Szymański
// SPDX-License-Identifier: MPL-2.0

package internal

type Payload struct {
	Context Context           `json:"context"`
	Data    map[string]string `json:",inline"`
}

type Context struct {
	Client Client `json:"client"`
}

type Client struct {
	ClientName    string `json:"clientName"`
	ClientVersion string `json:"clientVersion"`
	Hl            string `json:"hl"`
	Gl            string `json:"gl"`
}
