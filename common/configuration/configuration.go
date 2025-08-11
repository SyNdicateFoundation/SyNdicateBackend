package configuration

import (
	"SyNdicateBackend/common/logger"
	"github.com/BurntSushi/toml"
	"os"
)

// HTTPSServer -------------- HTTPS config holders --------------
type HTTPSServer struct {
	Enabled          bool                  `toml:"enabled"`
	Address          string                `toml:"address"`
	Port             int                   `toml:"port"`
	APIUserAgent     string                `toml:"api_user_agent"`
	TlsConfiguration HttpsTlsConfiguration `toml:"tls_configuration"`
}
type HttpsTlsConfiguration struct {
	Enable   bool   `toml:"enable"`
	CertFile string `toml:"cert_file"`
	KeyFile  string `toml:"key_file"`
}

// SQLLiteConfiguration -------------- SQLLite config holders --------------
type SQLLiteConfiguration struct {
	DatabaseFileLocation string `toml:"file_location"`
}

type Tokenizer struct {
	TokenizerSecret string `toml:"tokenizer_secret"`
	TokenExpiration int    `toml:"token_expiration"`
}

type Holder struct {
	debug                bool                 `toml:"debug"`
	HTTPSServer          HTTPSServer          `toml:"https_server"`
	SQLLiteConfiguration SQLLiteConfiguration `toml:"database"`
	Tokenizer            Tokenizer            `toml:"tokenizer"`
}

var ConfigHolder Holder

func SetupConfig() {
	if _, err := os.Stat("config.toml"); os.IsNotExist(err) {
		ConfigHolder = Holder{
			debug: false,
			HTTPSServer: HTTPSServer{
				Enabled:      true,
				Address:      "0.0.0.0",
				Port:         2009,
				APIUserAgent: "LiteGuard Client 1.0/b (Software)",
				TlsConfiguration: HttpsTlsConfiguration{
					Enable:   false,
					CertFile: "cert.pem",
					KeyFile:  "key.pem",
				},
			},
			SQLLiteConfiguration: SQLLiteConfiguration{
				DatabaseFileLocation: "database.db",
			},
			Tokenizer: Tokenizer{
				TokenizerSecret: "TBJU8H91IEJu7g/fygTjEKM5kBx8qiDdTouuMmYQd3jlAt62Jmwq/3X7S1nmgcsE",
				TokenExpiration: 10,
			},
		}
		file, err := os.Create("config.toml")
		if err != nil {
			logger.Logger.Error(err)
		}
		defer func(file *os.File) {

			if err := file.Close(); err != nil {
				logger.Logger.Error(err)
			}
		}(file)

		encoder := toml.NewEncoder(file)
		if err := encoder.Encode(ConfigHolder); err != nil {
			logger.Logger.Error(err)
		}
	}

	if _, err := toml.DecodeFile("config.toml", &ConfigHolder); err != nil {
		logger.Logger.Error(err)
	}
}
