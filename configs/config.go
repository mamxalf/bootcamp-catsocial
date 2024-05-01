package configs

import (
	"sync"
	"time"

	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
)

// Config is a struct that will receive configuration options via environment
// variables.
type Config struct {
	AppName string `mapstructure:"APP_Name"`
	AppPort string `mapstructure:"APP_PORT"`
	AppEnv  string `mapstructure:"APP_ENV"`
	AppUrl  string `mapstructure:"APP_URL"`

	CorsAllowCredentials bool     `mapstructure:"CORS_ALLOW_CREDENTIALS"`
	CorsAllowedHeaders   []string `mapstructure:"CORS_ALLOWED_HEADERS"`
	CorsAllowedMethods   []string `mapstructure:"CORS_ALLOWED_METHODS"`
	CorsAllowedOrigins   []string `mapstructure:"CORS_ALLOWED_ORIGINS"`
	CorsEnable           bool     `mapstructure:"CORS_ENABLE"`
	CorsMaxAgeSeconds    int      `mapstructure:"CORS_MAX_AGE_SECONDS"`

	LogLevel   string        `mapstructure:"LOG_LEVEL"`
	JwtSecret  string        `mapstructure:"JWT_SECRET"`
	JwtExpiry  time.Duration `mapstructure:"JWT_EXPIRY"`
	BcryptSalt int           `mapstructure:"BCRYPT_SALT"`

	DbName     string `mapstructure:"DB_NAME"`
	DbPort     string `mapstructure:"DB_PORT"`
	DbHost     string `mapstructure:"DB_HOST"`
	DbUsername string `mapstructure:"DB_USERNAME"`
	DbPassword string `mapstructure:"DB_PASSWORD"`
	DbParams   string `mapstructure:"DB_PARAMS"`
}

var (
	conf Config
	once sync.Once
)

// Get are responsible to load env and get data a return the struct
func Get() *Config {
	viper.SetConfigFile(".env")
	err := viper.ReadInConfig()

	if err != nil {
		log.Fatal().Err(err).Msg("Failed reading config file")
	}

	once.Do(func() {
		log.Info().Msg("Service configuration initialized.")
		err = viper.Unmarshal(&conf)
		if err != nil {
			log.Fatal().Err(err)
		}
	})

	return &conf
}
