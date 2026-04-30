package config

import "time"

var defaultConfig = map[string]interface{}{
	"auth.refresh_expiration_time":          RefreshTokenExpireDuration,
	"auth.access_expiration_time":           AccessTokenExpireDuration,
	"auth.access_subject":                   AccessTokenSubject,
	"auth.refresh_subject":                  RefreshTokenSubject,
	"application.graceful_shutdown_timeout": 5 * time.Second,
}
