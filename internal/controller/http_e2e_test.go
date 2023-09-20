package controller_test

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/p1xray/port-api/internal/controller"
	"github.com/p1xray/port-api/internal/repository/inmem"
	"github.com/p1xray/port-api/internal/services"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type HttpTestSuite struct {
	suite.Suite
	portService controller.PortService
	portHandler controller.PortHandler
}

func NewHttpTestSuite() *HttpTestSuite {
	suite := &HttpTestSuite{}

	portRepo := inmem.NewPortRepository()

	suite.portService = services.NewPortService(portRepo)

	suite.portHandler = controller.NewPortHandler(suite.portService)

	return suite
}

func TestHttpTestSuite(t *testing.T) {
	suite.Run(t, NewHttpTestSuite())
}

func (suite *HttpTestSuite) TestUploadPorts() {
	// Получаем json тело запроса из файла
	portsRequest, err := os.ReadFile("testdata/ports_request.json")
	require.NoError(suite.T(), err)

	// Подсчитываем количество портов в полученном json
	requestPortsTotal := countJsonPorts(suite.T(), portsRequest)

	// Получаем json ответ от файла
	portsResponse, err := os.ReadFile("testdata/ports_response.json")
	require.NoError(suite.T(), err)

	// Формируем запрос на загрузку портов
	req := httptest.NewRequest(http.MethodPost, "/ports", bytes.NewBuffer(portsRequest))
	w := httptest.NewRecorder()

	// Запускаем запрос на загрузку портов
	suite.portHandler.UploadPorts(w, req)

	// Получаем результат выполненного запроса
	res := w.Result()
	defer res.Body.Close()

	// Читаем тело ответа
	data, err := io.ReadAll(res.Body)
	require.NoError(suite.T(), err)

	assert.Equal(suite.T(), http.StatusOK, res.StatusCode)
	assert.Equal(suite.T(), portsResponse, data)

	// Сравниваем количество портов в переданном json и в хранилище
	storedPortsTotal, err := suite.portService.CountPorts(context.Background())
	require.NoError(suite.T(), err)

	assert.Equal(suite.T(), requestPortsTotal, storedPortsTotal)
}

func (suite *HttpTestSuite) TestUploadPorts_badJson() {
	// Формируем запрос на загрузку портов
	req := httptest.NewRequest(http.MethodPost, "/ports", bytes.NewBuffer([]byte("jgfhdkgjhfd")))
	w := httptest.NewRecorder()

	// Запускаем запрос на загрузку портов
	suite.portHandler.UploadPorts(w, req)

	// Получаем результат выполненного запроса
	res := w.Result()
	defer res.Body.Close()

	require.Equal(suite.T(), http.StatusBadRequest, res.StatusCode)
}

func countJsonPorts(t *testing.T, portsJson []byte) int {
	t.Helper()

	var ports map[string]struct{}
	err := json.Unmarshal(portsJson, &ports)
	require.NoError(t, err)
	return len(ports)
}
