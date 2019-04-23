package config

import "fmt"

// HTTP contains the HTTP configuration.
type HTTP struct {
	Hostname   string      `default:""`
	Port       int         `default:"3000"`
	IsSecure   bool        `yaml:"is_secure"`
	SecureHTTP *SecureHTTP `yaml:"secure_http"`
}

// SecureHTTP is the configuration for HTTPS.
type SecureHTTP struct {
	KeyFilePath  string `yaml:"key_file_path"`
	CertFilePath string `yaml:"cert_file_path"`
	Port         int    `default:"3443"`
}

// Address returns the HTTP address based on the configuration.
func (h *HTTP) Address(withScheme bool) string {
	port := h.Port
	if h.IsSecure && h.SecureHTTP != nil {
		port = h.SecureHTTP.Port
	}
	var scheme string
	if withScheme {
		scheme = "http://"
		if h.IsSecure {
			scheme = "https://"
		}
	}
	return fmt.Sprintf("%s%s:%d", scheme, h.Hostname, port)
}
