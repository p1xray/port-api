package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
)

// Парсит порты из переданного json
func parsePorts(ctx context.Context, r io.Reader, portChan chan Port) error {
	decoder := json.NewDecoder(r)

	// Читаем открывающий разделитель
	t, err := decoder.Token()
	if err != nil {
		return fmt.Errorf("failed to read opening delimiter: %w", err)
	}

	// Проверяем чтобы открывающий разделитель был '{'
	if t != json.Delim('{') {
		return fmt.Errorf("expected {, got %v", t)
	}

	for decoder.More() {
		if ctx.Err() != nil {
			return ctx.Err()
		}

		// Читаем идентификатор порта
		t, err := decoder.Token()
		if err != nil {
			return fmt.Errorf("failed to read port Id: %w", err)
		}

		// Проверяем чтобы идентификатор был строкой
		portId, ok := t.(string)
		if !ok {
			return fmt.Errorf("expected string, got %v", t)
		}

		// Читаем оставшуюся информацию о порте
		var port Port
		if err := decoder.Decode(&port); err != nil {
			return fmt.Errorf("failed to decode port: %w", err)
		}
		port.Id = portId

		// Отправляем прочитанный порт в канал
		portChan <- port
	}

	return nil
}
