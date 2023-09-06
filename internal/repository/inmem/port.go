package inmem

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/p1xray/port-api/internal/domain"
)

// Репозиторий портов
type PortRepository struct {
	data map[string]*Port
	mu   sync.RWMutex
}

// Создает новый репозиторий портов
func NewPortRepository() *PortRepository {
	return &PortRepository{
		data: make(map[string]*Port),
	}
}

// Возвращает доменную модель порта по переданному идентификатору
func (pr *PortRepository) GetPort(_ context.Context, id string) (*domain.Port, error) {
	pr.mu.RLock()
	defer pr.mu.RUnlock()

	storePort, exists := pr.data[id]
	if !exists {
		return nil, domain.ErrNotFound
	}

	domainPort, err := storePort.ToDomain()
	if err != nil {
		return nil, fmt.Errorf("storePort.ToDomain failed: %w", err)
	}

	return domainPort, nil
}

// Возвращает количетсво хранящихся в памяти портов
func (pr *PortRepository) CountPorts(_ context.Context) (int, error) {
	pr.mu.RLock()
	defer pr.mu.RUnlock()

	return len(pr.data), nil
}

// Создает или обновляет порт по переданной доменной модели
func (pr *PortRepository) CreateOrUpdatePort(ctx context.Context, port *domain.Port) error {
	if port == nil {
		return domain.ErrNil
	}

	storePort := &Port{}
	storePort.FillFromDomain(port)

	pr.mu.Lock()
	defer pr.mu.Unlock()

	_, exists := pr.data[storePort.Id]
	if exists {
		return pr.updatePort(ctx, storePort)
	} else {
		return pr.createPort(ctx, storePort)
	}
}

// Создает новый порт
func (pr *PortRepository) createPort(ctx context.Context, port *Port) error {
	if port == nil {
		return domain.ErrNil
	}

	now := time.Now()
	port.CreatedAt = now
	port.UpdatedAt = now

	pr.data[port.Id] = port

	return nil
}

// Обновляет существующий порт
func (pr *PortRepository) updatePort(ctx context.Context, port *Port) error {
	if port == nil {
		return domain.ErrNil
	}

	storePort, exists := pr.data[port.Id]
	if !exists {
		return domain.ErrNotFound
	}

	storePortCopy := storePort.Copy()

	storePortCopy.Name = port.Name
	storePortCopy.Code = port.Code
	storePortCopy.City = port.City
	storePortCopy.Country = port.Country
	storePortCopy.Alias = append([]string(nil), port.Alias...)
	storePortCopy.Regions = append([]string(nil), port.Regions...)
	storePortCopy.Coordinates = append([]float64(nil), port.Coordinates...)
	storePortCopy.Province = port.Province
	storePortCopy.Timezone = port.Timezone
	storePortCopy.Unlocs = append([]string(nil), port.Unlocs...)

	port.UpdatedAt = time.Now()

	pr.data[port.Id] = storePortCopy

	return nil
}
