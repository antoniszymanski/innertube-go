package shared

import (
	"time"

	"github.com/dop251/goja"
	"github.com/dsnet/try"
)

type OAuthOptions struct {
	Enabled      bool   `js:"enabled"`
	RefreshToken string `js:"refreshToken"`
}

type OAuthProps struct {
	Token        string    `js:"token"`
	ExpiresAt    time.Time `js:"expiresAt"`
	RefreshToken string    `js:"refreshToken"`
}

type PotOptions struct {
	Token       string `js:"token"`
	VisitorData string `js:"visitorData"`
}

type ClientOptions struct {
	APIKey               string
	BaseURL              string
	ClientName           string
	ClientVersion        string
	YoutubeClientOptions map[string]any
	InitialCookie        string
	OAuth                OAuthOptions
	Pot                  PotOptions
}

func (x *ClientOptions) ToValue(vm *goja.Runtime) (val goja.Value, err error) {
	defer try.Handle(&err)
	obj := vm.NewObject()
	if x.APIKey != "" {
		try.E(obj.Set("apiKey", x.APIKey))
	}
	if x.BaseURL != "" {
		try.E(obj.Set("baseUrl", x.BaseURL))
	}
	if x.ClientName != "" {
		try.E(obj.Set("clientName", x.ClientName))
	}
	if x.ClientVersion != "" {
		try.E(obj.Set("clientVersion", x.ClientVersion))
	}
	if x.YoutubeClientOptions != nil {
		try.E(obj.Set("youtubeClientOptions", x.YoutubeClientOptions))
	}
	if x.InitialCookie != "" {
		try.E(obj.Set("initialCookie", x.InitialCookie))
	}
	if x.OAuth.Enabled {
		try.E(obj.Set("oauth", x.OAuth))
	}
	if (x.Pot != PotOptions{}) {
		try.E(obj.Set("pot", x.Pot))
	}
	return obj, nil
}
