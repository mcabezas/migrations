package migrations

import (
	"github.com/mcabezas/migrations/commons"
	"github.com/mcabezas/migrations/persistence"
	"github.com/mcabezas/migrations/scriptfinder"
)

type Strategy int
type scriptFinderBackend int
type persistenceBackend int

const (
	merge Strategy = iota
)

const (
	file    scriptFinderBackend = iota
	binData scriptFinderBackend = iota
)

const (
	memory   persistenceBackend = iota
	postgres persistenceBackend = iota
)

type Config struct {
	deltaPath     string
	mergeVersions bool
	skipMigration bool
	ScriptFinder  scriptFinderBackend
	Persistence   persistenceBackend
	strategy      Strategy
}

type Migration struct {
	conf Config
	commons.ScriptFinderBackend
	commons.PersistenceBackend
	packageDeltas map[string]string
	appliedDeltas map[string]string
}

func New(conf Config) *Migration {
	var finder commons.ScriptFinderBackend
	switch conf.ScriptFinder {
	case file:
		finder = scriptfinder.NewFileBackend("")
	}
	var persistenceBackend commons.PersistenceBackend
	switch conf.Persistence {
	case memory:
		persistenceBackend = persistence.NewInMemoryBackend()
	}

	return &Migration{conf: conf, ScriptFinderBackend: finder, PersistenceBackend: persistenceBackend}
}
