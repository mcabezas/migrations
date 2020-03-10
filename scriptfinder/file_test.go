package scriptfinder

import (
	"github.com/stretchr/testify/assert"
	"path"
	"testing"
)

func TestCanFindScriptsFromDir(t *testing.T) {
	finder := NewFileBackend(path.Join("..", "test", "scriptfinder", "deltas"))
	deltas, _ := finder.GetRequiredDeltas()
	assert.Equalf(t, 1, len(deltas), "there must be 2 deltas found!")
}

func TestFailWhenUpDeltaIsMissing(t *testing.T) {
	finder := NewFileBackend(path.Join("..", "test", "scriptfinder", "missingupdelta"))
	deltas, err := finder.GetRequiredDeltas()
	assert.Equalf(t, 0, len(deltas), "there must be 0 deltas found!")
	assert.Error(t, err)
}

func TestFailWhenDownDeltaIsMissing(t *testing.T) {
	finder := NewFileBackend(path.Join("..", "test", "scriptfinder", "missingdowndelta"))
	deltas, err := finder.GetRequiredDeltas()
	assert.Equalf(t, 0, len(deltas), "there must be 0 deltas found!")
	assert.Error(t, err)
}
