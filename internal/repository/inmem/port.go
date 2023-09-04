package inmem

import (
	"context"
	"fmt"
	"sync"

	"github.com/p1xray/port-api/internal/domain"
)

type PortRepository struct {
	data map[string]*Port
	mu   sync.RWMutex
}

func NewPortRepository() *PortRepository {
	return &PortRepository{
		data: make(map[string]*Port),
	}
}

func (pr *PortRepository) GetPort(ctx context.Context, id string) (*domain.Port, error) {
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
