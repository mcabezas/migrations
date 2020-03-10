package persistence

import (
	"github.com/mcabezas/migrations/commons"
)

type memory struct {

}

func NewInMemoryBackend() commons.PersistenceBackend {
	return &memory{}
}

func (m *memory) Setup() {}

func (m *memory) GetAppliedDeltas() ([]*commons.Delta,error) {
	return nil, nil
}

func (m *memory) ApplyDelta(versionID string, content []byte) {}