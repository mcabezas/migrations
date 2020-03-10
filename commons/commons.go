package commons

import "errors"

type Delta struct {
	ID         string
	UpScript   []byte
	DownScript []byte
	Dirty      bool
}

func ValidateDelta(delta *Delta) error {
	if delta.ID == "" {
		return errors.New("DELTA ID can not be empty")
	}
	if delta.UpScript == nil {
		return errors.New("DELTA UP can not be nil")
	}
	if delta.DownScript == nil {
		return errors.New("DELTA DOWN can not be nil")
	}
	return nil
}

type ScriptFinderBackend interface {
	GetRequiredDeltas() ([]*Delta, error)
}

type PersistenceBackend interface {
	Setup()
	GetAppliedDeltas() ([]*Delta, error)
	ApplyDelta(versionID string, content []byte)
}
