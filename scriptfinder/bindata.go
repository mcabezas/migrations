package scriptfinder

import (
	"github.com/mcabezas/migrations/commons"
)

type bindataBackend struct {
	path string
}

func NewBindataScriptFinder() commons.ScriptFinderBackend {
	return &bindataBackend{}
}

func (m *bindataBackend) GetRequiredDeltas() ([]*commons.Delta, error) {
	return nil, nil
}
