package config

import (
	"myapp/pkg/richerror"
	"strings"

	"github.com/knadh/koanf/parsers/yaml"
	"github.com/knadh/koanf/providers/env"
	"github.com/knadh/koanf/providers/file"
	"github.com/knadh/koanf/v2"
)

func Load(configPath string, cfg any) error {
	const op = "config/load"
	var k = koanf.New(".")

	if err := k.Load(file.Provider(configPath), yaml.Parser()); err != nil {

		return richerror.New(op).WithKind(richerror.KindUnexpected).WithErr(err)
	}

	eErr := k.Load(env.Provider("MYAPP_", ".", func(s string) string {

		str := strings.Replace(strings.ToLower(
			strings.TrimPrefix(s, "MYAPP_")), "_", ".", -1)

		return strings.Replace(str, "..", "_", -1)

	}), nil)
	if eErr != nil {
		return richerror.New(op).WithKind(richerror.KindUnexpected).WithErr(eErr)
	}

	if err := k.Unmarshal("", cfg); err != nil {
		return richerror.New(op).WithKind(richerror.KindUnexpected).WithErr(err)
	}

	return nil

}
