package storage

import (
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	flag "github.com/spf13/pflag"
	"github.com/spf13/viper"
	"os"
)

type Config struct {
	ServiceAccountFile string
}

func defineFlags(flagSet *flag.FlagSet) {
	flagSet.String("gcs_service_account", "", "Service Account Credential File for GCS access.")
}

func getConfig() (*Config, error) {
	cwd, err := os.Getwd()
	if err != nil {
		return nil, err
	}

	flags := flag.NewFlagSet("config", flag.ExitOnError)
	defineFlags(flags)

	if err = flags.Parse(os.Args[1:]); err != nil {
		return nil, err
	}

	viper.SetConfigName("ponzu") // ponzu config file
	viper.SetConfigType("props")
	viper.AddConfigPath(cwd)

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil && errors.As(err, &viper.ConfigFileNotFoundError{}) {
		log.Info("config file not found. will default to provided flags")
		err = nil
	}

	if err = viper.BindPFlags(flags); err != nil {
		return nil, err
	}

	return &Config{
		ServiceAccountFile: viper.GetString("gcs_service_account"),
	}, nil
}
