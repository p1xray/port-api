package domain

import "fmt"

// Доменная модель порта
type Port struct {
	id          string
	name        string
	code        string
	city        string
	country     string
	alias       []string
	regions     []string
	coordinates []float64
	province    string
	timezone    string
	unlocs      []string
}

// Создает доменную модель порта
func NewPort(
	id, name, code, city, country string,
	alias, regions []string,
	coordinates []float64,
	province, timezone string,
	unlocs []string) (*Port, error) {

	if id == "" {
		return nil, fmt.Errorf("%w: port id is required", ErrRequired)
	}
	if name == "" {
		return nil, fmt.Errorf("%w: port name is required", ErrRequired)
	}
	if city == "" {
		return nil, fmt.Errorf("%w: port city is required", ErrRequired)
	}
	if country == "" {
		return nil, fmt.Errorf("%w: port country is required", ErrRequired)
	}

	port := &Port{
		id:          id,
		name:        name,
		code:        code,
		city:        city,
		country:     country,
		alias:       alias,
		regions:     regions,
		coordinates: coordinates,
		province:    province,
		timezone:    timezone,
		unlocs:      unlocs,
	}

	return port, nil
}

// Возвращает идентификатор порта
func (p *Port) Id() string {
	return p.id
}

// Возвращает наименование порта
func (p *Port) Name() string {
	return p.name
}

// Возвращает код порта
func (p *Port) Code() string {
	return p.code
}

// Возвращает город порта
func (p *Port) City() string {
	return p.city
}

// Возвращает страну порта
func (p *Port) Country() string {
	return p.country
}

// Возвращает список алиасов порта
func (p *Port) Alias() []string {
	return p.alias
}

// Возвращает список регионов порта
func (p *Port) Regions() []string {
	return p.regions
}

// Возвращает координаты порта
func (p *Port) Coordinates() []float64 {
	return p.coordinates
}

// Возвращает провинцию порта
func (p *Port) Province() string {
	return p.province
}

// Возвращает часовой пояс порта
func (p *Port) Timezone() string {
	return p.timezone
}

// Возвращает список незаблокированных портов
func (p *Port) Unlocs() []string {
	return p.unlocs
}
