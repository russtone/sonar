package server

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	"github.com/spf13/viper"

	"github.com/russtone/sonar/internal/database"
)

type Config struct {
	DB     database.Config `json:"db"`
	Domain string          `json:"domain"`
	IP     string          `json:"ip"`

	DNS DNSConfig `json:"dns"`

	TLS TLSConfig `json:"tls"`

	Modules ModulesConfig `json:"modules"`
}

func init() {
	viper.SetDefault("db.dsn", "")
	viper.SetDefault("db.migrations", "/opt/app/migrations")

	viper.SetDefault("domain", "")
	viper.SetDefault("ip", "")

	viper.SetDefault("dns.zone", "")

	viper.SetDefault("tls.type", "letsencrypt")

	viper.SetDefault("tls.custom.key", "")
	viper.SetDefault("tls.custom.cert", "")

	viper.SetDefault("tls.letsencrypt.email", "")
	viper.SetDefault("tls.letsencrypt.directory", "./tls")
	viper.SetDefault("tls.letsencrypt.caDirUrl", "https://acme-v02.api.letsencrypt.org/directory")
	viper.SetDefault("tls.letsencrypt.caInsecure", false)

	viper.SetDefault("modules.enabled", "api")

	viper.SetDefault("modules.api.admin", "")
	viper.SetDefault("modules.api.port", 31337)

	viper.SetDefault("modules.telegram.admin", 0)
	viper.SetDefault("modules.telegram.token", "")
	viper.SetDefault("modules.telegram.proxy", "")

	viper.SetDefault("modules.lark.admin", "")
	viper.SetDefault("modules.lark.appId", "")
	viper.SetDefault("modules.lark.appSecret", "")
	viper.SetDefault("modules.lark.verificationToken", "")
	viper.SetDefault("modules.lark.encryptKey", "")
	viper.SetDefault("modules.lark.tlsEnabled", true)
	viper.SetDefault("modules.lark.proxyUrl", "")
	viper.SetDefault("modules.lark.proxyInsecure", false)
}

func (c Config) Validate() error {
	return validation.ValidateStruct(&c,
		validation.Field(&c.DB),
		validation.Field(&c.Domain, validation.Required, is.Domain),
		validation.Field(&c.IP, validation.Required, is.IP),
		validation.Field(&c.DNS),
		validation.Field(&c.TLS),
		validation.Field(&c.Modules),
	)
}
