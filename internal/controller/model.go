package controller

import "github.com/p1xray/port-api/internal/domain"

type Port struct {
	Id          string    `json:"id"`
	Name        string    `json:"name"`
	Code        string    `json:"code"`
	City        string    `json:"city"`
	Country     string    `json:"country"`
	Alias       []string  `json:"alias"`
	Regions     []string  `json:"regions"`
	Coordinates []float64 `json:"coordinates"`
	Province    string    `json:"province"`
	Timezone    string    `json:"timezone"`
	Unlocs      []string  `json:"unlocs"`
}

func (p *Port) ToDomain() (*domain.Port, error) {
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
