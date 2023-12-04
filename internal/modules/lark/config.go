package lark

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type Config struct {
	Admin             string `json:"admin"`
	AppID             string `json:"appId"`
	AppSecret         string `json:"appSecret"`
	VerificationToken string `json:"verificationToken"`
	EncryptKey        string `json:"encryptKey"`
	TLSEnabled        bool   `json:"tlsEnabled"`
	ProxyURL          string `json:"proxyUrl"`
	ProxyInsecure     bool   `json:"proxyInsecure"`
}

func (c Config) Validate() error {
	return validation.ValidateStruct(&c,
		validation.Field(&c.Admin, validation.Required),
		validation.Field(&c.AppID, validation.Required),
		validation.Field(&c.AppSecret, validation.Required),
		validation.Field(&c.VerificationToken, validation.Required),
		validation.Field(&c.EncryptKey),
	)
}
