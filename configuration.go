package configuration

import (
	"strings"

	"github.com/mitchellh/mapstructure"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
)

type Configuration struct {
	config            *viper.Viper
	decoderConfigOpts []viper.DecoderConfigOption
}

func (c *Configuration) Unmarshal(dst any) error {
	err := c.config.ReadInConfig()
	if err != nil {
		return errors.Wrap(err, "unable to read configuration")
	}

	err = c.config.Unmarshal(dst, c.decoderConfigOpts...)
	if err != nil {
		return errors.Wrap(err, "unable to unmarshal configuration")
	}

	return nil
}

func New() *Configuration {
	config := viper.NewWithOptions(
		viper.KeyDelimiter("_"),
		viper.EnvKeyReplacer(strings.NewReplacer("-", "_")),
	)

	config.AutomaticEnv()
	config.SetConfigName("config")
	config.SetConfigType("yaml")
	config.AddConfigPath(".")

	decoderConfigOpts := []viper.DecoderConfigOption{
		func(config *mapstructure.DecoderConfig) {
			config.Squash = true
		},
		func(config *mapstructure.DecoderConfig) {
			config.TagName = "config"
		},
	}

	return &Configuration{
		config:            config,
		decoderConfigOpts: decoderConfigOpts,
	}
}
