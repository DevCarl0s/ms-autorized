package env

import (
	"genesis/ho/ms-autorized/comunes/dominio/constantes"
	comunes_entidades "genesis/ho/ms-autorized/comunes/dominio/entidades"
	"os"
)

func LoadConfig() (*comunes_entidades.ConfigEnv, error) {
	useHTTPS := os.Getenv("USE_HTTPS") == "true"

	return &comunes_entidades.ConfigEnv{
		HOST_IP:       getEnvOrDefault("HOST_IP", constantes.HOST_IP),
		HOST_PORT:     getEnvOrDefault("HOST_PORT", constantes.HOST_PORT),
		USE_HTTPS:     useHTTPS,
		SSL_CERT_FILE: getEnvOrDefault("SSL_CERT_FILE", ""),
		SSL_KEY_FILE:  getEnvOrDefault("SSL_KEY_FILE", ""),
	}, nil
}

func getEnvOrDefault(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}
