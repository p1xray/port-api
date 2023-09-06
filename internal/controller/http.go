package controller

import (
	"context"
	"errors"
	"log"
	"net/http"

	"github.com/p1xray/port-api/internal/domain"
)

// Сервис портов
type PortService interface {
	GetPort(ctx context.Context, id string) (*domain.Port, error)
	CountPorts(ctx context.Context) (int, error)
	CreateOrUpdatePort(ctx context.Context, port *domain.Port) error
}

// Обработчик запросов к портам
type PortHandler struct {
	portService PortService
}

// Создает новый обработчик запросов к портам
func NewPortHandler(ps PortService) PortHandler {
	return PortHandler{
		portService: ps,
	}
}

// Возвращает порт по переданному в запросе идентификатору
func (ph PortHandler) GetPort(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	port, err := ph.portService.GetPort(r.Context(), id)
	if err != nil {
		if errors.Is(err, domain.ErrNotFound) {
			NotFound("port-not-found", err, w, r)
			return
		}

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

// Возвращает количетсво хранящихся портов
func (ph PortHandler) CountPorts(w http.ResponseWriter, r *http.Request) {
	count, err := ph.portService.CountPorts(r.Context())
	if err != nil {
		RespondWithError(err, w, r)
		return
	}

	RespondOK(count, w, r)
}

// Загружает переданные в json формате порты в хранилище
func (ph PortHandler) UploadPorts(w http.ResponseWriter, r *http.Request) {
	log.Println("uploading ports")

	// Канал для передачи распаршенных портов
	portChan := make(chan Port)

	// Канал для передачи ошибки парсинга
	errChan := make(chan error)

	// Канал для передачи завершения работы парсера
	doneChan := make(chan struct{})

	go func() {
		err := parsePorts(r.Context(), r.Body, portChan)
		if err != nil {
			errChan <- err
		} else {
			doneChan <- struct{}{}
		}
	}()

	portCnt := 0
	for {
		select {
		case <-r.Context().Done():
			log.Printf("request cancelled")
			return
		case <-doneChan:
			log.Printf("finished reading ports")
			RespondOK(portCnt, w, r)
			return
		case port := <-portChan:
			portCnt++
			log.Printf("[%d] received port: %+v", portCnt, port)

			domainPort, err := port.ToDomain()
			if err != nil {
				BadRequest("port-to-domain", err, w, r)
				return
			}

			if err := ph.portService.CreateOrUpdatePort(r.Context(), domainPort); err != nil {
				RespondWithError(err, w, r)
				return
			}
		}

	}
}
