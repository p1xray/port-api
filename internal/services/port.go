package services

import (
	"context"

	"github.com/p1xray/port-api/internal/domain"
)

// Репозиторий портов
type PortRepository interface {
	GetPort(ctx context.Context, id string) (*domain.Port, error)
	CountPorts(ctx context.Context) (int, error)
}

// Сервис портов
type PortService struct {
	portRepo PortRepository
}

// Создает новый сервис портов
func NewPortService(pr PortRepository) PortService {
	return PortService{
		portRepo: pr,
	}
}

// Возвращает порт по переданному идентификатору
func (ps PortService) GetPort(ctx context.Context, id string) (*domain.Port, error) {
	return ps.portRepo.GetPort(ctx, id)
}

// Возвращает количетсво хранящихся портов
func (ps PortService) CountPorts(ctx context.Context) (int, error) {
	return ps.portRepo.CountPorts(ctx)
}
