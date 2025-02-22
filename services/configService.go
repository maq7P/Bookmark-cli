package services

import "os"

// ConfigService отвечает за работу с переменными окружения.
type ConfigService struct{}

// NewConfigService создаёт новый экземпляр ConfigService.
func NewConfigService() *ConfigService {
	return &ConfigService{}
}

// GetEnvValue возвращает значение переменной окружения или значение по умолчанию.
func (cs *ConfigService) GetEnvValue(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}