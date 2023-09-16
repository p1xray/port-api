package inmem_test

import (
	"context"
	"testing"

	"github.com/google/uuid"
	"github.com/p1xray/port-api/internal/domain"
	"github.com/p1xray/port-api/internal/repository/inmem"
	"github.com/stretchr/testify/assert"
)

func TestPortRepository_CreateOrUpdatePort(t *testing.T) {
	t.Parallel()

	repo := inmem.NewPortRepository()

	t.Run("create port", func(t *testing.T) {
		t.Parallel()

		randomPort := newRandomDomainPort(t)
		err := repo.CreateOrUpdatePort(context.Background(), randomPort)
		assert.NoError(t, err)

		port, err := repo.GetPort(context.Background(), randomPort.Id())
		assert.NoError(t, err)

		assert.Equal(t, randomPort, port)
	})

	t.Run("update port", func(t *testing.T) {
		t.Parallel()

		randomPort := newRandomDomainPort(t)
		err := repo.CreateOrUpdatePort(context.Background(), randomPort)
		assert.NoError(t, err)

		beforeUpdatePort, err := repo.GetPort(context.Background(), randomPort.Id())
		assert.NoError(t, err)

		assert.Equal(t, randomPort, beforeUpdatePort)

		err = randomPort.SetName("new name")
		assert.NoError(t, err)

		err = repo.CreateOrUpdatePort(context.Background(), randomPort)
		assert.NoError(t, err)

		updatedPort, err := repo.GetPort(context.Background(), randomPort.Id())
		assert.NoError(t, err)

		assert.NotEqual(t, beforeUpdatePort.Name(), updatedPort.Name())
	})

	t.Run("nil port", func(t *testing.T) {
		t.Parallel()
		err := repo.CreateOrUpdatePort(context.Background(), nil)
		assert.ErrorIs(t, err, domain.ErrNil)
	})
}

func newRandomDomainPort(t *testing.T) *domain.Port {
	t.Helper()

	randomId := uuid.New().String()
	port, err := domain.NewPort(randomId, randomId, randomId, randomId, randomId,
		[]string{randomId}, []string{randomId}, []float64{1.0, 2.0}, randomId, randomId, nil)

	assert.NoError(t, err)
	return port
}
