package inmem

import (
	"errors"
	"time"

	"github.com/p1xray/port-api/internal/domain"
)

// Модель порта для хранения в памяти
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

// Возвращает копию модели порта
func (p *Port) Copy() *Port {
	if p == nil {
		return nil
	}

	return &Port{
		Id:          p.Id,
		Name:        p.Name,
		Code:        p.Code,
		City:        p.City,
		Country:     p.Country,
		Alias:       append([]string(nil), p.Alias...),
		Regions:     append([]string(nil), p.Regions...),
		Coordinates: append([]float64(nil), p.Coordinates...),
		Province:    p.Province,
		Timezone:    p.Timezone,
		Unlocs:      append([]string(nil), p.Unlocs...),
		CreatedAt:   p.CreatedAt,
		UpdatedAt:   p.UpdatedAt,
	}
}

// Возвращает доменную модель порта, созданную на основе модели порта, хранящейся в памяти
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

// Заполняет поля модели порта полями доменной модели порта
func (p *Port) FillFromDomain(domainPort *domain.Port) {
	p.Id = domainPort.Id()
	p.Name = domainPort.Name()
	p.Code = domainPort.Code()
	p.City = domainPort.City()
	p.Country = domainPort.Country()
	p.Alias = append([]string(nil), domainPort.Alias()...)
	p.Regions = append([]string(nil), domainPort.Regions()...)
	p.Coordinates = append([]float64(nil), domainPort.Coordinates()...)
	p.Province = domainPort.Province()
	p.Timezone = domainPort.Timezone()
	p.Unlocs = append([]string(nil), domainPort.Unlocs()...)
}
