package inmem_test

import (
	"testing"
	"time"

	"github.com/p1xray/port-api/internal/domain"
	"github.com/p1xray/port-api/internal/repository/inmem"
	"github.com/stretchr/testify/assert"
)

func Test_storePortToDomain(t *testing.T) {
	testCases := []struct {
		name       string
		storePort  *inmem.Port
		domainPort *domain.Port
		isValid    bool
	}{
		{
			name:       "should return error when store port is nil",
			storePort:  nil,
			domainPort: nil,
			isValid:    false,
		},
		{
			name:       "should return domain port when store port is not nil",
			storePort:  newTestStorePort(t),
			domainPort: newTestDomainPort(t),
			isValid:    true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			domainPort, err := tc.storePort.ToDomain()
			if tc.isValid {
				assert.NoError(t, err)

				assert.Equal(t, tc.domainPort, domainPort)
			} else {
				assert.Error(t, err)
			}
		})
	}
}

const testString = "test"

func newTestStorePort(t *testing.T) *inmem.Port {
	t.Helper()
	return &inmem.Port{
		Id:          testString,
		Name:        testString,
		Code:        testString,
		City:        testString,
		Country:     testString,
		Alias:       []string{testString},
		Regions:     []string{testString},
		Coordinates: []float64{1.0, 2.0},
		Province:    testString,
		Timezone:    testString,
		Unlocs:      []string{testString},
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
}

func newTestDomainPort(t *testing.T) *domain.Port {
	t.Helper()
	port, err := domain.NewPort(testString, testString, testString, testString, testString,
		[]string{testString}, []string{testString}, []float64{1.0, 2.0}, testString, testString, []string{testString})
	assert.NoError(t, err)
	return port
}
