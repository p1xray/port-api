package controller

import (
	"context"
	"net/http"

	"github.com/p1xray/port-api/internal/domain"
)

// Сервис портов
type PortService interface {
	GetPort(ctx context.Context, id string) (*domain.Port, error)
}

type PortHandler struct {
	portService PortService
}

func NewPortHandler(ps PortService) PortHandler {
	return PortHandler{
		portService: ps,
	}
}

func (ph PortHandler) GetPort(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	port, err := ph.portService.GetPort(r.Context(), id)
	if err != nil {
		RespondWithError(err, w, r)
		return
	}

	response := Port{
		Id:          port.Id(),
		Name:        port.Name(),
		City:        port.City(),
		Country:     port.Country(),
		Alias:       port.Alias(),
		Regions:     port.Regions(),
		Coordinates: port.Coordinates(),
		Province:    port.Province(),
		Timezone:    port.Timezone(),
		Unlocs:      port.Unlocs(),
	}

	RespondOK(response, w, r)
}
