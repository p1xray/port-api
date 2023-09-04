package services

import (
	"context"

	"github.com/google/uuid"
	"github.com/p1xray/port-api/internal/domain"
)

// Сервис портов
type PortService struct {
}

// Создает новый сервис портов
func NewPortService() PortService {
	return PortService{}
}

// Возвращает порт по переданному идентификатору
func (ps PortService) GetPort(ctx context.Context, id string) (*domain.Port, error) {
	randomId := uuid.New().String()
	return domain.NewPort(randomId, randomId, randomId, randomId, randomId,
		[]string{randomId}, []string{randomId}, []float64{1.0, 2.0}, randomId, randomId, nil)
}
