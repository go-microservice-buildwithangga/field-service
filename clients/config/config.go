package clients

import "github.com/parnurzeal/gorequest"

type ClientsConfig struct {
	client       *gorequest.SuperAgent
	baseUrl      string
	signatureKey string
}

type IClientConfig interface {
	Client() *gorequest.SuperAgent
	BaseUrl() string
	SignatureKey() string
}

type Option func(*ClientsConfig)

func NewClientConfig(opts ...Option) IClientConfig {
	cfg := &ClientsConfig{
		client: gorequest.New().
			Set("content-type", "application/json").
			Set("accept", "application/json"),
	}
	for _, opt := range opts {
		opt(cfg)
	}
	return cfg

}

func (c *ClientsConfig) Client() *gorequest.SuperAgent {
	return c.client
}
func (c *ClientsConfig) BaseUrl() string {
	return c.baseUrl
}
func (c *ClientsConfig) SignatureKey() string {
	return c.signatureKey
}

func WithBaseUrl(baseUrl string) Option {
	return func(cfg *ClientsConfig) {
		cfg.baseUrl = baseUrl
	}
}

func WithSignatureKey(signatureKey string) Option {
	return func(cfg *ClientsConfig) {
		cfg.signatureKey = signatureKey
	}
}
