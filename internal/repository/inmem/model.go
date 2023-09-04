package inmem

import (
	"errors"
	"time"

	"github.com/p1xray/port-api/internal/domain"
)

type Port struct {
	Id          string
	Name        string
	Code        string
	City        string
	Country     string
	Alias       []string
	Regions     []string
	Coordinates []float64
	Province    string
	Timezone    string
	Unlocs      []string

	CreatedAt time.Time
	UpdatedAt time.Time
}

func (p *Port) ToDomain() (*domain.Port, error) {
	if p == nil {
		return nil, errors.New("store port is nil")
	}
	return domain.NewPort(
		p.Id,
		p.Name,
		p.Code,
		p.City,
		p.Country,
		append([]string(nil), p.Alias...),
		append([]string(nil), p.Regions...),
		append([]float64(nil), p.Coordinates...),
		p.Province,
		p.Timezone,
		append([]string(nil), p.Unlocs...),
	)
}
