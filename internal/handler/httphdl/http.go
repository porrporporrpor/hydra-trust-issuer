package httphdl

import "github.com/porrporporrpor/hydra-trust-issuer/config"

type HTTPHandler struct {
	config config.Config
}

func NewHTTP(cfg config.Config) *HTTPHandler {
	return &HTTPHandler{
		config: cfg,
	}
}
