package shared

import "time"

type OAuthProps struct {
	Token        string    `js:"token"`
	ExpiresAt    time.Time `js:"expiresAt"`
	RefreshToken string    `js:"refreshToken"`
}
