package envx

import (
	envLoader "github.com/joho/godotenv"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
	"io/fs"
)

func LoadFromDotEnv() error {
	err := envLoader.Load()
	if err == nil {
		log.Info().Msg("Loaded configuration from local .env file")
		return nil
	}

	if !errors.Is(err, fs.ErrNotExist) {
		return errors.Wrap(err, "failed parsing .env file")
	}

	return nil
}
