package config

type Account struct {
	AccessToken     string `json:"accessToken"`
	OpenId          string `json:"openId"`
	Sig             string `json:"sig"`
	Timestamp       string `json:"timestamp"`
	MsdkEncodeParam string `json:"msdkEncodeParam"`
	MsdkToken       string `json:"msdkToken"`
}

type Wxpush struct {
	Pushplus string `json:"pushplus"`
}

// Qywxpush 使用企业微信机器人推送,使用webhook协议推送消息
type Qywxpush struct {
	Webhookurl string `json:"webhookurl"`
}

type Config struct {
	Account  []Account
	Wxpush   `json:"wxpush"`
	Qywxpush `json:"qywxpush"`
}
